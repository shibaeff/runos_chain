/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "runos_chain.configstore";

export interface MsgSetPort {
  creator: string;
  dpid: number;
  mac: string;
  inPort: number;
}

export interface MsgSetPortResponse {
}

function createBaseMsgSetPort(): MsgSetPort {
  return { creator: "", dpid: 0, mac: "", inPort: 0 };
}

export const MsgSetPort = {
  encode(message: MsgSetPort, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.dpid !== 0) {
      writer.uint32(16).uint64(message.dpid);
    }
    if (message.mac !== "") {
      writer.uint32(26).string(message.mac);
    }
    if (message.inPort !== 0) {
      writer.uint32(32).uint64(message.inPort);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetPort {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetPort();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.dpid = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.mac = reader.string();
          break;
        case 4:
          message.inPort = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetPort {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      dpid: isSet(object.dpid) ? Number(object.dpid) : 0,
      mac: isSet(object.mac) ? String(object.mac) : "",
      inPort: isSet(object.inPort) ? Number(object.inPort) : 0,
    };
  },

  toJSON(message: MsgSetPort): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.dpid !== undefined && (obj.dpid = Math.round(message.dpid));
    message.mac !== undefined && (obj.mac = message.mac);
    message.inPort !== undefined && (obj.inPort = Math.round(message.inPort));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetPort>, I>>(object: I): MsgSetPort {
    const message = createBaseMsgSetPort();
    message.creator = object.creator ?? "";
    message.dpid = object.dpid ?? 0;
    message.mac = object.mac ?? "";
    message.inPort = object.inPort ?? 0;
    return message;
  },
};

function createBaseMsgSetPortResponse(): MsgSetPortResponse {
  return {};
}

export const MsgSetPortResponse = {
  encode(_: MsgSetPortResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetPortResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetPortResponse();
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

  fromJSON(_: any): MsgSetPortResponse {
    return {};
  },

  toJSON(_: MsgSetPortResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetPortResponse>, I>>(_: I): MsgSetPortResponse {
    const message = createBaseMsgSetPortResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetPort(request: MsgSetPort): Promise<MsgSetPortResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.SetPort = this.SetPort.bind(this);
  }
  SetPort(request: MsgSetPort): Promise<MsgSetPortResponse> {
    const data = MsgSetPort.encode(request).finish();
    const promise = this.rpc.request("runos_chain.configstore.Msg", "SetPort", data);
    return promise.then((data) => MsgSetPortResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}