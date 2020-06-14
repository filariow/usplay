// package: bookmastergrpc
// file: bookmastergrpc/activity.proto

import * as bookmastergrpc_activity_pb from "../bookmastergrpc/activity_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ActivitySvcCreate = {
  readonly methodName: string;
  readonly service: typeof ActivitySvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activity_pb.CreateActivityRequest;
  readonly responseType: typeof bookmastergrpc_activity_pb.CreateActivityReply;
};

type ActivitySvcRead = {
  readonly methodName: string;
  readonly service: typeof ActivitySvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activity_pb.ReadActivityRequest;
  readonly responseType: typeof bookmastergrpc_activity_pb.ReadActivityReply;
};

type ActivitySvcDelete = {
  readonly methodName: string;
  readonly service: typeof ActivitySvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activity_pb.DeleteActivityRequest;
  readonly responseType: typeof bookmastergrpc_activity_pb.DeleteActivityReply;
};

type ActivitySvcUpdate = {
  readonly methodName: string;
  readonly service: typeof ActivitySvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activity_pb.UpdateActivityRequest;
  readonly responseType: typeof bookmastergrpc_activity_pb.UpdateActivityReply;
};

type ActivitySvcList = {
  readonly methodName: string;
  readonly service: typeof ActivitySvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activity_pb.ListActivitiesRequest;
  readonly responseType: typeof bookmastergrpc_activity_pb.ListActivitiesReply;
};

type ActivitySvcListInInterval = {
  readonly methodName: string;
  readonly service: typeof ActivitySvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activity_pb.ListInIntervalActivitiesRequest;
  readonly responseType: typeof bookmastergrpc_activity_pb.ListActivitiesReply;
};

export class ActivitySvc {
  static readonly serviceName: string;
  static readonly Create: ActivitySvcCreate;
  static readonly Read: ActivitySvcRead;
  static readonly Delete: ActivitySvcDelete;
  static readonly Update: ActivitySvcUpdate;
  static readonly List: ActivitySvcList;
  static readonly ListInInterval: ActivitySvcListInInterval;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ActivitySvcClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  create(
    requestMessage: bookmastergrpc_activity_pb.CreateActivityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.CreateActivityReply|null) => void
  ): UnaryResponse;
  create(
    requestMessage: bookmastergrpc_activity_pb.CreateActivityRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.CreateActivityReply|null) => void
  ): UnaryResponse;
  read(
    requestMessage: bookmastergrpc_activity_pb.ReadActivityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.ReadActivityReply|null) => void
  ): UnaryResponse;
  read(
    requestMessage: bookmastergrpc_activity_pb.ReadActivityRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.ReadActivityReply|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: bookmastergrpc_activity_pb.DeleteActivityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.DeleteActivityReply|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: bookmastergrpc_activity_pb.DeleteActivityRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.DeleteActivityReply|null) => void
  ): UnaryResponse;
  update(
    requestMessage: bookmastergrpc_activity_pb.UpdateActivityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.UpdateActivityReply|null) => void
  ): UnaryResponse;
  update(
    requestMessage: bookmastergrpc_activity_pb.UpdateActivityRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.UpdateActivityReply|null) => void
  ): UnaryResponse;
  list(
    requestMessage: bookmastergrpc_activity_pb.ListActivitiesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.ListActivitiesReply|null) => void
  ): UnaryResponse;
  list(
    requestMessage: bookmastergrpc_activity_pb.ListActivitiesRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.ListActivitiesReply|null) => void
  ): UnaryResponse;
  listInInterval(
    requestMessage: bookmastergrpc_activity_pb.ListInIntervalActivitiesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.ListActivitiesReply|null) => void
  ): UnaryResponse;
  listInInterval(
    requestMessage: bookmastergrpc_activity_pb.ListInIntervalActivitiesRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activity_pb.ListActivitiesReply|null) => void
  ): UnaryResponse;
}

