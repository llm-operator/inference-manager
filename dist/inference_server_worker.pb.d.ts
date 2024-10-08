import * as LlmoperatorChatServerV1Inference_server from "./inference_server.pb";
type Absent<T, K extends keyof T> = {
    [k in Exclude<keyof T, K>]?: undefined;
};
type OneOf<T> = {
    [k in keyof T]?: undefined;
} | (keyof T extends infer K ? (K extends string & keyof T ? {
    [k in K]: T[K];
} & Absent<T, K> : never) : never);
export type EngineStatusSyncStatus = {
    inProgressModelIds?: string[];
};
export type EngineStatus = {
    engineId?: string;
    modelIds?: string[];
    syncStatus?: EngineStatusSyncStatus;
};
export type HeaderValue = {
    values?: string[];
};
export type HttpResponse = {
    statusCode?: number;
    status?: string;
    header?: {
        [key: string]: HeaderValue;
    };
    body?: Uint8Array;
};
export type ServerSentEvent = {
    data?: Uint8Array;
    isLastEvent?: boolean;
};
type BaseTaskResult = {
    taskId?: string;
};
export type TaskResult = BaseTaskResult & OneOf<{
    httpResponse: HttpResponse;
    serverSentEvent: ServerSentEvent;
}>;
type BaseProcessTasksRequest = {};
export type ProcessTasksRequest = BaseProcessTasksRequest & OneOf<{
    engineStatus: EngineStatus;
    taskResult: TaskResult;
}>;
export type Task = {
    id?: string;
    request?: LlmoperatorChatServerV1Inference_server.CreateChatCompletionRequest;
    header?: {
        [key: string]: HeaderValue;
    };
};
export type ProcessTasksResponse = {
    newTask?: Task;
};
export declare class InferenceWorkerService {
}
export {};
