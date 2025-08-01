import { DataTable } from '@/components/v1/molecules/data-table/data-table.tsx';
import { columns } from './v1/task-runs-columns';
import { useCallback, useMemo, useState } from 'react';
import { RowSelectionState, VisibilityState } from '@tanstack/react-table';
import { useQuery } from '@tanstack/react-query';
import { queries } from '@/lib/api';
import { Button } from '@/components/v1/ui/button';
import { ArrowPathIcon, XCircleIcon } from '@heroicons/react/24/outline';
import { V1WorkflowRunsMetricsView } from './task-runs-metrics';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/v1/ui/select';
import { Skeleton } from '@/components/v1/ui/skeleton';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from '@/components/v1/ui/dialog';
import { CodeHighlighter } from '@/components/v1/ui/code-highlighter';
import { Separator } from '@/components/v1/ui/separator';
import {
  DataPoint,
  ZoomableChart,
} from '@/components/v1/molecules/charts/zoomable';
import { DateTimePicker } from '@/components/v1/molecules/time-picker/date-time-picker';
import { Sheet, SheetContent } from '@/components/v1/ui/sheet';
import {
  TabOption,
  TaskRunDetail,
} from '../$run/v2components/step-run-detail/step-run-detail';
import { TaskRunActionButton } from '../../task-runs-v1/actions';
import { TimeWindow, useColumnFilters } from '../hooks/column-filters';
import { usePagination } from '../hooks/pagination';
import { useTaskRuns } from '../hooks/task-runs';
import { useMetrics } from '../hooks/metrics';
import { useToolbarFilters } from '../hooks/toolbar-filters';
import { IntroDocsEmptyState } from '@/pages/onboarding/intro-docs-empty-state';
import { useCurrentTenantId } from '@/hooks/use-tenant';
import { TriggerWorkflowForm } from '../../workflows/$workflow/components/trigger-workflow-form';
import { useToast } from '@/components/v1/hooks/use-toast';
import { Toaster } from '@/components/v1/ui/toaster';

export interface TaskRunsTableProps {
  createdAfter?: string;
  createdBefore?: string;
  workflowId?: string;
  workerId?: string;
  initColumnVisibility?: VisibilityState;
  filterVisibility?: { [key: string]: boolean };
  refetchInterval?: number;
  showMetrics?: boolean;
  showCounts?: boolean;
  showDateFilter?: boolean;
  showTriggerRunButton?: boolean;
  parentTaskExternalId?: string;
  triggeringEventExternalId?: string;
  disableTaskRunPagination?: boolean;
  headerClassName?: string;
}

type StepDetailSheetState = {
  isOpen: boolean;
  taskRunId: string | undefined;
};

export function TaskRunsTable({
  workflowId,
  workerId,
  parentTaskExternalId,
  triggeringEventExternalId,
  createdAfter: createdAfterProp,
  initColumnVisibility = {},
  filterVisibility = {},
  refetchInterval = 5000,
  showMetrics = false,
  showCounts = true,
  showDateFilter = true,
  disableTaskRunPagination = false,
  showTriggerRunButton = true,
  headerClassName,
}: TaskRunsTableProps) {
  const { tenantId } = useCurrentTenantId();
  const { toast } = useToast();

  const [selectedAdditionalMetaRunId, setSelectedAdditionalMetaRunId] =
    useState<string | null>(null);
  const [triggerWorkflow, setTriggerWorkflow] = useState(false);
  const [viewQueueMetrics, setViewQueueMetrics] = useState(false);
  const [rotate, setRotate] = useState(false);
  const [rowSelection, setRowSelection] = useState<RowSelectionState>({});
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({
    // IMPORTANT: the parentTaskExternalId column is hidden by default and shouldn't be shown
    // It's here for filtering
    ...initColumnVisibility,
    parentTaskExternalId: false,
  });
  const [stepDetailSheetState, setStepDetailSheetState] =
    useState<StepDetailSheetState>({
      isOpen: false,
      taskRunId: undefined,
    });

  const [taskIdsPendingAction, setTaskIdsPendingAction] = useState<string[]>(
    [],
  );

  const cf = useColumnFilters();

  const toolbarFilters = useToolbarFilters({ filterVisibility });
  const { pagination, setPagination, setPageSize } = usePagination();

  const workflow = workflowId || cf.filters.workflowId;
  const derivedParentTaskExternalId =
    parentTaskExternalId || cf.filters.parentTaskExternalId;

  const hasOpenUI =
    !!selectedAdditionalMetaRunId || stepDetailSheetState.isOpen;

  const {
    tableRows,
    selectedRuns,
    numPages,
    isLoading: isTaskRunsLoading,
    isFetching: isTaskRunsFetching,
    refetch: refetchTaskRuns,
    getRowId,
  } = useTaskRuns({
    rowSelection,
    workerId,
    workflow,
    parentTaskExternalId: derivedParentTaskExternalId,
    triggeringEventExternalId: triggeringEventExternalId,
    disablePagination: disableTaskRunPagination,
    pauseRefetch: hasOpenUI,
  });

  const {
    metrics,
    tenantMetrics,
    isLoading: isMetricsLoading,
    isFetching: isMetricsFetching,
    refetch: refetchMetrics,
  } = useMetrics({
    workflow,
    refetchInterval,
    parentTaskExternalId: derivedParentTaskExternalId,
    pauseRefetch: hasOpenUI,
  });

  const onTaskRunIdClick = useCallback((taskRunId: string) => {
    setStepDetailSheetState({
      taskRunId,
      isOpen: true,
    });
  }, []);

  const parentTaskRun = useQuery({
    ...queries.v1Tasks.get(derivedParentTaskExternalId || ''),
    enabled: !!derivedParentTaskExternalId,
  });

  const v1TaskFilters = useMemo(
    () => ({
      since: cf.filters.createdAfter,
      until: cf.filters.finishedBefore,
      statuses: cf.filters.status ? [cf.filters.status] : undefined,
      workflowIds: workflow ? [workflow] : undefined,
      additionalMetadata: cf.filters.additionalMetadata,
    }),
    [
      cf.filters.createdAfter,
      cf.filters.finishedBefore,
      cf.filters.status,
      workflow,
      cf.filters.additionalMetadata,
    ],
  );

  const hasRowsSelected = Object.values(rowSelection).some(
    (selected) => !!selected,
  );
  const hasTaskFiltersSelected = Object.values(v1TaskFilters).some(
    (filter) => !!filter,
  );

  const hasLoaded = useMemo(() => {
    return !isTaskRunsLoading && !isMetricsLoading;
  }, [isTaskRunsLoading, isMetricsLoading]);

  const isFetching = !hasLoaded && (isTaskRunsFetching || isMetricsFetching);

  const onActionProcessed = useCallback(
    (action: 'cancel' | 'replay', ids: string[]) => {
      const prefix = action === 'cancel' ? 'Canceling' : 'Replaying';
      const count = ids.length;

      setTaskIdsPendingAction(ids);
      const t = toast({
        title: `${prefix} ${count} task run${count > 1 ? 's' : ''}`,
        description: `This may take a few seconds. You don't need to hit ${action} again.`,
      });

      setTimeout(() => {
        setTaskIdsPendingAction([]);
        t.dismiss();
      }, 5000);
    },
    [toast],
  );

  const actions = useMemo(() => {
    let localActions = [
      <TaskRunActionButton
        key="cancel"
        actionType="cancel"
        disabled={
          !(hasRowsSelected || hasTaskFiltersSelected) ||
          taskIdsPendingAction.length > 0
        }
        params={
          selectedRuns.length > 0
            ? { externalIds: selectedRuns.map((run) => run?.metadata.id) }
            : { filter: v1TaskFilters }
        }
        showModal
        onActionProcessed={(ids) => onActionProcessed('cancel', ids)}
        onActionSubmit={() => {
          toast({
            title: 'Cancel request submitted',
            description: "No need to hit 'Cancel' again.",
          });
        }}
      />,
      <TaskRunActionButton
        key="replay"
        actionType="replay"
        disabled={
          !(hasRowsSelected || hasTaskFiltersSelected) ||
          taskIdsPendingAction.length > 0
        }
        params={
          selectedRuns.length > 0
            ? { externalIds: selectedRuns.map((run) => run?.metadata.id) }
            : { filter: v1TaskFilters }
        }
        showModal
        onActionProcessed={(ids) => onActionProcessed('replay', ids)}
        onActionSubmit={() => {
          toast({
            title: 'Replay request submitted',
            description: "No need to hit 'Replay' again.",
          });
        }}
      />,
      <Button
        key="refresh"
        className="h-8 px-2 lg:px-3"
        size="sm"
        onClick={() => {
          refetchTaskRuns();
          refetchMetrics();
          setRotate(!rotate);
        }}
        variant={'outline'}
        aria-label="Refresh events list"
      >
        <ArrowPathIcon
          className={`h-4 w-4 transition-transform ${rotate ? 'rotate-180' : ''}`}
        />
      </Button>,
    ];

    if (showTriggerRunButton) {
      localActions = [
        <Button
          key="trigger"
          className="h-8 border"
          onClick={() => setTriggerWorkflow(true)}
        >
          Trigger Run
        </Button>,
        ...localActions,
      ];
    }
    return localActions;
  }, [
    showTriggerRunButton,
    hasRowsSelected,
    hasTaskFiltersSelected,
    selectedRuns,
    v1TaskFilters,
    refetchTaskRuns,
    refetchMetrics,
    rotate,
    onActionProcessed,
    taskIdsPendingAction.length,
    toast,
  ]);

  const handleSetSelectedAdditionalMetaRunId = useCallback(
    (runId: string | null) => {
      setSelectedAdditionalMetaRunId(runId);
    },
    [],
  );

  return (
    <div className="flex flex-col h-full overflow-hidden">
      <Toaster />
      <TriggerWorkflowForm
        defaultWorkflow={undefined}
        show={triggerWorkflow}
        onClose={() => setTriggerWorkflow(false)}
      />
      {cf.filters.parentTaskExternalId &&
        !parentTaskRun.isLoading &&
        parentTaskRun.data && (
          <div className="flex flex-row items-center gap-x-2">
            <p>Child runs of parent:</p>
            <p className="font-semibold text-orange-300">
              {' '}
              {parentTaskRun.data.displayName}
            </p>
            <Button
              variant="outline"
              className="ml-4"
              onClick={() => {
                cf.clearParentTaskExternalId();
              }}
            >
              Clear
            </Button>
          </div>
        )}
      {showMetrics && !derivedParentTaskExternalId && (
        <Dialog
          open={viewQueueMetrics}
          onOpenChange={(open) => {
            if (!open) {
              setViewQueueMetrics(false);
            }
          }}
        >
          <DialogContent className="w-fit max-w-[80%] min-w-[500px]">
            <DialogHeader>
              <DialogTitle>Queue Metrics</DialogTitle>
            </DialogHeader>
            <Separator />
            {tenantMetrics && (
              <CodeHighlighter
                language="json"
                className="max-h-[400px] overflow-y-auto"
                code={JSON.stringify(tenantMetrics || '{}', null, 2)}
              />
            )}
            {isMetricsLoading && 'Loading...'}
          </DialogContent>
        </Dialog>
      )}
      {showDateFilter && !createdAfterProp && !derivedParentTaskExternalId && (
        <div className="flex flex-row justify-end items-center mb-4 gap-2">
          {cf.filters.isCustomTimeRange && [
            <Button
              key="clear"
              onClick={() => {
                cf.setCustomTimeRange(undefined);
              }}
              variant="outline"
              size="sm"
              className="text-xs h-9 py-2"
            >
              <XCircleIcon className="h-[18px] w-[18px] mr-2" />
              Clear
            </Button>,
            <DateTimePicker
              key="after"
              label="After"
              date={
                cf.filters.createdAfter
                  ? new Date(cf.filters.createdAfter)
                  : undefined
              }
              setDate={(date) => {
                cf.setCreatedAfter(date?.toISOString());
              }}
            />,
            <DateTimePicker
              key="before"
              label="Before"
              date={
                cf.filters.finishedBefore
                  ? new Date(cf.filters.finishedBefore)
                  : undefined
              }
              setDate={(date) => {
                cf.setFinishedBefore(date?.toISOString());
              }}
            />,
          ]}
          <Select
            value={
              cf.filters.isCustomTimeRange ? 'custom' : cf.filters.timeWindow
            }
            onValueChange={(value: TimeWindow | 'custom') => {
              if (value !== 'custom') {
                cf.setFilterValues([
                  { key: 'isCustomTimeRange', value: false },
                  { key: 'timeWindow', value: value },
                ]);
              } else {
                cf.setFilterValues([{ key: 'isCustomTimeRange', value: true }]);
              }
            }}
          >
            <SelectTrigger className="w-fit">
              <SelectValue id="timerange" placeholder="Choose time range" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="1h">1 hour</SelectItem>
              <SelectItem value="6h">6 hours</SelectItem>
              <SelectItem value="1d">1 day</SelectItem>
              <SelectItem value="7d">7 days</SelectItem>
              <SelectItem value="custom">Custom</SelectItem>
            </SelectContent>
          </Select>
        </div>
      )}
      {showMetrics && !derivedParentTaskExternalId && (
        <GetWorkflowChart
          createdAfter={cf.filters.createdAfter}
          zoom={(createdAfter, createdBefore) => {
            cf.setCustomTimeRange({ start: createdAfter, end: createdBefore });
          }}
          finishedBefore={cf.filters.finishedBefore}
          refetchInterval={refetchInterval}
          pauseRefetch={hasOpenUI}
        />
      )}
      {showCounts && (
        <div className="flex flex-row justify-between items-center my-4">
          {metrics.length > 0 ? (
            <V1WorkflowRunsMetricsView
              metrics={metrics}
              onViewQueueMetricsClick={() => {
                setViewQueueMetrics(true);
              }}
              showQueueMetrics={showMetrics}
              onClick={(status) => {
                cf.setStatus(status);
              }}
            />
          ) : (
            <Skeleton className="max-w-[800px] w-[40vw] h-8" />
          )}
        </div>
      )}
      {stepDetailSheetState.taskRunId && (
        <Sheet
          open={stepDetailSheetState.isOpen}
          onOpenChange={(isOpen) =>
            setStepDetailSheetState((prev) => ({
              ...prev,
              isOpen,
            }))
          }
        >
          <SheetContent className="w-fit min-w-[56rem] max-w-4xl sm:max-w-2xl z-[60] h-full overflow-auto">
            <TaskRunDetail
              taskRunId={stepDetailSheetState.taskRunId}
              defaultOpenTab={TabOption.Output}
              showViewTaskRunButton
            />
          </SheetContent>
        </Sheet>
      )}
      <div className="flex-1 min-h-0">
        <DataTable
          emptyState={
            <IntroDocsEmptyState
              link="/home/your-first-task"
              title="No Runs Found"
              linkPreambleText="To learn more about how workflows function in Hatchet,"
              linkText="check out our documentation."
            />
          }
          isLoading={isFetching}
          columns={columns(
            tenantId,
            selectedAdditionalMetaRunId,
            handleSetSelectedAdditionalMetaRunId,
            cf.setAdditionalMetadata,
            onTaskRunIdClick,
          )}
          columnVisibility={columnVisibility}
          setColumnVisibility={setColumnVisibility}
          data={tableRows}
          filters={toolbarFilters}
          actions={actions}
          columnFilters={cf.filters.columnFilters}
          setColumnFilters={(updaterOrValue) => {
            cf.setColumnFilters(updaterOrValue);
          }}
          pagination={pagination}
          setPagination={setPagination}
          onSetPageSize={setPageSize}
          rowSelection={rowSelection}
          setRowSelection={setRowSelection}
          pageCount={numPages}
          showColumnToggle={true}
          getSubRows={(row) => row.children || []}
          getRowId={getRowId}
          onToolbarReset={cf.clearColumnFilters}
          headerClassName={headerClassName}
        />
      </div>
    </div>
  );
}

const GetWorkflowChart = ({
  createdAfter,
  finishedBefore,
  refetchInterval,
  zoom,
  pauseRefetch = false,
}: {
  createdAfter?: string;
  finishedBefore?: string;
  refetchInterval?: number;
  zoom: (startTime: string, endTime: string) => void;
  pauseRefetch?: boolean;
}) => {
  const { tenantId } = useCurrentTenantId();
  const workflowRunEventsMetricsQuery = useQuery({
    ...queries.v1TaskRuns.pointMetrics(tenantId, {
      createdAfter,
      finishedBefore,
    }),
    placeholderData: (prev) => prev,
    refetchInterval: pauseRefetch ? false : refetchInterval,
  });

  if (workflowRunEventsMetricsQuery.isLoading) {
    return <Skeleton className="w-full h-36" />;
  }

  return (
    <ZoomableChart
      kind="bar"
      data={
        workflowRunEventsMetricsQuery.data?.results?.map(
          (result): DataPoint<'SUCCEEDED' | 'FAILED'> => ({
            date: result.time,
            SUCCEEDED: result.SUCCEEDED,
            FAILED: result.FAILED,
          }),
        ) || []
      }
      colors={{
        SUCCEEDED: 'rgb(34 197 94 / 0.5)',
        FAILED: 'hsl(var(--destructive))',
      }}
      zoom={zoom}
      showYAxis={false}
    />
  );
};
