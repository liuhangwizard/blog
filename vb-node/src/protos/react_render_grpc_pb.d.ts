// package: react_render
// file: react_render.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as react_render_pb from "./react_render_pb";

interface IRenderServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getReactHTMLContent: IRenderServiceService_IgetReactHTMLContent;
}

interface IRenderServiceService_IgetReactHTMLContent extends grpc.MethodDefinition<react_render_pb.RenderRequest, react_render_pb.RenderResponse> {
    path: string; // "/react_render.RenderService/getReactHTMLContent"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<react_render_pb.RenderRequest>;
    requestDeserialize: grpc.deserialize<react_render_pb.RenderRequest>;
    responseSerialize: grpc.serialize<react_render_pb.RenderResponse>;
    responseDeserialize: grpc.deserialize<react_render_pb.RenderResponse>;
}

export const RenderServiceService: IRenderServiceService;

export interface IRenderServiceServer {
    getReactHTMLContent: grpc.handleUnaryCall<react_render_pb.RenderRequest, react_render_pb.RenderResponse>;
}

export interface IRenderServiceClient {
    getReactHTMLContent(request: react_render_pb.RenderRequest, callback: (error: grpc.ServiceError | null, response: react_render_pb.RenderResponse) => void): grpc.ClientUnaryCall;
    getReactHTMLContent(request: react_render_pb.RenderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: react_render_pb.RenderResponse) => void): grpc.ClientUnaryCall;
    getReactHTMLContent(request: react_render_pb.RenderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: react_render_pb.RenderResponse) => void): grpc.ClientUnaryCall;
}

export class RenderServiceClient extends grpc.Client implements IRenderServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getReactHTMLContent(request: react_render_pb.RenderRequest, callback: (error: grpc.ServiceError | null, response: react_render_pb.RenderResponse) => void): grpc.ClientUnaryCall;
    public getReactHTMLContent(request: react_render_pb.RenderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: react_render_pb.RenderResponse) => void): grpc.ClientUnaryCall;
    public getReactHTMLContent(request: react_render_pb.RenderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: react_render_pb.RenderResponse) => void): grpc.ClientUnaryCall;
}
