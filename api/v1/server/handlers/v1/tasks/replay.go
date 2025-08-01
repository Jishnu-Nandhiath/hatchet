package tasks

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hatchet-dev/hatchet/api/v1/server/oas/gen"
	"github.com/hatchet-dev/hatchet/api/v1/server/oas/transformers/v1"
	contracts "github.com/hatchet-dev/hatchet/internal/services/shared/proto/v1"
	"github.com/hatchet-dev/hatchet/pkg/repository/postgres/dbsqlc"
)

func (t *TasksService) V1TaskReplay(ctx echo.Context, request gen.V1TaskReplayRequestObject) (gen.V1TaskReplayResponseObject, error) {
	tenant := ctx.Get("tenant").(*dbsqlc.Tenant)

	var err error

	grpcReq := &contracts.ReplayTasksRequest{}

	if request.Body.ExternalIds != nil {
		externalIds := make([]string, 0)

		for _, id := range *request.Body.ExternalIds {
			externalIds = append(externalIds, id.String())
		}

		grpcReq.ExternalIds = externalIds
	}

	if request.Body.Filter != nil {
		filter := &contracts.TasksFilter{
			Since: timestamppb.New(request.Body.Filter.Since),
		}

		if request.Body.Filter.Until != nil {
			filter.Until = timestamppb.New(*request.Body.Filter.Until)
		}

		if request.Body.Filter.Statuses != nil {
			filter.Statuses = make([]string, len(*request.Body.Filter.Statuses))

			for i, status := range *request.Body.Filter.Statuses {
				filter.Statuses[i] = string(status)
			}
		}

		if request.Body.Filter.WorkflowIds != nil {
			filter.WorkflowIds = make([]string, len(*request.Body.Filter.WorkflowIds))

			for i, id := range *request.Body.Filter.WorkflowIds {
				filter.WorkflowIds[i] = id.String()
			}
		}

		if request.Body.Filter.AdditionalMetadata != nil {
			filter.AdditionalMetadata = make([]string, len(*request.Body.Filter.AdditionalMetadata))

			copy(filter.AdditionalMetadata, *request.Body.Filter.AdditionalMetadata)
		}

		grpcReq.Filter = filter
	}

	resp, err := t.proxyReplay.Do(
		ctx.Request().Context(),
		tenant,
		grpcReq,
	)

	if err != nil {
		return nil, err
	}

	ids := transformers.ToCancelledOrReplayedTaskResponse(resp.ReplayedTasks)

	return gen.V1TaskReplay200JSONResponse(
		ids,
	), nil
}
