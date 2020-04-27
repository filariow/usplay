import * as jspb from "google-protobuf"

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as activitytypecomm_activitytype_pb from '../activitytypecomm/activitytype_pb';
import * as ordercomm_order_pb from '../ordercomm/order_pb';

export class CreateActivityRequest extends jspb.Message {
  getActtypeid(): string;
  setActtypeid(value: string): void;

  getOrderid(): string;
  setOrderid(value: string): void;

  getPeriod(): Interval | undefined;
  setPeriod(value?: Interval): void;
  hasPeriod(): boolean;
  clearPeriod(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateActivityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateActivityRequest): CreateActivityRequest.AsObject;
  static serializeBinaryToWriter(message: CreateActivityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateActivityRequest;
  static deserializeBinaryFromReader(message: CreateActivityRequest, reader: jspb.BinaryReader): CreateActivityRequest;
}

export namespace CreateActivityRequest {
  export type AsObject = {
    acttypeid: string,
    orderid: string,
    period?: Interval.AsObject,
  }
}

export class CreateActivityReply extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateActivityReply.AsObject;
  static toObject(includeInstance: boolean, msg: CreateActivityReply): CreateActivityReply.AsObject;
  static serializeBinaryToWriter(message: CreateActivityReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateActivityReply;
  static deserializeBinaryFromReader(message: CreateActivityReply, reader: jspb.BinaryReader): CreateActivityReply;
}

export namespace CreateActivityReply {
  export type AsObject = {
    id: string,
  }
}

export class ReadActivityRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadActivityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadActivityRequest): ReadActivityRequest.AsObject;
  static serializeBinaryToWriter(message: ReadActivityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadActivityRequest;
  static deserializeBinaryFromReader(message: ReadActivityRequest, reader: jspb.BinaryReader): ReadActivityRequest;
}

export namespace ReadActivityRequest {
  export type AsObject = {
    id: string,
  }
}

export class ReadActivityReply extends jspb.Message {
  getActivity(): Activity | undefined;
  setActivity(value?: Activity): void;
  hasActivity(): boolean;
  clearActivity(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadActivityReply.AsObject;
  static toObject(includeInstance: boolean, msg: ReadActivityReply): ReadActivityReply.AsObject;
  static serializeBinaryToWriter(message: ReadActivityReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadActivityReply;
  static deserializeBinaryFromReader(message: ReadActivityReply, reader: jspb.BinaryReader): ReadActivityReply;
}

export namespace ReadActivityReply {
  export type AsObject = {
    activity?: Activity.AsObject,
  }
}

export class DeleteActivityRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteActivityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteActivityRequest): DeleteActivityRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteActivityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteActivityRequest;
  static deserializeBinaryFromReader(message: DeleteActivityRequest, reader: jspb.BinaryReader): DeleteActivityRequest;
}

export namespace DeleteActivityRequest {
  export type AsObject = {
    id: string,
  }
}

export class DeleteActivityReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteActivityReply.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteActivityReply): DeleteActivityReply.AsObject;
  static serializeBinaryToWriter(message: DeleteActivityReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteActivityReply;
  static deserializeBinaryFromReader(message: DeleteActivityReply, reader: jspb.BinaryReader): DeleteActivityReply;
}

export namespace DeleteActivityReply {
  export type AsObject = {
  }
}

export class UpdateActivityRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getActtypeid(): string;
  setActtypeid(value: string): void;

  getOrderid(): string;
  setOrderid(value: string): void;

  getPeriod(): Interval | undefined;
  setPeriod(value?: Interval): void;
  hasPeriod(): boolean;
  clearPeriod(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateActivityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateActivityRequest): UpdateActivityRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateActivityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateActivityRequest;
  static deserializeBinaryFromReader(message: UpdateActivityRequest, reader: jspb.BinaryReader): UpdateActivityRequest;
}

export namespace UpdateActivityRequest {
  export type AsObject = {
    id: string,
    acttypeid: string,
    orderid: string,
    period?: Interval.AsObject,
  }
}

export class UpdateActivityReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateActivityReply.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateActivityReply): UpdateActivityReply.AsObject;
  static serializeBinaryToWriter(message: UpdateActivityReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateActivityReply;
  static deserializeBinaryFromReader(message: UpdateActivityReply, reader: jspb.BinaryReader): UpdateActivityReply;
}

export namespace UpdateActivityReply {
  export type AsObject = {
  }
}

export class ListActivitiesRequest extends jspb.Message {
  getFilteridsList(): Array<string>;
  setFilteridsList(value: Array<string>): void;
  clearFilteridsList(): void;
  addFilterids(value: string, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListActivitiesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListActivitiesRequest): ListActivitiesRequest.AsObject;
  static serializeBinaryToWriter(message: ListActivitiesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListActivitiesRequest;
  static deserializeBinaryFromReader(message: ListActivitiesRequest, reader: jspb.BinaryReader): ListActivitiesRequest;
}

export namespace ListActivitiesRequest {
  export type AsObject = {
    filteridsList: Array<string>,
  }
}

export class ListInIntervalActivitiesRequest extends jspb.Message {
  getPeriod(): Interval | undefined;
  setPeriod(value?: Interval): void;
  hasPeriod(): boolean;
  clearPeriod(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListInIntervalActivitiesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListInIntervalActivitiesRequest): ListInIntervalActivitiesRequest.AsObject;
  static serializeBinaryToWriter(message: ListInIntervalActivitiesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListInIntervalActivitiesRequest;
  static deserializeBinaryFromReader(message: ListInIntervalActivitiesRequest, reader: jspb.BinaryReader): ListInIntervalActivitiesRequest;
}

export namespace ListInIntervalActivitiesRequest {
  export type AsObject = {
    period?: Interval.AsObject,
  }
}

export class ListActivitiesReply extends jspb.Message {
  getActivitiesList(): Array<Activity>;
  setActivitiesList(value: Array<Activity>): void;
  clearActivitiesList(): void;
  addActivities(value?: Activity, index?: number): Activity;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListActivitiesReply.AsObject;
  static toObject(includeInstance: boolean, msg: ListActivitiesReply): ListActivitiesReply.AsObject;
  static serializeBinaryToWriter(message: ListActivitiesReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListActivitiesReply;
  static deserializeBinaryFromReader(message: ListActivitiesReply, reader: jspb.BinaryReader): ListActivitiesReply;
}

export namespace ListActivitiesReply {
  export type AsObject = {
    activitiesList: Array<Activity.AsObject>,
  }
}

export class Activity extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getActtype(): activitytypecomm_activitytype_pb.ActivityType | undefined;
  setActtype(value?: activitytypecomm_activitytype_pb.ActivityType): void;
  hasActtype(): boolean;
  clearActtype(): void;

  getCreationtime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreationtime(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasCreationtime(): boolean;
  clearCreationtime(): void;

  getOrder(): ordercomm_order_pb.Order | undefined;
  setOrder(value?: ordercomm_order_pb.Order): void;
  hasOrder(): boolean;
  clearOrder(): void;

  getPeriod(): Interval | undefined;
  setPeriod(value?: Interval): void;
  hasPeriod(): boolean;
  clearPeriod(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Activity.AsObject;
  static toObject(includeInstance: boolean, msg: Activity): Activity.AsObject;
  static serializeBinaryToWriter(message: Activity, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Activity;
  static deserializeBinaryFromReader(message: Activity, reader: jspb.BinaryReader): Activity;
}

export namespace Activity {
  export type AsObject = {
    id: string,
    acttype?: activitytypecomm_activitytype_pb.ActivityType.AsObject,
    creationtime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    order?: ordercomm_order_pb.Order.AsObject,
    period?: Interval.AsObject,
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

