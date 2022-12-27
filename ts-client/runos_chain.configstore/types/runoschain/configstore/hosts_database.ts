/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "runos_chain.configstore";

export interface HostsDatabase {
  dpid: string;
  mac: string;
  inport: string;
  creator: string;
}

function createBaseHostsDatabase(): HostsDatabase {
  return { dpid: "", mac: "", inport: "", creator: "" };
}

export const HostsDatabase = {
  encode(message: HostsDatabase, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.dpid !== "") {
      writer.uint32(10).string(message.dpid);
    }
    if (message.mac !== "") {
      writer.uint32(18).string(message.mac);
    }
    if (message.inport !== "") {
      writer.uint32(26).string(message.inport);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HostsDatabase {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHostsDatabase();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.dpid = reader.string();
          break;
        case 2:
          message.mac = reader.string();
          break;
        case 3:
          message.inport = reader.string();
          break;
        case 4:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HostsDatabase {
    return {
      dpid: isSet(object.dpid) ? String(object.dpid) : "",
      mac: isSet(object.mac) ? String(object.mac) : "",
      inport: isSet(object.inport) ? String(object.inport) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: HostsDatabase): unknown {
    const obj: any = {};
    message.dpid !== undefined && (obj.dpid = message.dpid);
    message.mac !== undefined && (obj.mac = message.mac);
    message.inport !== undefined && (obj.inport = message.inport);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<HostsDatabase>, I>>(object: I): HostsDatabase {
    const message = createBaseHostsDatabase();
    message.dpid = object.dpid ?? "";
    message.mac = object.mac ?? "";
    message.inport = object.inport ?? "";
    message.creator = object.creator ?? "";
    return message;
  },
};

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
