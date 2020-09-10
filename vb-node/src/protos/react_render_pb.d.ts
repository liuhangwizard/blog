// package: react_render
// file: react_render.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class RenderRequest extends jspb.Message { 
    getStore(): string;
    setStore(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RenderRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RenderRequest): RenderRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RenderRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RenderRequest;
    static deserializeBinaryFromReader(message: RenderRequest, reader: jspb.BinaryReader): RenderRequest;
}

export namespace RenderRequest {
    export type AsObject = {
        store: string,
    }
}

export class RenderResponse extends jspb.Message { 
    getHtml(): string;
    setHtml(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RenderResponse.AsObject;
    static toObject(includeInstance: boolean, msg: RenderResponse): RenderResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RenderResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RenderResponse;
    static deserializeBinaryFromReader(message: RenderResponse, reader: jspb.BinaryReader): RenderResponse;
}

export namespace RenderResponse {
    export type AsObject = {
        html: string,
    }
}
