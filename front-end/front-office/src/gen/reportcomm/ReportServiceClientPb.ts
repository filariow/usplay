/**
 * @fileoverview gRPC-Web generated client stub for reportgrpc
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as bookmastergrpc_activity_pb from '../bookmastergrpc/activity_pb';

import {
  CreateReportReply,
  CreateReportRequest,
  DeleteReportReply,
  DeleteReportRequest,
  ListReportsReply,
  ListReportsRequest,
  ReadReportReply,
  ReadReportRequest} from './report_pb';

export class ReportSvcClient {
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
    CreateReportReply,
    (request: CreateReportRequest) => {
      return request.serializeBinary();
    },
    CreateReportReply.deserializeBinary
  );

  create(
    request: CreateReportRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: CreateReportReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/reportgrpc.ReportSvc/Create',
      request,
      metadata || {},
      this.methodInfoCreate,
      callback);
  }

  methodInfoRead = new grpcWeb.AbstractClientBase.MethodInfo(
    ReadReportReply,
    (request: ReadReportRequest) => {
      return request.serializeBinary();
    },
    ReadReportReply.deserializeBinary
  );

  read(
    request: ReadReportRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ReadReportReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/reportgrpc.ReportSvc/Read',
      request,
      metadata || {},
      this.methodInfoRead,
      callback);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    DeleteReportReply,
    (request: DeleteReportRequest) => {
      return request.serializeBinary();
    },
    DeleteReportReply.deserializeBinary
  );

  delete(
    request: DeleteReportRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: DeleteReportReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/reportgrpc.ReportSvc/Delete',
      request,
      metadata || {},
      this.methodInfoDelete,
      callback);
  }

  methodInfoList = new grpcWeb.AbstractClientBase.MethodInfo(
    ListReportsReply,
    (request: ListReportsRequest) => {
      return request.serializeBinary();
    },
    ListReportsReply.deserializeBinary
  );

  list(
    request: ListReportsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ListReportsReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/reportgrpc.ReportSvc/List',
      request,
      metadata || {},
      this.methodInfoList,
      callback);
  }

}

