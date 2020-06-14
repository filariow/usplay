// package: bookmastergrpc
// file: bookmastergrpc/activitytype.proto

import * as bookmastergrpc_activitytype_pb from "../bookmastergrpc/activitytype_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ActivityTypeSvcCreate = {
  readonly methodName: string;
  readonly service: typeof ActivityTypeSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activitytype_pb.CreateActivityTypeRequest;
  readonly responseType: typeof bookmastergrpc_activitytype_pb.CreateActivityTypeReply;
};

type ActivityTypeSvcExist = {
  readonly methodName: string;
  readonly service: typeof ActivityTypeSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activitytype_pb.ExistActivityTypeRequest;
  readonly responseType: typeof bookmastergrpc_activitytype_pb.ExistActivityTypeReply;
};

type ActivityTypeSvcRead = {
  readonly methodName: string;
  readonly service: typeof ActivityTypeSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activitytype_pb.ReadActivityTypeRequest;
  readonly responseType: typeof bookmastergrpc_activitytype_pb.ReadActivityTypeReply;
};

type ActivityTypeSvcDelete = {
  readonly methodName: string;
  readonly service: typeof ActivityTypeSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activitytype_pb.DeleteActivityTypeRequest;
  readonly responseType: typeof bookmastergrpc_activitytype_pb.DeleteActivityTypeReply;
};

type ActivityTypeSvcUpdate = {
  readonly methodName: string;
  readonly service: typeof ActivityTypeSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activitytype_pb.UpdateActivityTypeRequest;
  readonly responseType: typeof bookmastergrpc_activitytype_pb.UpdateActivityTypeReply;
};

type ActivityTypeSvcList = {
  readonly methodName: string;
  readonly service: typeof ActivityTypeSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof bookmastergrpc_activitytype_pb.ListActivityTypesRequest;
  readonly responseType: typeof bookmastergrpc_activitytype_pb.ListActivityTypesReply;
};

export class ActivityTypeSvc {
  static readonly serviceName: string;
  static readonly Create: ActivityTypeSvcCreate;
  static readonly Exist: ActivityTypeSvcExist;
  static readonly Read: ActivityTypeSvcRead;
  static readonly Delete: ActivityTypeSvcDelete;
  static readonly Update: ActivityTypeSvcUpdate;
  static readonly List: ActivityTypeSvcList;
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

export class ActivityTypeSvcClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  create(
    requestMessage: bookmastergrpc_activitytype_pb.CreateActivityTypeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.CreateActivityTypeReply|null) => void
  ): UnaryResponse;
  create(
    requestMessage: bookmastergrpc_activitytype_pb.CreateActivityTypeRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.CreateActivityTypeReply|null) => void
  ): UnaryResponse;
  exist(
    requestMessage: bookmastergrpc_activitytype_pb.ExistActivityTypeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.ExistActivityTypeReply|null) => void
  ): UnaryResponse;
  exist(
    requestMessage: bookmastergrpc_activitytype_pb.ExistActivityTypeRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.ExistActivityTypeReply|null) => void
  ): UnaryResponse;
  read(
    requestMessage: bookmastergrpc_activitytype_pb.ReadActivityTypeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.ReadActivityTypeReply|null) => void
  ): UnaryResponse;
  read(
    requestMessage: bookmastergrpc_activitytype_pb.ReadActivityTypeRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.ReadActivityTypeReply|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: bookmastergrpc_activitytype_pb.DeleteActivityTypeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.DeleteActivityTypeReply|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: bookmastergrpc_activitytype_pb.DeleteActivityTypeRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.DeleteActivityTypeReply|null) => void
  ): UnaryResponse;
  update(
    requestMessage: bookmastergrpc_activitytype_pb.UpdateActivityTypeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.UpdateActivityTypeReply|null) => void
  ): UnaryResponse;
  update(
    requestMessage: bookmastergrpc_activitytype_pb.UpdateActivityTypeRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.UpdateActivityTypeReply|null) => void
  ): UnaryResponse;
  list(
    requestMessage: bookmastergrpc_activitytype_pb.ListActivityTypesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.ListActivityTypesReply|null) => void
  ): UnaryResponse;
  list(
    requestMessage: bookmastergrpc_activitytype_pb.ListActivityTypesRequest,
    callback: (error: ServiceError|null, responseMessage: bookmastergrpc_activitytype_pb.ListActivityTypesReply|null) => void
  ): UnaryResponse;
}

