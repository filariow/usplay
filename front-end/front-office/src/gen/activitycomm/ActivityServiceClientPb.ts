/**
 * @fileoverview gRPC-Web generated client stub for activitycomm
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as activitytypecomm_activitytype_pb from '../activitytypecomm/activitytype_pb';
import * as ordercomm_order_pb from '../ordercomm/order_pb';

import {
  CreateActivityReply,
  CreateActivityRequest,
  DeleteActivityReply,
  DeleteActivityRequest,
  ListActivitiesReply,
  ListActivitiesRequest,
  ListInIntervalActivitiesRequest,
  ReadActivityReply,
  ReadActivityRequest,
  UpdateActivityReply,
  UpdateActivityRequest} from './activity_pb';

export class ActivitySvcClient {
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
    CreateActivityReply,
    (request: CreateActivityRequest) => {
      return request.serializeBinary();
    },
    CreateActivityReply.deserializeBinary
  );

  create(
    request: CreateActivityRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: CreateActivityReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitycomm.ActivitySvc/Create',
      request,
      metadata || {},
      this.methodInfoCreate,
      callback);
  }

  methodInfoRead = new grpcWeb.AbstractClientBase.MethodInfo(
    ReadActivityReply,
    (request: ReadActivityRequest) => {
      return request.serializeBinary();
    },
    ReadActivityReply.deserializeBinary
  );

  read(
    request: ReadActivityRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ReadActivityReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitycomm.ActivitySvc/Read',
      request,
      metadata || {},
      this.methodInfoRead,
      callback);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    DeleteActivityReply,
    (request: DeleteActivityRequest) => {
      return request.serializeBinary();
    },
    DeleteActivityReply.deserializeBinary
  );

  delete(
    request: DeleteActivityRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: DeleteActivityReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitycomm.ActivitySvc/Delete',
      request,
      metadata || {},
      this.methodInfoDelete,
      callback);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    UpdateActivityReply,
    (request: UpdateActivityRequest) => {
      return request.serializeBinary();
    },
    UpdateActivityReply.deserializeBinary
  );

  update(
    request: UpdateActivityRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: UpdateActivityReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitycomm.ActivitySvc/Update',
      request,
      metadata || {},
      this.methodInfoUpdate,
      callback);
  }

  methodInfoList = new grpcWeb.AbstractClientBase.MethodInfo(
    ListActivitiesReply,
    (request: ListActivitiesRequest) => {
      return request.serializeBinary();
    },
    ListActivitiesReply.deserializeBinary
  );

  list(
    request: ListActivitiesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ListActivitiesReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitycomm.ActivitySvc/List',
      request,
      metadata || {},
      this.methodInfoList,
      callback);
  }

  methodInfoListInInterval = new grpcWeb.AbstractClientBase.MethodInfo(
    ListActivitiesReply,
    (request: ListInIntervalActivitiesRequest) => {
      return request.serializeBinary();
    },
    ListActivitiesReply.deserializeBinary
  );

  listInInterval(
    request: ListInIntervalActivitiesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ListActivitiesReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/activitycomm.ActivitySvc/ListInInterval',
      request,
      metadata || {},
      this.methodInfoListInInterval,
      callback);
  }

}

