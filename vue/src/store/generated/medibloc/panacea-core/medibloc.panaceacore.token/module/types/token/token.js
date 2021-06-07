/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "medibloc.panaceacore.token";
const baseToken = {
    creator: "",
    id: 0,
    name: "",
    symbol: "",
    totalSupply: 0,
    mintable: false,
    ownerAddress: "",
};
export const Token = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.id !== 0) {
            writer.uint32(16).uint64(message.id);
        }
        if (message.name !== "") {
            writer.uint32(26).string(message.name);
        }
        if (message.symbol !== "") {
            writer.uint32(34).string(message.symbol);
        }
        if (message.totalSupply !== 0) {
            writer.uint32(40).int32(message.totalSupply);
        }
        if (message.mintable === true) {
            writer.uint32(48).bool(message.mintable);
        }
        if (message.ownerAddress !== "") {
            writer.uint32(58).string(message.ownerAddress);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseToken };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.name = reader.string();
                    break;
                case 4:
                    message.symbol = reader.string();
                    break;
                case 5:
                    message.totalSupply = reader.int32();
                    break;
                case 6:
                    message.mintable = reader.bool();
                    break;
                case 7:
                    message.ownerAddress = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseToken };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.symbol !== undefined && object.symbol !== null) {
            message.symbol = String(object.symbol);
        }
        else {
            message.symbol = "";
        }
        if (object.totalSupply !== undefined && object.totalSupply !== null) {
            message.totalSupply = Number(object.totalSupply);
        }
        else {
            message.totalSupply = 0;
        }
        if (object.mintable !== undefined && object.mintable !== null) {
            message.mintable = Boolean(object.mintable);
        }
        else {
            message.mintable = false;
        }
        if (object.ownerAddress !== undefined && object.ownerAddress !== null) {
            message.ownerAddress = String(object.ownerAddress);
        }
        else {
            message.ownerAddress = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        message.name !== undefined && (obj.name = message.name);
        message.symbol !== undefined && (obj.symbol = message.symbol);
        message.totalSupply !== undefined &&
            (obj.totalSupply = message.totalSupply);
        message.mintable !== undefined && (obj.mintable = message.mintable);
        message.ownerAddress !== undefined &&
            (obj.ownerAddress = message.ownerAddress);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseToken };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.symbol !== undefined && object.symbol !== null) {
            message.symbol = object.symbol;
        }
        else {
            message.symbol = "";
        }
        if (object.totalSupply !== undefined && object.totalSupply !== null) {
            message.totalSupply = object.totalSupply;
        }
        else {
            message.totalSupply = 0;
        }
        if (object.mintable !== undefined && object.mintable !== null) {
            message.mintable = object.mintable;
        }
        else {
            message.mintable = false;
        }
        if (object.ownerAddress !== undefined && object.ownerAddress !== null) {
            message.ownerAddress = object.ownerAddress;
        }
        else {
            message.ownerAddress = "";
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
