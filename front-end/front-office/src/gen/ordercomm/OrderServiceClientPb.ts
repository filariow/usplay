/**
 * @fileoverview gRPC-Web generated client stub for ordercomm
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  CreateOrderReply,
  CreateOrderRequest,
  DeleteOrderReply,
  DeleteOrderRequest,
  ExistOrderReply,
  ExistOrderRequest,
  ListOrdersReply,
  ListOrdersRequest,
  ReadOrderReply,
  ReadOrderRequest,
  UpdateOrderReply,
  UpdateOrderRequest} from './order_pb';

export class OrderSvcClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoCreate = new grpcWeb.AbstractClientBase.MethodInfo(
    CreateOrderReply,
    (request: CreateOrderRequest) => {
      return request.serializeBinary();
    },
    CreateOrderReply.deserializeBinary
  );

  create(
    request: CreateOrderRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: CreateOrderReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/ordercomm.OrderSvc/Create',
      request,
      metadata || {},
      this.methodInfoCreate,
      callback);
  }

  methodInfoRead = new grpcWeb.AbstractClientBase.MethodInfo(
    ReadOrderReply,
    (request: ReadOrderRequest) => {
      return request.serializeBinary();
    },
    ReadOrderReply.deserializeBinary
  );

  read(
    request: ReadOrderRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ReadOrderReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/ordercomm.OrderSvc/Read',
      request,
      metadata || {},
      this.methodInfoRead,
      callback);
  }

  methodInfoExist = new grpcWeb.AbstractClientBase.MethodInfo(
    ExistOrderReply,
    (request: ExistOrderRequest) => {
      return request.serializeBinary();
    },
    ExistOrderReply.deserializeBinary
  );

  exist(
    request: ExistOrderRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ExistOrderReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/ordercomm.OrderSvc/Exist',
      request,
      metadata || {},
      this.methodInfoExist,
      callback);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    DeleteOrderReply,
    (request: DeleteOrderRequest) => {
      return request.serializeBinary();
    },
    DeleteOrderReply.deserializeBinary
  );

  delete(
    request: DeleteOrderRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: DeleteOrderReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/ordercomm.OrderSvc/Delete',
      request,
      metadata || {},
      this.methodInfoDelete,
      callback);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    UpdateOrderReply,
    (request: UpdateOrderRequest) => {
      return request.serializeBinary();
    },
    UpdateOrderReply.deserializeBinary
  );

  update(
    request: UpdateOrderRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: UpdateOrderReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/ordercomm.OrderSvc/Update',
      request,
      metadata || {},
      this.methodInfoUpdate,
      callback);
  }

  methodInfoList = new grpcWeb.AbstractClientBase.MethodInfo(
    ListOrdersReply,
    (request: ListOrdersRequest) => {
      return request.serializeBinary();
    },
    ListOrdersReply.deserializeBinary
  );

  list(
    request: ListOrdersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ListOrdersReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/ordercomm.OrderSvc/List',
      request,
      metadata || {},
      this.methodInfoList,
      callback);
  }

}

