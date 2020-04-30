import * as jspb from "google-protobuf"

export class CreateOrderRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getCode(): string;
  setCode(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateOrderRequest): CreateOrderRequest.AsObject;
  static serializeBinaryToWriter(message: CreateOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateOrderRequest;
  static deserializeBinaryFromReader(message: CreateOrderRequest, reader: jspb.BinaryReader): CreateOrderRequest;
}

export namespace CreateOrderRequest {
  export type AsObject = {
    name: string,
    code: string,
    description: string,
  }
}

export class CreateOrderReply extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateOrderReply.AsObject;
  static toObject(includeInstance: boolean, msg: CreateOrderReply): CreateOrderReply.AsObject;
  static serializeBinaryToWriter(message: CreateOrderReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateOrderReply;
  static deserializeBinaryFromReader(message: CreateOrderReply, reader: jspb.BinaryReader): CreateOrderReply;
}

export namespace CreateOrderReply {
  export type AsObject = {
    id: string,
  }
}

export class ReadOrderRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadOrderRequest): ReadOrderRequest.AsObject;
  static serializeBinaryToWriter(message: ReadOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadOrderRequest;
  static deserializeBinaryFromReader(message: ReadOrderRequest, reader: jspb.BinaryReader): ReadOrderRequest;
}

export namespace ReadOrderRequest {
  export type AsObject = {
    id: string,
  }
}

export class ReadOrderReply extends jspb.Message {
  getOrder(): Order | undefined;
  setOrder(value?: Order): void;
  hasOrder(): boolean;
  clearOrder(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadOrderReply.AsObject;
  static toObject(includeInstance: boolean, msg: ReadOrderReply): ReadOrderReply.AsObject;
  static serializeBinaryToWriter(message: ReadOrderReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadOrderReply;
  static deserializeBinaryFromReader(message: ReadOrderReply, reader: jspb.BinaryReader): ReadOrderReply;
}

export namespace ReadOrderReply {
  export type AsObject = {
    order?: Order.AsObject,
  }
}

export class ExistOrderRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExistOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ExistOrderRequest): ExistOrderRequest.AsObject;
  static serializeBinaryToWriter(message: ExistOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExistOrderRequest;
  static deserializeBinaryFromReader(message: ExistOrderRequest, reader: jspb.BinaryReader): ExistOrderRequest;
}

export namespace ExistOrderRequest {
  export type AsObject = {
    id: string,
  }
}

export class ExistOrderReply extends jspb.Message {
  getExists(): boolean;
  setExists(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExistOrderReply.AsObject;
  static toObject(includeInstance: boolean, msg: ExistOrderReply): ExistOrderReply.AsObject;
  static serializeBinaryToWriter(message: ExistOrderReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExistOrderReply;
  static deserializeBinaryFromReader(message: ExistOrderReply, reader: jspb.BinaryReader): ExistOrderReply;
}

export namespace ExistOrderReply {
  export type AsObject = {
    exists: boolean,
  }
}

export class DeleteOrderRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteOrderRequest): DeleteOrderRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteOrderRequest;
  static deserializeBinaryFromReader(message: DeleteOrderRequest, reader: jspb.BinaryReader): DeleteOrderRequest;
}

export namespace DeleteOrderRequest {
  export type AsObject = {
    id: string,
  }
}

export class DeleteOrderReply extends jspb.Message {
  getOrder(): Order | undefined;
  setOrder(value?: Order): void;
  hasOrder(): boolean;
  clearOrder(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteOrderReply.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteOrderReply): DeleteOrderReply.AsObject;
  static serializeBinaryToWriter(message: DeleteOrderReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteOrderReply;
  static deserializeBinaryFromReader(message: DeleteOrderReply, reader: jspb.BinaryReader): DeleteOrderReply;
}

export namespace DeleteOrderReply {
  export type AsObject = {
    order?: Order.AsObject,
  }
}

export class UpdateOrderRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getCode(): string;
  setCode(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateOrderRequest): UpdateOrderRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateOrderRequest;
  static deserializeBinaryFromReader(message: UpdateOrderRequest, reader: jspb.BinaryReader): UpdateOrderRequest;
}

export namespace UpdateOrderRequest {
  export type AsObject = {
    id: string,
    name: string,
    code: string,
    description: string,
  }
}

export class UpdateOrderReply extends jspb.Message {
  getOrder(): Order | undefined;
  setOrder(value?: Order): void;
  hasOrder(): boolean;
  clearOrder(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateOrderReply.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateOrderReply): UpdateOrderReply.AsObject;
  static serializeBinaryToWriter(message: UpdateOrderReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateOrderReply;
  static deserializeBinaryFromReader(message: UpdateOrderReply, reader: jspb.BinaryReader): UpdateOrderReply;
}

export namespace UpdateOrderReply {
  export type AsObject = {
    order?: Order.AsObject,
  }
}

export class ListOrdersRequest extends jspb.Message {
  getFilteridsList(): Array<string>;
  setFilteridsList(value: Array<string>): void;
  clearFilteridsList(): void;
  addFilterids(value: string, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListOrdersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListOrdersRequest): ListOrdersRequest.AsObject;
  static serializeBinaryToWriter(message: ListOrdersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListOrdersRequest;
  static deserializeBinaryFromReader(message: ListOrdersRequest, reader: jspb.BinaryReader): ListOrdersRequest;
}

export namespace ListOrdersRequest {
  export type AsObject = {
    filteridsList: Array<string>,
  }
}

export class ListOrdersReply extends jspb.Message {
  getOrdersList(): Array<Order>;
  setOrdersList(value: Array<Order>): void;
  clearOrdersList(): void;
  addOrders(value?: Order, index?: number): Order;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListOrdersReply.AsObject;
  static toObject(includeInstance: boolean, msg: ListOrdersReply): ListOrdersReply.AsObject;
  static serializeBinaryToWriter(message: ListOrdersReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListOrdersReply;
  static deserializeBinaryFromReader(message: ListOrdersReply, reader: jspb.BinaryReader): ListOrdersReply;
}

export namespace ListOrdersReply {
  export type AsObject = {
    ordersList: Array<Order.AsObject>,
  }
}

export class Order extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getCode(): string;
  setCode(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Order.AsObject;
  static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
  static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Order;
  static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
}

export namespace Order {
  export type AsObject = {
    id: string,
    name: string,
    code: string,
    description: string,
  }
}

