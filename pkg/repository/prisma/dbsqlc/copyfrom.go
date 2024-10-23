// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: copyfrom.go

package dbsqlc

import (
	"context"
)

// iteratorForCreateEvents implements pgx.CopyFromSource.
type iteratorForCreateEvents struct {
	rows                 []CreateEventsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateEvents) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateEvents) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].Key,
		r.rows[0].TenantId,
		r.rows[0].ReplayedFromId,
		r.rows[0].Data,
		r.rows[0].AdditionalMetadata,
		r.rows[0].InsertOrder,
	}, nil
}

func (r iteratorForCreateEvents) Err() error {
	return nil
}

func (q *Queries) CreateEvents(ctx context.Context, db DBTX, arg []CreateEventsParams) (int64, error) {
	return db.CopyFrom(ctx, []string{"Event"}, []string{"id", "key", "tenantId", "replayedFromId", "data", "additionalMetadata", "insertOrder"}, &iteratorForCreateEvents{rows: arg})
}

// iteratorForCreateGetGroupKeyRuns implements pgx.CopyFromSource.
type iteratorForCreateGetGroupKeyRuns struct {
	rows                 []CreateGetGroupKeyRunsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateGetGroupKeyRuns) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateGetGroupKeyRuns) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].TenantId,
		r.rows[0].WorkflowRunId,
		r.rows[0].Input,
		r.rows[0].RequeueAfter,
		r.rows[0].ScheduleTimeoutAt,
		r.rows[0].Status,
	}, nil
}

func (r iteratorForCreateGetGroupKeyRuns) Err() error {
	return nil
}

func (q *Queries) CreateGetGroupKeyRuns(ctx context.Context, db DBTX, arg []CreateGetGroupKeyRunsParams) (int64, error) {
	return db.CopyFrom(ctx, []string{"GetGroupKeyRun"}, []string{"id", "tenantId", "workflowRunId", "input", "requeueAfter", "scheduleTimeoutAt", "status"}, &iteratorForCreateGetGroupKeyRuns{rows: arg})
}

// iteratorForCreateQueueItemsBulk implements pgx.CopyFromSource.
type iteratorForCreateQueueItemsBulk struct {
	rows                 []CreateQueueItemsBulkParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateQueueItemsBulk) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateQueueItemsBulk) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].StepRunId,
		r.rows[0].StepId,
		r.rows[0].ActionId,
		r.rows[0].ScheduleTimeoutAt,
		r.rows[0].StepTimeout,
		r.rows[0].Priority,
		r.rows[0].IsQueued,
		r.rows[0].TenantId,
		r.rows[0].Queue,
		r.rows[0].Sticky,
		r.rows[0].DesiredWorkerId,
	}, nil
}

func (r iteratorForCreateQueueItemsBulk) Err() error {
	return nil
}

func (q *Queries) CreateQueueItemsBulk(ctx context.Context, db DBTX, arg []CreateQueueItemsBulkParams) (int64, error) {
	return db.CopyFrom(ctx, []string{"QueueItem"}, []string{"stepRunId", "stepId", "actionId", "scheduleTimeoutAt", "stepTimeout", "priority", "isQueued", "tenantId", "queue", "sticky", "desiredWorkerId"}, &iteratorForCreateQueueItemsBulk{rows: arg})
}

// iteratorForCreateStepRuns implements pgx.CopyFromSource.
type iteratorForCreateStepRuns struct {
	rows                 []CreateStepRunsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateStepRuns) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateStepRuns) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].TenantId,
		r.rows[0].JobRunId,
		r.rows[0].StepId,
		r.rows[0].Status,
		r.rows[0].RequeueAfter,
		r.rows[0].Queue,
		r.rows[0].Priority,
	}, nil
}

func (r iteratorForCreateStepRuns) Err() error {
	return nil
}

func (q *Queries) CreateStepRuns(ctx context.Context, db DBTX, arg []CreateStepRunsParams) (int64, error) {
	return db.CopyFrom(ctx, []string{"StepRun"}, []string{"id", "tenantId", "jobRunId", "stepId", "status", "requeueAfter", "queue", "priority"}, &iteratorForCreateStepRuns{rows: arg})
}

// iteratorForCreateWorkflowRunTriggeredBys implements pgx.CopyFromSource.
type iteratorForCreateWorkflowRunTriggeredBys struct {
	rows                 []CreateWorkflowRunTriggeredBysParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateWorkflowRunTriggeredBys) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateWorkflowRunTriggeredBys) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].TenantId,
		r.rows[0].ParentId,
		r.rows[0].EventId,
		r.rows[0].CronParentId,
		r.rows[0].CronSchedule,
		r.rows[0].ScheduledId,
	}, nil
}

func (r iteratorForCreateWorkflowRunTriggeredBys) Err() error {
	return nil
}

func (q *Queries) CreateWorkflowRunTriggeredBys(ctx context.Context, db DBTX, arg []CreateWorkflowRunTriggeredBysParams) (int64, error) {
	return db.CopyFrom(ctx, []string{"WorkflowRunTriggeredBy"}, []string{"id", "tenantId", "parentId", "eventId", "cronParentId", "cronSchedule", "scheduledId"}, &iteratorForCreateWorkflowRunTriggeredBys{rows: arg})
}

// iteratorForCreateWorkflowRuns implements pgx.CopyFromSource.
type iteratorForCreateWorkflowRuns struct {
	rows                 []CreateWorkflowRunsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateWorkflowRuns) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateWorkflowRuns) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].DisplayName,
		r.rows[0].TenantId,
		r.rows[0].WorkflowVersionId,
		r.rows[0].Status,
		r.rows[0].ChildIndex,
		r.rows[0].ChildKey,
		r.rows[0].ParentId,
		r.rows[0].ParentStepRunId,
		r.rows[0].AdditionalMetadata,
		r.rows[0].Priority,
		r.rows[0].InsertOrder,
	}, nil
}

func (r iteratorForCreateWorkflowRuns) Err() error {
	return nil
}

func (q *Queries) CreateWorkflowRuns(ctx context.Context, db DBTX, arg []CreateWorkflowRunsParams) (int64, error) {
	return db.CopyFrom(ctx, []string{"WorkflowRun"}, []string{"id", "displayName", "tenantId", "workflowVersionId", "status", "childIndex", "childKey", "parentId", "parentStepRunId", "additionalMetadata", "priority", "insertOrder"}, &iteratorForCreateWorkflowRuns{rows: arg})
}
