import { Input } from '@/components/v1/ui/input';
import { Label } from '@/components/v1/ui/label';
import { WorkflowVersion } from '@/lib/api';
import { useWorkflowDetails } from '@/next/hooks/use-workflow-details';
import CronPrettifier from 'cronstrue';

export default function WorkflowGeneralSettings() {
  const { workflowVersion } = useWorkflowDetails();

  return (
    <>
      <h3 className="text-lg font-semibold mb-4">Trigger</h3>
      <TriggerSettings workflow={workflowVersion} />
      <h3 className="text-lg font-semibold my-4">Concurrency</h3>
      <ConcurrencySettings workflow={workflowVersion} />
      <h3 className="text-lg font-semibold my-4">Sticky Strategy</h3>
      <StickyStrategy workflow={workflowVersion} />
      <h3 className="text-lg font-semibold my-4">Default Priority</h3>
      <DefaultPriority workflow={workflowVersion} />
    </>
  );
}

function TriggerSettings({
  workflow,
}: {
  workflow: WorkflowVersion | undefined;
}) {
  if (!workflow?.triggers || Object.keys(workflow.triggers).length === 0) {
    return (
      <div className="text-[0.8rem] text-gray-700 dark:text-gray-300">
        There are no trigger settings for this workflow.
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-2">
      {workflow.triggers.events ? (
        <>
          <Label>Events</Label>
          {workflow.triggers.events.map((event) => (
            <Input
              key={event.event_key}
              disabled
              placeholder="shadcn"
              value={event.event_key}
            />
          ))}
        </>
      ) : null}
      {workflow.triggers.crons ? (
        <>
          <Label>Crons</Label>
          {workflow.triggers.crons.map((event) => (
            <>
              <Input
                key={event.cron}
                disabled
                placeholder="shadcn"
                value={event.cron}
              />
              {event.cron ? (
                <span className="text-sm mb-2 text-gray-500">
                  (runs {CronPrettifier.toString(event.cron).toLowerCase()} UTC)
                </span>
              ) : null}
            </>
          ))}
        </>
      ) : null}
    </div>
  );
}

function ConcurrencySettings({
  workflow,
}: {
  workflow: WorkflowVersion | undefined;
}) {
  if (!workflow?.concurrency) {
    return (
      <div className="text-[0.8rem] text-gray-700 dark:text-gray-300">
        There are no concurrency settings for this workflow.
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-2">
      <Label>Max runs</Label>
      <Input
        disabled
        placeholder="shadcn"
        value={workflow.concurrency?.maxRuns}
      />
      <div className="text-[0.8rem] text-gray-700 dark:text-gray-300">
        The maximum number of concurrency runs.
      </div>
      <Label className="mt-4">Concurrency strategy</Label>
      <Input
        disabled
        placeholder="shadcn"
        value={workflow.concurrency?.limitStrategy}
      />
      <div className="text-[0.8rem] text-gray-700 dark:text-gray-300">
        The strategy to use when the maximum number of concurrency runs is
        reached.
      </div>
    </div>
  );
}

function StickyStrategy({
  workflow,
}: {
  workflow: WorkflowVersion | undefined;
}) {
  if (!workflow?.sticky) {
    return (
      <div className="text-[0.8rem] text-gray-700 dark:text-gray-300">
        There is no sticky strategy set for this workflow.
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-2">
      <Label>Strategy</Label>
      <Input disabled value={workflow.sticky} />
    </div>
  );
}

function DefaultPriority({
  workflow,
}: {
  workflow: WorkflowVersion | undefined;
}) {
  if (!workflow?.defaultPriority) {
    return (
      <div className="text-[0.8rem] text-gray-700 dark:text-gray-300">
        There is no default priority set for this workflow.
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-2">
      <Label>Default Priority</Label>
      <Input disabled value={workflow.defaultPriority} />
    </div>
  );
}
