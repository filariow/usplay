import * as jspb from "google-protobuf"

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as bookmastergrpc_activity_pb from '../bookmastergrpc/activity_pb';

export class CreateReportRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getPeriod(): Interval | undefined;
  setPeriod(value?: Interval): void;
  hasPeriod(): boolean;
  clearPeriod(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateReportRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateReportRequest): CreateReportRequest.AsObject;
  static serializeBinaryToWriter(message: CreateReportRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateReportRequest;
  static deserializeBinaryFromReader(message: CreateReportRequest, reader: jspb.BinaryReader): CreateReportRequest;
}

export namespace CreateReportRequest {
  export type AsObject = {
    name: string,
    period?: Interval.AsObject,
  }
}

export class CreateReportReply extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateReportReply.AsObject;
  static toObject(includeInstance: boolean, msg: CreateReportReply): CreateReportReply.AsObject;
  static serializeBinaryToWriter(message: CreateReportReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateReportReply;
  static deserializeBinaryFromReader(message: CreateReportReply, reader: jspb.BinaryReader): CreateReportReply;
}

export namespace CreateReportReply {
  export type AsObject = {
    id: string,
  }
}

export class ReadReportRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadReportRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadReportRequest): ReadReportRequest.AsObject;
  static serializeBinaryToWriter(message: ReadReportRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadReportRequest;
  static deserializeBinaryFromReader(message: ReadReportRequest, reader: jspb.BinaryReader): ReadReportRequest;
}

export namespace ReadReportRequest {
  export type AsObject = {
    id: string,
  }
}

export class ReadReportReply extends jspb.Message {
  getReport(): Report | undefined;
  setReport(value?: Report): void;
  hasReport(): boolean;
  clearReport(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadReportReply.AsObject;
  static toObject(includeInstance: boolean, msg: ReadReportReply): ReadReportReply.AsObject;
  static serializeBinaryToWriter(message: ReadReportReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadReportReply;
  static deserializeBinaryFromReader(message: ReadReportReply, reader: jspb.BinaryReader): ReadReportReply;
}

export namespace ReadReportReply {
  export type AsObject = {
    report?: Report.AsObject,
  }
}

export class DeleteReportRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteReportRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteReportRequest): DeleteReportRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteReportRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteReportRequest;
  static deserializeBinaryFromReader(message: DeleteReportRequest, reader: jspb.BinaryReader): DeleteReportRequest;
}

export namespace DeleteReportRequest {
  export type AsObject = {
    id: string,
  }
}

export class DeleteReportReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteReportReply.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteReportReply): DeleteReportReply.AsObject;
  static serializeBinaryToWriter(message: DeleteReportReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteReportReply;
  static deserializeBinaryFromReader(message: DeleteReportReply, reader: jspb.BinaryReader): DeleteReportReply;
}

export namespace DeleteReportReply {
  export type AsObject = {
  }
}

export class UpdateReportRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getCode(): string;
  setCode(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateReportRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateReportRequest): UpdateReportRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateReportRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateReportRequest;
  static deserializeBinaryFromReader(message: UpdateReportRequest, reader: jspb.BinaryReader): UpdateReportRequest;
}

export namespace UpdateReportRequest {
  export type AsObject = {
    id: string,
    name: string,
    code: string,
    description: string,
  }
}

export class UpdateReportReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateReportReply.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateReportReply): UpdateReportReply.AsObject;
  static serializeBinaryToWriter(message: UpdateReportReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateReportReply;
  static deserializeBinaryFromReader(message: UpdateReportReply, reader: jspb.BinaryReader): UpdateReportReply;
}

export namespace UpdateReportReply {
  export type AsObject = {
  }
}

export class ListReportsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListReportsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListReportsRequest): ListReportsRequest.AsObject;
  static serializeBinaryToWriter(message: ListReportsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListReportsRequest;
  static deserializeBinaryFromReader(message: ListReportsRequest, reader: jspb.BinaryReader): ListReportsRequest;
}

export namespace ListReportsRequest {
  export type AsObject = {
  }
}

export class ListReportsReply extends jspb.Message {
  getReportsList(): Array<Report>;
  setReportsList(value: Array<Report>): void;
  clearReportsList(): void;
  addReports(value?: Report, index?: number): Report;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListReportsReply.AsObject;
  static toObject(includeInstance: boolean, msg: ListReportsReply): ListReportsReply.AsObject;
  static serializeBinaryToWriter(message: ListReportsReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListReportsReply;
  static deserializeBinaryFromReader(message: ListReportsReply, reader: jspb.BinaryReader): ListReportsReply;
}

export namespace ListReportsReply {
  export type AsObject = {
    reportsList: Array<Report.AsObject>,
  }
}

export class Report extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getOrderid(): string;
  setOrderid(value: string): void;

  getPeriod(): Interval | undefined;
  setPeriod(value?: Interval): void;
  hasPeriod(): boolean;
  clearPeriod(): void;

  getActivitiesList(): Array<bookmastergrpc_activity_pb.Activity>;
  setActivitiesList(value: Array<bookmastergrpc_activity_pb.Activity>): void;
  clearActivitiesList(): void;
  addActivities(value?: bookmastergrpc_activity_pb.Activity, index?: number): bookmastergrpc_activity_pb.Activity;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Report.AsObject;
  static toObject(includeInstance: boolean, msg: Report): Report.AsObject;
  static serializeBinaryToWriter(message: Report, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Report;
  static deserializeBinaryFromReader(message: Report, reader: jspb.BinaryReader): Report;
}

export namespace Report {
  export type AsObject = {
    id: string,
    name: string,
    orderid: string,
    period?: Interval.AsObject,
    activitiesList: Array<bookmastergrpc_activity_pb.Activity.AsObject>,
  }
}

export class Interval extends jspb.Message {
  getFrom(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFrom(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasFrom(): boolean;
  clearFrom(): void;

  getTo(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTo(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasTo(): boolean;
  clearTo(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Interval.AsObject;
  static toObject(includeInstance: boolean, msg: Interval): Interval.AsObject;
  static serializeBinaryToWriter(message: Interval, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Interval;
  static deserializeBinaryFromReader(message: Interval, reader: jspb.BinaryReader): Interval;
}

export namespace Interval {
  export type AsObject = {
    from?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    to?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

