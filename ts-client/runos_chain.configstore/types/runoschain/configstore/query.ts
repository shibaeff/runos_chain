/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { HostsDatabase } from "./hosts_database";
import { Params } from "./params";

export const protobufPackage = "runos_chain.configstore";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetHostsDatabaseRequest {
  dpid: string;
  mac: string;
}

export interface QueryGetHostsDatabaseResponse {
  hostsDatabase: HostsDatabase | undefined;
}

export interface QueryAllHostsDatabaseRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllHostsDatabaseResponse {
  hostsDatabase: HostsDatabase[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetHostsDatabaseRequest(): QueryGetHostsDatabaseRequest {
  return { dpid: "", mac: "" };
}

export const QueryGetHostsDatabaseRequest = {
  encode(message: QueryGetHostsDatabaseRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.dpid !== "") {
      writer.uint32(10).string(message.dpid);
    }
    if (message.mac !== "") {
      writer.uint32(18).string(message.mac);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetHostsDatabaseRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetHostsDatabaseRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.dpid = reader.string();
          break;
        case 2:
          message.mac = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetHostsDatabaseRequest {
    return { dpid: isSet(object.dpid) ? String(object.dpid) : "", mac: isSet(object.mac) ? String(object.mac) : "" };
  },

  toJSON(message: QueryGetHostsDatabaseRequest): unknown {
    const obj: any = {};
    message.dpid !== undefined && (obj.dpid = message.dpid);
    message.mac !== undefined && (obj.mac = message.mac);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetHostsDatabaseRequest>, I>>(object: I): QueryGetHostsDatabaseRequest {
    const message = createBaseQueryGetHostsDatabaseRequest();
    message.dpid = object.dpid ?? "";
    message.mac = object.mac ?? "";
    return message;
  },
};

function createBaseQueryGetHostsDatabaseResponse(): QueryGetHostsDatabaseResponse {
  return { hostsDatabase: undefined };
}

export const QueryGetHostsDatabaseResponse = {
  encode(message: QueryGetHostsDatabaseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.hostsDatabase !== undefined) {
      HostsDatabase.encode(message.hostsDatabase, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetHostsDatabaseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetHostsDatabaseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hostsDatabase = HostsDatabase.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetHostsDatabaseResponse {
    return { hostsDatabase: isSet(object.hostsDatabase) ? HostsDatabase.fromJSON(object.hostsDatabase) : undefined };
  },

  toJSON(message: QueryGetHostsDatabaseResponse): unknown {
    const obj: any = {};
    message.hostsDatabase !== undefined
      && (obj.hostsDatabase = message.hostsDatabase ? HostsDatabase.toJSON(message.hostsDatabase) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetHostsDatabaseResponse>, I>>(
    object: I,
  ): QueryGetHostsDatabaseResponse {
    const message = createBaseQueryGetHostsDatabaseResponse();
    message.hostsDatabase = (object.hostsDatabase !== undefined && object.hostsDatabase !== null)
      ? HostsDatabase.fromPartial(object.hostsDatabase)
      : undefined;
    return message;
  },
};

function createBaseQueryAllHostsDatabaseRequest(): QueryAllHostsDatabaseRequest {
  return { pagination: undefined };
}

export const QueryAllHostsDatabaseRequest = {
  encode(message: QueryAllHostsDatabaseRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllHostsDatabaseRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllHostsDatabaseRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllHostsDatabaseRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllHostsDatabaseRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllHostsDatabaseRequest>, I>>(object: I): QueryAllHostsDatabaseRequest {
    const message = createBaseQueryAllHostsDatabaseRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllHostsDatabaseResponse(): QueryAllHostsDatabaseResponse {
  return { hostsDatabase: [], pagination: undefined };
}

export const QueryAllHostsDatabaseResponse = {
  encode(message: QueryAllHostsDatabaseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.hostsDatabase) {
      HostsDatabase.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllHostsDatabaseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllHostsDatabaseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hostsDatabase.push(HostsDatabase.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllHostsDatabaseResponse {
    return {
      hostsDatabase: Array.isArray(object?.hostsDatabase)
        ? object.hostsDatabase.map((e: any) => HostsDatabase.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllHostsDatabaseResponse): unknown {
    const obj: any = {};
    if (message.hostsDatabase) {
      obj.hostsDatabase = message.hostsDatabase.map((e) => e ? HostsDatabase.toJSON(e) : undefined);
    } else {
      obj.hostsDatabase = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllHostsDatabaseResponse>, I>>(
    object: I,
  ): QueryAllHostsDatabaseResponse {
    const message = createBaseQueryAllHostsDatabaseResponse();
    message.hostsDatabase = object.hostsDatabase?.map((e) => HostsDatabase.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a HostsDatabase by index. */
  HostsDatabase(request: QueryGetHostsDatabaseRequest): Promise<QueryGetHostsDatabaseResponse>;
  /** Queries a list of HostsDatabase items. */
  HostsDatabaseAll(request: QueryAllHostsDatabaseRequest): Promise<QueryAllHostsDatabaseResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.HostsDatabase = this.HostsDatabase.bind(this);
    this.HostsDatabaseAll = this.HostsDatabaseAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("runos_chain.configstore.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  HostsDatabase(request: QueryGetHostsDatabaseRequest): Promise<QueryGetHostsDatabaseResponse> {
    const data = QueryGetHostsDatabaseRequest.encode(request).finish();
    const promise = this.rpc.request("runos_chain.configstore.Query", "HostsDatabase", data);
    return promise.then((data) => QueryGetHostsDatabaseResponse.decode(new _m0.Reader(data)));
  }

  HostsDatabaseAll(request: QueryAllHostsDatabaseRequest): Promise<QueryAllHostsDatabaseResponse> {
    const data = QueryAllHostsDatabaseRequest.encode(request).finish();
    const promise = this.rpc.request("runos_chain.configstore.Query", "HostsDatabaseAll", data);
    return promise.then((data) => QueryAllHostsDatabaseResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
