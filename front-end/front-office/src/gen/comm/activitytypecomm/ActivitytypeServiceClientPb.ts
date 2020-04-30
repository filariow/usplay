/**
 * @fileoverview gRPC-Web generated client stub for activitytypecomm
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  CreateActivityTypeReply,
  CreateActivityTypeRequest,
  DeleteActivityTypeReply,
  DeleteActivityTypeRequest,
  ExistActivityTypeReply,
  ExistActivityTypeRequest,
  ListActivityTypesReply,
  ListActivityTypesRequest,
  ReadActivityTypeReply,
  ReadActivityTypeRequest,
  UpdateActivityTypeReply,
  UpdateActivityTypeRequest} from './activitytype_pb';

export class ActivityTypeSvcClient {
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
    CreateActivityTypeReply,
    (request: CreateActivityTypeRequest) => {
      return request.serializeBinary();
    },
    CreateActivityTypeReply.deserializeBinary
  );

  create(
    request: CreateActivityTypeRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: CreateActivityTypeReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitytypecomm.ActivityTypeSvc/Create',
      request,
      metadata || {},
      this.methodInfoCreate,
      callback);
  }

  methodInfoExist = new grpcWeb.AbstractClientBase.MethodInfo(
    ExistActivityTypeReply,
    (request: ExistActivityTypeRequest) => {
      return request.serializeBinary();
    },
    ExistActivityTypeReply.deserializeBinary
  );

  exist(
    request: ExistActivityTypeRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ExistActivityTypeReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitytypecomm.ActivityTypeSvc/Exist',
      request,
      metadata || {},
      this.methodInfoExist,
      callback);
  }

  methodInfoRead = new grpcWeb.AbstractClientBase.MethodInfo(
    ReadActivityTypeReply,
    (request: ReadActivityTypeRequest) => {
      return request.serializeBinary();
    },
    ReadActivityTypeReply.deserializeBinary
  );

  read(
    request: ReadActivityTypeRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ReadActivityTypeReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitytypecomm.ActivityTypeSvc/Read',
      request,
      metadata || {},
      this.methodInfoRead,
      callback);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    DeleteActivityTypeReply,
    (request: DeleteActivityTypeRequest) => {
      return request.serializeBinary();
    },
    DeleteActivityTypeReply.deserializeBinary
  );

  delete(
    request: DeleteActivityTypeRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: DeleteActivityTypeReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitytypecomm.ActivityTypeSvc/Delete',
      request,
      metadata || {},
      this.methodInfoDelete,
      callback);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    UpdateActivityTypeReply,
    (request: UpdateActivityTypeRequest) => {
      return request.serializeBinary();
    },
    UpdateActivityTypeReply.deserializeBinary
  );

  update(
    request: UpdateActivityTypeRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: UpdateActivityTypeReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitytypecomm.ActivityTypeSvc/Update',
      request,
      metadata || {},
      this.methodInfoUpdate,
      callback);
  }

  methodInfoList = new grpcWeb.AbstractClientBase.MethodInfo(
    ListActivityTypesReply,
    (request: ListActivityTypesRequest) => {
      return request.serializeBinary();
    },
    ListActivityTypesReply.deserializeBinary
  );

  list(
    request: ListActivityTypesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ListActivityTypesReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitytypecomm.ActivityTypeSvc/List',
      request,
      metadata || {},
      this.methodInfoList,
      callback);
  }

}

