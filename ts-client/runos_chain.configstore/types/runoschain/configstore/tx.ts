/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "runos_chain.configstore";

export interface MsgSetPort {
  creator: string;
  dpid: string;
  mac: string;
  inport: string;
}

export interface MsgSetPortResponse {
}

function createBaseMsgSetPort(): MsgSetPort {
  return { creator: "", dpid: "", mac: "", inport: "" };
}

export const MsgSetPort = {
  encode(message: MsgSetPort, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.dpid !== "") {
      writer.uint32(18).string(message.dpid);
    }
    if (message.mac !== "") {
      writer.uint32(26).string(message.mac);
    }
    if (message.inport !== "") {
      writer.uint32(34).string(message.inport);
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
          message.dpid = reader.string();
          break;
        case 3:
          message.mac = reader.string();
          break;
        case 4:
          message.inport = reader.string();
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
      dpid: isSet(object.dpid) ? String(object.dpid) : "",
      mac: isSet(object.mac) ? String(object.mac) : "",
      inport: isSet(object.inport) ? String(object.inport) : "",
    };
  },

  toJSON(message: MsgSetPort): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.dpid !== undefined && (obj.dpid = message.dpid);
    message.mac !== undefined && (obj.mac = message.mac);
    message.inport !== undefined && (obj.inport = message.inport);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetPort>, I>>(object: I): MsgSetPort {
    const message = createBaseMsgSetPort();
    message.creator = object.creator ?? "";
    message.dpid = object.dpid ?? "";
    message.mac = object.mac ?? "";
    message.inport = object.inport ?? "";
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
