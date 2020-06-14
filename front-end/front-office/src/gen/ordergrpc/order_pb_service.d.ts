// package: ordergrpc
// file: ordergrpc/order.proto

import * as ordergrpc_order_pb from "../ordergrpc/order_pb";
import {grpc} from "@improbable-eng/grpc-web";

type OrderSvcCreate = {
  readonly methodName: string;
  readonly service: typeof OrderSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ordergrpc_order_pb.CreateOrderRequest;
  readonly responseType: typeof ordergrpc_order_pb.CreateOrderReply;
};

type OrderSvcRead = {
  readonly methodName: string;
  readonly service: typeof OrderSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ordergrpc_order_pb.ReadOrderRequest;
  readonly responseType: typeof ordergrpc_order_pb.ReadOrderReply;
};

type OrderSvcExist = {
  readonly methodName: string;
  readonly service: typeof OrderSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ordergrpc_order_pb.ExistOrderRequest;
  readonly responseType: typeof ordergrpc_order_pb.ExistOrderReply;
};

type OrderSvcDelete = {
  readonly methodName: string;
  readonly service: typeof OrderSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ordergrpc_order_pb.DeleteOrderRequest;
  readonly responseType: typeof ordergrpc_order_pb.DeleteOrderReply;
};

type OrderSvcUpdate = {
  readonly methodName: string;
  readonly service: typeof OrderSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ordergrpc_order_pb.UpdateOrderRequest;
  readonly responseType: typeof ordergrpc_order_pb.UpdateOrderReply;
};

type OrderSvcList = {
  readonly methodName: string;
  readonly service: typeof OrderSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ordergrpc_order_pb.ListOrdersRequest;
  readonly responseType: typeof ordergrpc_order_pb.ListOrdersReply;
};

export class OrderSvc {
  static readonly serviceName: string;
  static readonly Create: OrderSvcCreate;
  static readonly Read: OrderSvcRead;
  static readonly Exist: OrderSvcExist;
  static readonly Delete: OrderSvcDelete;
  static readonly Update: OrderSvcUpdate;
  static readonly List: OrderSvcList;
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

export class OrderSvcClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  create(
    requestMessage: ordergrpc_order_pb.CreateOrderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.CreateOrderReply|null) => void
  ): UnaryResponse;
  create(
    requestMessage: ordergrpc_order_pb.CreateOrderRequest,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.CreateOrderReply|null) => void
  ): UnaryResponse;
  read(
    requestMessage: ordergrpc_order_pb.ReadOrderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.ReadOrderReply|null) => void
  ): UnaryResponse;
  read(
    requestMessage: ordergrpc_order_pb.ReadOrderRequest,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.ReadOrderReply|null) => void
  ): UnaryResponse;
  exist(
    requestMessage: ordergrpc_order_pb.ExistOrderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.ExistOrderReply|null) => void
  ): UnaryResponse;
  exist(
    requestMessage: ordergrpc_order_pb.ExistOrderRequest,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.ExistOrderReply|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: ordergrpc_order_pb.DeleteOrderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.DeleteOrderReply|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: ordergrpc_order_pb.DeleteOrderRequest,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.DeleteOrderReply|null) => void
  ): UnaryResponse;
  update(
    requestMessage: ordergrpc_order_pb.UpdateOrderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.UpdateOrderReply|null) => void
  ): UnaryResponse;
  update(
    requestMessage: ordergrpc_order_pb.UpdateOrderRequest,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.UpdateOrderReply|null) => void
  ): UnaryResponse;
  list(
    requestMessage: ordergrpc_order_pb.ListOrdersRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.ListOrdersReply|null) => void
  ): UnaryResponse;
  list(
    requestMessage: ordergrpc_order_pb.ListOrdersRequest,
    callback: (error: ServiceError|null, responseMessage: ordergrpc_order_pb.ListOrdersReply|null) => void
  ): UnaryResponse;
}

