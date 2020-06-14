// package: bookmastergrpc
// file: bookmastergrpc/activitytype.proto

import * as jspb from "google-protobuf";

export class CreateActivityTypeRequest extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getName(): string;
  setName(value: string): void;

  getNeedsorder(): boolean;
  setNeedsorder(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateActivityTypeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateActivityTypeRequest): CreateActivityTypeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateActivityTypeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateActivityTypeRequest;
  static deserializeBinaryFromReader(message: CreateActivityTypeRequest, reader: jspb.BinaryReader): CreateActivityTypeRequest;
}

export namespace CreateActivityTypeRequest {
  export type AsObject = {
    code: number,
    name: string,
    needsorder: boolean,
  }
}

export class CreateActivityTypeReply extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateActivityTypeReply.AsObject;
  static toObject(includeInstance: boolean, msg: CreateActivityTypeReply): CreateActivityTypeReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateActivityTypeReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateActivityTypeReply;
  static deserializeBinaryFromReader(message: CreateActivityTypeReply, reader: jspb.BinaryReader): CreateActivityTypeReply;
}

export namespace CreateActivityTypeReply {
  export type AsObject = {
    id: string,
  }
}

export class ExistActivityTypeRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExistActivityTypeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ExistActivityTypeRequest): ExistActivityTypeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExistActivityTypeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExistActivityTypeRequest;
  static deserializeBinaryFromReader(message: ExistActivityTypeRequest, reader: jspb.BinaryReader): ExistActivityTypeRequest;
}

export namespace ExistActivityTypeRequest {
  export type AsObject = {
    id: string,
  }
}

export class ExistActivityTypeReply extends jspb.Message {
  getExists(): boolean;
  setExists(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExistActivityTypeReply.AsObject;
  static toObject(includeInstance: boolean, msg: ExistActivityTypeReply): ExistActivityTypeReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExistActivityTypeReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExistActivityTypeReply;
  static deserializeBinaryFromReader(message: ExistActivityTypeReply, reader: jspb.BinaryReader): ExistActivityTypeReply;
}

export namespace ExistActivityTypeReply {
  export type AsObject = {
    exists: boolean,
  }
}

export class ReadActivityTypeRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadActivityTypeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadActivityTypeRequest): ReadActivityTypeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadActivityTypeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadActivityTypeRequest;
  static deserializeBinaryFromReader(message: ReadActivityTypeRequest, reader: jspb.BinaryReader): ReadActivityTypeRequest;
}

export namespace ReadActivityTypeRequest {
  export type AsObject = {
    id: string,
  }
}

export class ReadActivityTypeReply extends jspb.Message {
  hasActivitytype(): boolean;
  clearActivitytype(): void;
  getActivitytype(): ActivityType | undefined;
  setActivitytype(value?: ActivityType): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadActivityTypeReply.AsObject;
  static toObject(includeInstance: boolean, msg: ReadActivityTypeReply): ReadActivityTypeReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadActivityTypeReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadActivityTypeReply;
  static deserializeBinaryFromReader(message: ReadActivityTypeReply, reader: jspb.BinaryReader): ReadActivityTypeReply;
}

export namespace ReadActivityTypeReply {
  export type AsObject = {
    activitytype?: ActivityType.AsObject,
  }
}

export class DeleteActivityTypeRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteActivityTypeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteActivityTypeRequest): DeleteActivityTypeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteActivityTypeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteActivityTypeRequest;
  static deserializeBinaryFromReader(message: DeleteActivityTypeRequest, reader: jspb.BinaryReader): DeleteActivityTypeRequest;
}

export namespace DeleteActivityTypeRequest {
  export type AsObject = {
    id: string,
  }
}

export class DeleteActivityTypeReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteActivityTypeReply.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteActivityTypeReply): DeleteActivityTypeReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteActivityTypeReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteActivityTypeReply;
  static deserializeBinaryFromReader(message: DeleteActivityTypeReply, reader: jspb.BinaryReader): DeleteActivityTypeReply;
}

export namespace DeleteActivityTypeReply {
  export type AsObject = {
  }
}

export class UpdateActivityTypeRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getCode(): number;
  setCode(value: number): void;

  getName(): string;
  setName(value: string): void;

  getNeedsorder(): boolean;
  setNeedsorder(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateActivityTypeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateActivityTypeRequest): UpdateActivityTypeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateActivityTypeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateActivityTypeRequest;
  static deserializeBinaryFromReader(message: UpdateActivityTypeRequest, reader: jspb.BinaryReader): UpdateActivityTypeRequest;
}

export namespace UpdateActivityTypeRequest {
  export type AsObject = {
    id: string,
    code: number,
    name: string,
    needsorder: boolean,
  }
}

export class UpdateActivityTypeReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateActivityTypeReply.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateActivityTypeReply): UpdateActivityTypeReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateActivityTypeReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateActivityTypeReply;
  static deserializeBinaryFromReader(message: UpdateActivityTypeReply, reader: jspb.BinaryReader): UpdateActivityTypeReply;
}

export namespace UpdateActivityTypeReply {
  export type AsObject = {
  }
}

export class ListActivityTypesRequest extends jspb.Message {
  clearFilteridsList(): void;
  getFilteridsList(): Array<string>;
  setFilteridsList(value: Array<string>): void;
  addFilterids(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListActivityTypesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListActivityTypesRequest): ListActivityTypesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListActivityTypesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListActivityTypesRequest;
  static deserializeBinaryFromReader(message: ListActivityTypesRequest, reader: jspb.BinaryReader): ListActivityTypesRequest;
}

export namespace ListActivityTypesRequest {
  export type AsObject = {
    filteridsList: Array<string>,
  }
}

export class ListActivityTypesReply extends jspb.Message {
  clearActivitytypesList(): void;
  getActivitytypesList(): Array<ActivityType>;
  setActivitytypesList(value: Array<ActivityType>): void;
  addActivitytypes(value?: ActivityType, index?: number): ActivityType;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListActivityTypesReply.AsObject;
  static toObject(includeInstance: boolean, msg: ListActivityTypesReply): ListActivityTypesReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListActivityTypesReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListActivityTypesReply;
  static deserializeBinaryFromReader(message: ListActivityTypesReply, reader: jspb.BinaryReader): ListActivityTypesReply;
}

export namespace ListActivityTypesReply {
  export type AsObject = {
    activitytypesList: Array<ActivityType.AsObject>,
  }
}

export class ActivityType extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getCode(): number;
  setCode(value: number): void;

  getName(): string;
  setName(value: string): void;

  getNeedsorder(): boolean;
  setNeedsorder(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActivityType.AsObject;
  static toObject(includeInstance: boolean, msg: ActivityType): ActivityType.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ActivityType, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActivityType;
  static deserializeBinaryFromReader(message: ActivityType, reader: jspb.BinaryReader): ActivityType;
}

export namespace ActivityType {
  export type AsObject = {
    id: string,
    code: number,
    name: string,
    needsorder: boolean,
  }
}

