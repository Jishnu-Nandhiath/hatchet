import { Channel, ClientFactory } from 'nice-grpc';
import {
  BulkPushEventRequest,
  EventsServiceClient,
  EventsServiceDefinition,
  PushEventRequest,
} from '@hatchet/protoc/events/events';
import HatchetError from '@util/errors/hatchet-error';
import { ClientConfig } from '@clients/hatchet-client/client-config';
import { Logger } from '@hatchet/util/logger';
import { retrier } from '@hatchet/util/retrier';
import { applyNamespace } from '@hatchet/util/apply-namespace';
import { HatchetClient } from '@hatchet/v1';
import { LegacyHatchetClient } from '../hatchet-client';

// eslint-disable-next-line no-shadow
export enum LogLevel {
  INFO = 'INFO',
  WARN = 'WARN',
  ERROR = 'ERROR',
  DEBUG = 'DEBUG',
}

export interface PushEventOptions {
  additionalMetadata?: Record<string, string>;
  priority?: number;
  scope?: string;
}

export interface EventWithMetadata<T> {
  payload: T;
  additionalMetadata?: Record<string, any>;
  priority?: number;
  scope?: string;
}

export class EventClient {
  config: ClientConfig;
  client: EventsServiceClient;
  retrier: typeof retrier;
  api: HatchetClient['api'];
  tenantId: string;

  logger: Logger;

  constructor(
    config: ClientConfig,
    channel: Channel,
    factory: ClientFactory,
    hatchetClient: LegacyHatchetClient
  ) {
    this.config = config;
    this.client = factory.create(EventsServiceDefinition, channel);
    this.logger = config.logger(`Dispatcher`, config.log_level);
    this.retrier = retrier;
    this.api = hatchetClient.api;
    this.tenantId = config.tenant_id;
  }

  push<T>(type: string, input: T, options: PushEventOptions = {}) {
    const namespacedType = applyNamespace(type, this.config.namespace);

    const req: PushEventRequest = {
      key: namespacedType,
      payload: JSON.stringify(input),
      eventTimestamp: new Date(),
      additionalMetadata: options.additionalMetadata
        ? JSON.stringify(options.additionalMetadata)
        : undefined,
      priority: options.priority,
      scope: options.scope,
    };

    try {
      const e = this.retrier(async () => this.client.push(req), this.logger);
      this.logger.info(`Event pushed: ${namespacedType}`);
      return e;
    } catch (e: any) {
      throw new HatchetError(e.message);
    }
  }

  bulkPush<T>(type: string, inputs: EventWithMetadata<T>[], options: PushEventOptions = {}) {
    const namespacedType = applyNamespace(type, this.config.namespace);

    const events = inputs.map((input) => {
      return {
        key: namespacedType,
        payload: JSON.stringify(input.payload),
        eventTimestamp: new Date(),
        additionalMetadata: (() => {
          if (input.additionalMetadata) {
            return JSON.stringify(input.additionalMetadata);
          }
          if (options.additionalMetadata) {
            return JSON.stringify(options.additionalMetadata);
          }
          return undefined;
        })(),
        priority: input.priority,
        scope: input.scope,
      };
    });

    const req: BulkPushEventRequest = {
      events,
    };

    try {
      const res = this.retrier(async () => this.client.bulkPush(req), this.logger);
      this.logger.info(`Bulk events pushed for type: ${namespacedType}`);
      return res;
    } catch (e: any) {
      throw new HatchetError(e.message);
    }
  }

  async putLog(
    stepRunId: string,
    log: string,
    level?: LogLevel,
    taskRetryCount?: number,
    metadata?: Record<string, any>
  ) {
    const createdAt = new Date();

    if (log.length > 1_000) {
      this.logger.warn(`log is too long, skipping: ${log.length} characters`);
      return;
    }

    //  fire and forget the log
    await this.client
      .putLog({
        stepRunId,
        createdAt,
        message: log,
        level: level || LogLevel.INFO,
        taskRetryCount,
        metadata: metadata ? JSON.stringify(metadata) : undefined,
      })
      .catch((e: any) => {
        // log a warning, but this is not a fatal error
        this.logger.warn(`Could not put log: ${e.message}`);
      });
  }

  async putStream(stepRunId: string, data: string | Uint8Array, index: number | undefined) {
    const createdAt = new Date();

    let dataBytes: Uint8Array;
    if (typeof data === 'string') {
      dataBytes = new TextEncoder().encode(data);
    } else if (data instanceof Uint8Array) {
      dataBytes = data;
    } else {
      throw new Error('Invalid data type. Expected string or Uint8Array.');
    }

    retrier(
      async () =>
        this.client.putStreamEvent({
          stepRunId,
          createdAt,
          message: dataBytes,
          eventIndex: index,
        }),
      this.logger
    ).catch((e: any) => {
      // log a warning, but this is not a fatal error
      this.logger.warn(`Could not put log: ${e.message}`);
    });
  }

  async list(opts?: Parameters<typeof this.api.v1EventList>[1]) {
    const { data } = await this.api.v1EventList(this.tenantId, opts);
    return data;
  }
}
