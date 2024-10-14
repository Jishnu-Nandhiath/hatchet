package workflows

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/errgroup"

	"github.com/hatchet-dev/hatchet/internal/cel"
	"github.com/hatchet-dev/hatchet/internal/datautils"
	"github.com/hatchet-dev/hatchet/internal/msgqueue"
	"github.com/hatchet-dev/hatchet/internal/services/shared/tasktypes"
	"github.com/hatchet-dev/hatchet/internal/telemetry"
	"github.com/hatchet-dev/hatchet/internal/telemetry/servertel"
	"github.com/hatchet-dev/hatchet/pkg/repository"
	"github.com/hatchet-dev/hatchet/pkg/repository/prisma/db"
	"github.com/hatchet-dev/hatchet/pkg/repository/prisma/dbsqlc"
	"github.com/hatchet-dev/hatchet/pkg/repository/prisma/sqlchelpers"
)

func (wc *WorkflowsControllerImpl) handleWorkflowRunQueued(ctx context.Context, task *msgqueue.Message) error {
	ctx, span := telemetry.NewSpan(ctx, "handle-workflow-run-queued")
	defer span.End()

	payload := tasktypes.WorkflowRunQueuedTaskPayload{}
	metadata := tasktypes.WorkflowRunQueuedTaskMetadata{}

	err := wc.dv.DecodeAndValidate(task.Payload, &payload)

	if err != nil {
		return fmt.Errorf("could not decode job task payload: %w", err)
	}

	err = wc.dv.DecodeAndValidate(task.Metadata, &metadata)

	if err != nil {

		return fmt.Errorf("could not decode job task metadata: %w", err)
	}

	// get the workflow run in the database
	workflowRun, err := wc.repo.WorkflowRun().GetWorkflowRunById(ctx, metadata.TenantId, payload.WorkflowRunId)

	if err != nil {
		return fmt.Errorf("could not get job run: %w", err)
	}

	isPaused := workflowRun.IsPaused.Valid && workflowRun.IsPaused.Bool

	workflowRunId := sqlchelpers.UUIDToStr(workflowRun.WorkflowRun.ID)

	servertel.WithWorkflowRunModel(span, workflowRun)

	wc.l.Info().Msgf("starting workflow run %s", workflowRunId)

	// determine if we should start this workflow run or we need to limit its concurrency
	// if the workflow has concurrency settings, then we need to check if we can start it
	if workflowRun.ConcurrencyLimitStrategy.Valid && workflowRun.GetGroupKeyRunId.Valid {
		wc.l.Info().Msgf("workflow %s has concurrency settings", workflowRunId)

		groupKeyRunId := sqlchelpers.UUIDToStr(workflowRun.GetGroupKeyRunId)

		if groupKeyRunId == "" {
			return fmt.Errorf("could not get group key run")
		}

		sqlcGroupKeyRun, err := wc.repo.GetGroupKeyRun().GetGroupKeyRunForEngine(ctx, metadata.TenantId, groupKeyRunId)

		if err != nil {
			return fmt.Errorf("could not get group key run for engine: %w", err)
		}

		err = wc.scheduleGetGroupAction(ctx, sqlcGroupKeyRun)

		if err != nil {
			return fmt.Errorf("could not trigger get group action: %w", err)
		}

		return nil
	} else if workflowRun.ConcurrencyLimitStrategy.Valid && workflowRun.ConcurrencyGroupExpression.Valid {
		// if the workflow has a concurrency group expression, then we need to evaluate it
		// and see if we can start the workflow run
		wc.l.Info().Msgf("workflow %s has concurrency settings", workflowRunId)

		groupKey, err := wc.evalWorkflowRunConcurrency(ctx, metadata.TenantId, workflowRunId, workflowRun.ConcurrencyGroupExpression.String)

		if err != nil {
			return fmt.Errorf("could not evaluate concurrency group expression: %w", err)
		}

		if groupKey == nil {
			// fail the workflow run
			workflowRun, err := wc.repo.WorkflowRun().GetWorkflowRunById(ctx, metadata.TenantId, workflowRunId)

			if err != nil {
				return fmt.Errorf("could not get workflow run: %w", err)
			}

			return wc.cancelWorkflowRunJobs(ctx, workflowRun, "GROUP_KEY_REQUIRED")
		}

		wc.checkTenantQueue(ctx, metadata.TenantId)

		return nil
	} else if workflowRun.ConcurrencyLimitStrategy.Valid {
		return fmt.Errorf("workflow run %s has concurrency settings but no group key run", workflowRunId)
	}

	err = wc.queueWorkflowRunJobs(ctx, workflowRun, isPaused)

	if err != nil {
		return fmt.Errorf("could not start workflow run: %w", err)
	}

	return nil
}

func (wc *WorkflowsControllerImpl) evalWorkflowRunConcurrency(ctx context.Context, tenantId, workflowRunId, expr string) (*string, error) {
	input, err := wc.repo.WorkflowRun().GetWorkflowRunInputData(tenantId, workflowRunId)

	if err != nil {
		return nil, fmt.Errorf("could not get workflow run input data: %w", err)
	}

	addMetaRes, err := wc.repo.WorkflowRun().GetWorkflowRunAdditionalMeta(ctx, tenantId, workflowRunId)

	if err != nil {
		return nil, fmt.Errorf("could not get workflow run additional meta: %w", err)
	}

	addMeta := map[string]interface{}{}

	if addMetaRes.AdditionalMetadata != nil {
		err = json.Unmarshal(addMetaRes.AdditionalMetadata, &addMeta)

		if err != nil {
			return nil, fmt.Errorf("could not decode additional metadata: %w", err)
		}
	}

	concurrencyGroupId, err := wc.celParser.ParseAndEvalWorkflowString(expr, cel.NewInput(
		cel.WithInput(input),
		cel.WithAdditionalMetadata(addMeta),
		cel.WithWorkflowRunID(workflowRunId),
	))

	opts := &repository.UpdateWorkflowRunFromGroupKeyEvalOpts{}

	if err != nil {
		opts.Error = repository.StringPtr(fmt.Sprintf("could not evaluate concurrency group expression %s: %s", expr, err.Error()))

		wc.l.Err(err).Msgf("could not evaluate concurrency group expression for tenant %s and workflow run %s", tenantId, workflowRunId)
	} else {
		opts.GroupKey = repository.StringPtr(concurrencyGroupId)
	}

	err = wc.repo.WorkflowRun().UpdateWorkflowRunFromGroupKeyEval(ctx, tenantId, workflowRunId, opts)

	if err != nil {
		return nil, fmt.Errorf("could not update workflow run from group key eval: %w", err)
	}

	return opts.GroupKey, nil
}

func (wc *WorkflowsControllerImpl) handleWorkflowRunFinished(ctx context.Context, task *msgqueue.Message) error {
	ctx, span := telemetry.NewSpan(ctx, "handle-workflow-run-finished")
	defer span.End()

	payload := tasktypes.WorkflowRunFinishedTask{}
	metadata := tasktypes.WorkflowRunFinishedTaskMetadata{}

	err := wc.dv.DecodeAndValidate(task.Payload, &payload)

	if err != nil {
		return fmt.Errorf("could not decode job task payload: %w", err)
	}

	err = wc.dv.DecodeAndValidate(task.Metadata, &metadata)

	if err != nil {
		return fmt.Errorf("could not decode job task metadata: %w", err)
	}

	// get the workflow run in the database
	workflowRun, err := wc.repo.WorkflowRun().GetWorkflowRunById(ctx, metadata.TenantId, payload.WorkflowRunId)

	if err != nil {
		return fmt.Errorf("could not get job run: %w", err)
	}

	workflowRunId := sqlchelpers.UUIDToStr(workflowRun.WorkflowRun.ID)

	servertel.WithWorkflowRunModel(span, workflowRun)

	wc.l.Info().Msgf("finishing workflow run %s", workflowRunId)

	shouldAlertFailure := workflowRun.WorkflowRun.Status == dbsqlc.WorkflowRunStatusFAILED

	// if there's an onFailure job, start that job
	if workflowRun.WorkflowVersion.OnFailureJobId.Valid {
		jobRun, err := wc.repo.JobRun().GetJobRunByWorkflowRunIdAndJobId(
			ctx,
			metadata.TenantId,
			workflowRunId,
			sqlchelpers.UUIDToStr(workflowRun.WorkflowVersion.OnFailureJobId),
		)

		if err != nil {
			return fmt.Errorf("could not get job run: %w", err)
		}

		if !repository.IsFinalJobRunStatus(jobRun.Status) {
			if workflowRun.WorkflowRun.Status == dbsqlc.WorkflowRunStatusFAILED {
				err = wc.startJobRun(ctx, metadata.TenantId, sqlchelpers.UUIDToStr(jobRun.ID))

				if err != nil {
					return err
				}
			} else if jobRun.Status != dbsqlc.JobRunStatus(db.JobRunStatusCancelled) {
				// cancel the onFailure job
				err = wc.mq.AddMessage(
					ctx,
					msgqueue.JOB_PROCESSING_QUEUE,
					tasktypes.JobRunCancelledToTask(metadata.TenantId, sqlchelpers.UUIDToStr(jobRun.ID), nil),
				)

				if err != nil {
					return fmt.Errorf("could not add job run to task queue: %w", err)
				}
			}
		} else {
			shouldAlertFailure = false
		}
	}

	if shouldAlertFailure {
		err := wc.tenantAlerter.HandleAlert(
			sqlchelpers.UUIDToStr(workflowRun.WorkflowRun.TenantId),
		)

		if err != nil {
			wc.l.Err(err).Msg("could not handle alert")
		}
	}

	wc.checkTenantQueue(ctx, metadata.TenantId)

	return nil
}

func (wc *WorkflowsControllerImpl) scheduleGetGroupAction(
	ctx context.Context,
	getGroupKeyRun *dbsqlc.GetGroupKeyRunForEngineRow,
) error {
	ctx, span := telemetry.NewSpan(ctx, "trigger-get-group-action")
	defer span.End()

	tenantId := sqlchelpers.UUIDToStr(getGroupKeyRun.GetGroupKeyRun.TenantId)
	getGroupKeyRunId := sqlchelpers.UUIDToStr(getGroupKeyRun.GetGroupKeyRun.ID)
	workflowRunId := sqlchelpers.UUIDToStr(getGroupKeyRun.WorkflowRunId)

	_, err := wc.repo.GetGroupKeyRun().UpdateGetGroupKeyRun(ctx, tenantId, getGroupKeyRunId, &repository.UpdateGetGroupKeyRunOpts{
		Status: repository.StepRunStatusPtr(db.StepRunStatusPendingAssignment),
	})

	if err != nil {
		return fmt.Errorf("could not update get group key run: %w", err)
	}

	selectedWorkerId, dispatcherId, err := wc.repo.GetGroupKeyRun().AssignGetGroupKeyRunToWorker(
		ctx,
		tenantId,
		getGroupKeyRunId,
	)

	if err != nil {
		if errors.Is(err, repository.ErrNoWorkerAvailable) {
			wc.l.Debug().Msgf("no worker available for get group key run %s, requeueing", getGroupKeyRunId)
			return nil
		}

		return fmt.Errorf("could not assign get group key run to worker: %w", err)
	}

	// send a task to the dispatcher
	err = wc.mq.AddMessage(
		ctx,
		msgqueue.QueueTypeFromDispatcherID(dispatcherId),
		getGroupActionTask(
			tenantId,
			workflowRunId,
			selectedWorkerId,
			dispatcherId,
		),
	)

	if err != nil {
		return fmt.Errorf("could not add job assigned task to task queue: %w", err)
	}

	return nil
}

func (wc *WorkflowsControllerImpl) queueWorkflowRunJobs(ctx context.Context, workflowRun *dbsqlc.GetWorkflowRunRow, isPaused bool) error {
	ctx, span := telemetry.NewSpan(ctx, "queue-workflow-run-jobs") // nolint:ineffassign
	defer span.End()

	tenantId := sqlchelpers.UUIDToStr(workflowRun.WorkflowRun.TenantId)
	workflowRunId := sqlchelpers.UUIDToStr(workflowRun.WorkflowRun.ID)
	workflowId := sqlchelpers.UUIDToStr(workflowRun.WorkflowVersion.WorkflowId)

	if isPaused {
		return wc.repo.WorkflowRun().QueuePausedWorkflowRun(ctx, tenantId, workflowId, workflowRunId)
	}

	jobRuns, err := wc.repo.JobRun().ListJobRunsForWorkflowRun(ctx, tenantId, workflowRunId)

	if err != nil {
		return fmt.Errorf("could not list job runs: %w", err)
	}

	jobRunIds := make([]string, 0)

	for i := range jobRuns {
		// don't start job runs that are onFailure
		if workflowRun.WorkflowVersion.OnFailureJobId.Valid && jobRuns[i].JobId == workflowRun.WorkflowVersion.OnFailureJobId {
			continue
		}

		jobRunIds = append(jobRunIds, sqlchelpers.UUIDToStr(jobRuns[i].ID))
	}

	return wc.startManyJobRuns(ctx, tenantId, jobRunIds)
}

func (wc *WorkflowsControllerImpl) cancelWorkflowRunJobs(ctx context.Context, workflowRun *dbsqlc.GetWorkflowRunRow, reason string) error {
	ctx, span := telemetry.NewSpan(ctx, "cancel-workflow-run-jobs") // nolint:ineffassign
	defer span.End()

	tenantId := sqlchelpers.UUIDToStr(workflowRun.WorkflowRun.TenantId)
	workflowRunId := sqlchelpers.UUIDToStr(workflowRun.WorkflowRun.ID)

	jobRuns, err := wc.repo.JobRun().ListJobRunsForWorkflowRun(ctx, tenantId, workflowRunId)

	if err != nil {
		return fmt.Errorf("could not list job runs: %w", err)
	}

	var returnErr error

	for i := range jobRuns {
		// don't cancel job runs that are onFailure
		if workflowRun.WorkflowVersion.OnFailureJobId.Valid && jobRuns[i].JobId == workflowRun.WorkflowVersion.OnFailureJobId {
			continue
		}

		jobRunId := sqlchelpers.UUIDToStr(jobRuns[i].ID)

		err := wc.mq.AddMessage(
			context.Background(),
			msgqueue.JOB_PROCESSING_QUEUE,
			tasktypes.JobRunCancelledToTask(tenantId, jobRunId, &reason),
		)

		if err != nil {
			returnErr = multierror.Append(err, fmt.Errorf("could not add job run to task queue: %w", err))
		}
	}

	return returnErr
}

func (wc *WorkflowsControllerImpl) runGetGroupKeyRunRequeue(ctx context.Context) func() {
	return func() {
		wc.l.Debug().Msgf("workflows controller: checking get group key run requeue")

		// list all tenants
		tenants, err := wc.repo.Tenant().ListTenantsByControllerPartition(ctx, wc.p.GetControllerPartitionId())

		if err != nil {
			wc.l.Err(err).Msg("could not list tenants")
			return
		}

		g := new(errgroup.Group)

		for i := range tenants {
			tenantId := sqlchelpers.UUIDToStr(tenants[i].ID)

			g.Go(func() error {
				return wc.runGetGroupKeyRunRequeueTenant(ctx, tenantId)
			})
		}

		err = g.Wait()

		if err != nil {
			wc.l.Err(err).Msg("could not run get group key run requeue")
		}
	}
}

// runGetGroupKeyRunRequeueTenant looks for any get group key runs that haven't been assigned that are past their
// requeue time
func (ec *WorkflowsControllerImpl) runGetGroupKeyRunRequeueTenant(ctx context.Context, tenantId string) error {
	ctx, span := telemetry.NewSpan(ctx, "handle-get-group-key-run-requeue")
	defer span.End()

	getGroupKeyRuns, err := ec.repo.GetGroupKeyRun().ListGetGroupKeyRunsToRequeue(ctx, tenantId)

	if err != nil {
		return fmt.Errorf("could not list group key runs: %w", err)
	}

	g := new(errgroup.Group)

	for i := range getGroupKeyRuns {
		getGroupKeyRunCp := getGroupKeyRuns[i]

		// wrap in func to get defer on the span to avoid leaking spans
		g.Go(func() (err error) {
			var innerGetGroupKeyRun *dbsqlc.GetGroupKeyRunForEngineRow

			ctx, span := telemetry.NewSpan(ctx, "handle-get-group-key-run-requeue-tenant")
			defer span.End()

			getGroupKeyRunId := sqlchelpers.UUIDToStr(getGroupKeyRunCp.ID)

			ec.l.Debug().Msgf("requeueing group key run %s", getGroupKeyRunId)

			now := time.Now().UTC().UTC()

			// if the current time is after the scheduleTimeoutAt, then mark this as timed out
			scheduleTimeoutAt := getGroupKeyRunCp.ScheduleTimeoutAt.Time

			// timed out if there was no scheduleTimeoutAt set and the current time is after the get group key run created at time plus the default schedule timeout,
			// or if the scheduleTimeoutAt is set and the current time is after the scheduleTimeoutAt
			isTimedOut := !scheduleTimeoutAt.IsZero() && scheduleTimeoutAt.Before(now)

			if isTimedOut {
				return ec.cancelGetGroupKeyRun(ctx, tenantId, getGroupKeyRunId, "SCHEDULING_TIMED_OUT")
			}

			requeueAfter := time.Now().UTC().Add(time.Second * 4)

			innerGetGroupKeyRun, err = ec.repo.GetGroupKeyRun().UpdateGetGroupKeyRun(ctx, tenantId, getGroupKeyRunId, &repository.UpdateGetGroupKeyRunOpts{
				RequeueAfter: &requeueAfter,
			})

			if err != nil {
				return fmt.Errorf("could not update get group key run %s: %w", getGroupKeyRunId, err)
			}

			return ec.scheduleGetGroupAction(ctx, innerGetGroupKeyRun)
		})
	}

	return g.Wait()
}

func (wc *WorkflowsControllerImpl) runGetGroupKeyRunReassign(ctx context.Context) func() {
	return func() {
		wc.l.Debug().Msgf("workflows controller: checking get group key run reassign")

		// list all tenants
		tenants, err := wc.repo.Tenant().ListTenantsByControllerPartition(ctx, wc.p.GetControllerPartitionId())

		if err != nil {
			wc.l.Err(err).Msg("could not list tenants")
			return
		}

		g := new(errgroup.Group)

		for i := range tenants {
			tenantId := sqlchelpers.UUIDToStr(tenants[i].ID)

			g.Go(func() error {
				return wc.runGetGroupKeyRunReassignTenant(ctx, tenantId)
			})
		}

		err = g.Wait()

		if err != nil {
			wc.l.Err(err).Msg("could not run get group key run reassign")
		}
	}
}

// runGetGroupKeyRunReassignTenant looks for any get group key runs that have been assigned to an inactive worker
func (ec *WorkflowsControllerImpl) runGetGroupKeyRunReassignTenant(ctx context.Context, tenantId string) error {
	ctx, span := telemetry.NewSpan(ctx, "handle-get-group-key-run-reassign")
	defer span.End()

	getGroupKeyRuns, err := ec.repo.GetGroupKeyRun().ListGetGroupKeyRunsToReassign(ctx, tenantId)

	if err != nil {
		return fmt.Errorf("could not list get group key runs: %w", err)
	}

	g := new(errgroup.Group)

	for i := range getGroupKeyRuns {
		getGroupKeyRunCp := getGroupKeyRuns[i]

		// wrap in func to get defer on the span to avoid leaking spans
		g.Go(func() (err error) {
			var innerGetGroupKeyRun *dbsqlc.GetGroupKeyRunForEngineRow

			ctx, span := telemetry.NewSpan(ctx, "handle-get-group-key-run-reassign-tenant")
			defer span.End()

			getGroupKeyRunId := sqlchelpers.UUIDToStr(getGroupKeyRunCp.ID)

			ec.l.Debug().Msgf("reassigning group key run %s", getGroupKeyRunId)

			requeueAfter := time.Now().UTC().Add(time.Second * 4)

			innerGetGroupKeyRun, err = ec.repo.GetGroupKeyRun().UpdateGetGroupKeyRun(ctx, tenantId, getGroupKeyRunId, &repository.UpdateGetGroupKeyRunOpts{
				RequeueAfter: &requeueAfter,
				Status:       repository.StepRunStatusPtr(db.StepRunStatusPendingAssignment),
			})

			if err != nil {
				return fmt.Errorf("could not update get group key run %s: %w", getGroupKeyRunId, err)
			}

			return ec.scheduleGetGroupAction(ctx, innerGetGroupKeyRun)
		})
	}

	return g.Wait()
}

func (wc *WorkflowsControllerImpl) getLock(key string) *sync.Mutex {
	actual, _ := wc.queueMutex.LoadOrStore(key, &sync.Mutex{})
	return actual.(*sync.Mutex)
}

func (wc *WorkflowsControllerImpl) queueByCancelInProgress(ctx context.Context, tenantId, groupKey string, workflowVersion *dbsqlc.GetWorkflowVersionForEngineRow) error {
	ctx, span := telemetry.NewSpan(ctx, "queue-by-cancel-in-progress")
	defer span.End()

	wc.l.Info().Msgf("handling queue with strategy CANCEL_IN_PROGRESS for %s", groupKey)

	mutex := wc.getLock(fmt.Sprintf("%s:%s", tenantId, groupKey))

	if ok := mutex.TryLock(); !ok {
		return nil
	}

	defer mutex.Unlock()

	running := db.WorkflowRunStatusRunning
	queued := db.WorkflowRunStatusQueued
	workflowVersionId := sqlchelpers.UUIDToStr(workflowVersion.WorkflowVersion.ID)
	maxRuns := int(workflowVersion.ConcurrencyMaxRuns.Int32)

	runningWorkflowRuns, err := wc.repo.WorkflowRun().ListWorkflowRuns(ctx, tenantId, &repository.ListWorkflowRunsOpts{
		WorkflowVersionId: &workflowVersionId,
		GroupKey:          &groupKey,
		Statuses:          &[]db.WorkflowRunStatus{running},
		OrderBy:           repository.StringPtr("createdAt"),
		OrderDirection:    repository.StringPtr("ASC"),
	})
	if err != nil {
		return fmt.Errorf("could not list running workflow runs: %w", err)
	}

	queuedWorkflowRuns, err := wc.repo.WorkflowRun().ListWorkflowRuns(ctx, tenantId, &repository.ListWorkflowRunsOpts{
		WorkflowVersionId: &workflowVersionId,
		GroupKey:          &groupKey,
		Statuses:          &[]db.WorkflowRunStatus{queued},
		OrderBy:           repository.StringPtr("createdAt"),
		OrderDirection:    repository.StringPtr("ASC"),
	})
	if err != nil {
		return fmt.Errorf("could not list queued workflow runs: %w", err)
	}

	runningCount := len(runningWorkflowRuns.Rows)
	queuedCount := len(queuedWorkflowRuns.Rows)

	// Calculate how many runs we need to cancel
	toCancel := max(0, runningCount+queuedCount-maxRuns)

	errGroup := new(errgroup.Group)

	// Cancel the oldest running workflows
	for i := 0; i < toCancel && i < runningCount; i++ {
		row := runningWorkflowRuns.Rows[i]
		workflowRunId := sqlchelpers.UUIDToStr(row.WorkflowRun.ID)

		errGroup.Go(func() error {
			return wc.cancelWorkflowRun(ctx, tenantId, workflowRunId)
		})

		errGroup.Go(func() error {
			workflowRun, err := wc.repo.WorkflowRun().GetWorkflowRunById(ctx, tenantId, workflowRunId)
			if err != nil {
				return err
			}
			return wc.cancelWorkflowRunJobs(ctx, workflowRun, "CANCEL_IN_PROGRESS")
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("could not cancel workflow runs: %w", err)
	}

	// Queue new runs
	toQueue := min(maxRuns-(runningCount-toCancel), queuedCount)
	errGroup = new(errgroup.Group)

	for i := 0; i < toQueue; i++ {
		row := queuedWorkflowRuns.Rows[i]
		workflowRunId := sqlchelpers.UUIDToStr(row.WorkflowRun.ID)

		errGroup.Go(func() error {
			workflowRun, err := wc.repo.WorkflowRun().GetWorkflowRunById(ctx, tenantId, workflowRunId)
			if err != nil {
				return fmt.Errorf("could not get workflow run: %w", err)
			}

			isPaused := workflowRun.IsPaused.Valid && workflowRun.IsPaused.Bool
			if isPaused {
				return nil
			}

			return wc.queueWorkflowRunJobs(ctx, workflowRun, isPaused)
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("could not queue workflow runs: %w", err)
	}

	return nil
}

func (wc *WorkflowsControllerImpl) queueByGroupRoundRobin(ctx context.Context, tenantId string, workflowVersion *dbsqlc.GetWorkflowVersionForEngineRow) error {
	ctx, span := telemetry.NewSpan(ctx, "queue-by-group-round-robin")
	defer span.End()

	workflowVersionId := sqlchelpers.UUIDToStr(workflowVersion.WorkflowVersion.ID)
	workflowId := sqlchelpers.UUIDToStr(workflowVersion.WorkflowVersion.WorkflowId)
	maxRuns := int(workflowVersion.ConcurrencyMaxRuns.Int32)

	wc.l.Info().Msgf("handling queue with strategy GROUP_ROUND_ROBIN for workflow version %s", workflowVersionId)

	// get workflow runs which are queued for this group key
	poppedWorkflowRuns, err := wc.repo.WorkflowRun().PopWorkflowRunsRoundRobin(ctx, tenantId, workflowId, maxRuns)

	if err != nil {
		return fmt.Errorf("could not list queued workflow runs: %w", err)
	}

	errGroup := new(errgroup.Group)

	for i := range poppedWorkflowRuns {
		row := poppedWorkflowRuns[i]

		errGroup.Go(func() error {
			workflowRunId := sqlchelpers.UUIDToStr(row.ID)

			wc.l.Info().Msgf("popped workflow run %s", workflowRunId)
			workflowRun, err := wc.repo.WorkflowRun().GetWorkflowRunById(ctx, tenantId, workflowRunId)

			if err != nil {
				return fmt.Errorf("could not get workflow run: %w", err)
			}

			isPaused := workflowRun.IsPaused.Valid && workflowRun.IsPaused.Bool

			if isPaused {
				return nil
			}

			return wc.queueWorkflowRunJobs(ctx, workflowRun, isPaused)
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("could not queue workflow runs: %w", err)
	}

	return nil
}

func (wc *WorkflowsControllerImpl) cancelWorkflowRun(ctx context.Context, tenantId, workflowRunId string) error {
	// cancel all running step runs
	runningStatus := dbsqlc.StepRunStatusRUNNING
	stepRuns, err := wc.repo.StepRun().ListStepRuns(ctx, tenantId, &repository.ListStepRunsOpts{
		Status: &runningStatus,
		WorkflowRunIds: []string{
			workflowRunId,
		},
	})

	if err != nil {
		return fmt.Errorf("could not list step runs: %w", err)
	}

	errGroup := new(errgroup.Group)

	for i := range stepRuns {
		stepRunCp := stepRuns[i]
		stepRunId := sqlchelpers.UUIDToStr(stepRunCp.SRID)

		errGroup.Go(func() error {
			return wc.mq.AddMessage(
				context.Background(),
				msgqueue.JOB_PROCESSING_QUEUE,
				getStepRunCancelTask(tenantId, stepRunId, "CANCELLED_BY_CONCURRENCY_LIMIT"),
			)
		})
	}

	return errGroup.Wait()
}

func getGroupActionTask(tenantId, workflowRunId, workerId, dispatcherId string) *msgqueue.Message {
	payload, _ := datautils.ToJSONMap(tasktypes.GroupKeyActionAssignedTaskPayload{
		WorkflowRunId: workflowRunId,
		WorkerId:      workerId,
	})

	metadata, _ := datautils.ToJSONMap(tasktypes.GroupKeyActionAssignedTaskMetadata{
		TenantId:     tenantId,
		DispatcherId: dispatcherId,
	})

	return &msgqueue.Message{
		ID:       "group-key-action-assigned",
		Payload:  payload,
		Metadata: metadata,
		Retries:  3,
	}
}

func getStepRunCancelTask(tenantId, stepRunId, reason string) *msgqueue.Message {
	payload, _ := datautils.ToJSONMap(tasktypes.StepRunCancelTaskPayload{
		StepRunId:       stepRunId,
		CancelledReason: reason,
	})

	metadata, _ := datautils.ToJSONMap(tasktypes.StepRunCancelTaskMetadata{
		TenantId: tenantId,
	})

	return &msgqueue.Message{
		ID:       "step-run-cancel",
		Payload:  payload,
		Metadata: metadata,
		Retries:  3,
	}
}
