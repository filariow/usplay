// package: reportgrpc
// file: reportgrpc/report.proto

import * as reportgrpc_report_pb from "../reportgrpc/report_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ReportSvcCreate = {
  readonly methodName: string;
  readonly service: typeof ReportSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof reportgrpc_report_pb.CreateReportRequest;
  readonly responseType: typeof reportgrpc_report_pb.CreateReportReply;
};

type ReportSvcRead = {
  readonly methodName: string;
  readonly service: typeof ReportSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof reportgrpc_report_pb.ReadReportRequest;
  readonly responseType: typeof reportgrpc_report_pb.ReadReportReply;
};

type ReportSvcDelete = {
  readonly methodName: string;
  readonly service: typeof ReportSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof reportgrpc_report_pb.DeleteReportRequest;
  readonly responseType: typeof reportgrpc_report_pb.DeleteReportReply;
};

type ReportSvcList = {
  readonly methodName: string;
  readonly service: typeof ReportSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof reportgrpc_report_pb.ListReportsRequest;
  readonly responseType: typeof reportgrpc_report_pb.ListReportsReply;
};

export class ReportSvc {
  static readonly serviceName: string;
  static readonly Create: ReportSvcCreate;
  static readonly Read: ReportSvcRead;
  static readonly Delete: ReportSvcDelete;
  static readonly List: ReportSvcList;
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

export class ReportSvcClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  create(
    requestMessage: reportgrpc_report_pb.CreateReportRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: reportgrpc_report_pb.CreateReportReply|null) => void
  ): UnaryResponse;
  create(
    requestMessage: reportgrpc_report_pb.CreateReportRequest,
    callback: (error: ServiceError|null, responseMessage: reportgrpc_report_pb.CreateReportReply|null) => void
  ): UnaryResponse;
  read(
    requestMessage: reportgrpc_report_pb.ReadReportRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: reportgrpc_report_pb.ReadReportReply|null) => void
  ): UnaryResponse;
  read(
    requestMessage: reportgrpc_report_pb.ReadReportRequest,
    callback: (error: ServiceError|null, responseMessage: reportgrpc_report_pb.ReadReportReply|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: reportgrpc_report_pb.DeleteReportRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: reportgrpc_report_pb.DeleteReportReply|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: reportgrpc_report_pb.DeleteReportRequest,
    callback: (error: ServiceError|null, responseMessage: reportgrpc_report_pb.DeleteReportReply|null) => void
  ): UnaryResponse;
  list(
    requestMessage: reportgrpc_report_pb.ListReportsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: reportgrpc_report_pb.ListReportsReply|null) => void
  ): UnaryResponse;
  list(
    requestMessage: reportgrpc_report_pb.ListReportsRequest,
    callback: (error: ServiceError|null, responseMessage: reportgrpc_report_pb.ListReportsReply|null) => void
  ): UnaryResponse;
}

