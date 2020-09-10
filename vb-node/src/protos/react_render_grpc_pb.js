// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var react_render_pb = require('./react_render_pb.js');

function serialize_react_render_RenderRequest(arg) {
    if (!(arg instanceof react_render_pb.RenderRequest)) {
        throw new Error('Expected argument of type react_render.RenderRequest');
    }
    return Buffer.from(arg.serializeBinary());
}

function deserialize_react_render_RenderRequest(buffer_arg) {
    return react_render_pb.RenderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_react_render_RenderResponse(arg) {
    if (!(arg instanceof react_render_pb.RenderResponse)) {
        throw new Error('Expected argument of type react_render.RenderResponse');
    }
    return Buffer.from(arg.serializeBinary());
}

function deserialize_react_render_RenderResponse(buffer_arg) {
    return react_render_pb.RenderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var RenderServiceService = exports.RenderServiceService = {
    getReactHTMLContent: {
        path: '/react_render.RenderService/getReactHTMLContent',
        requestStream: false,
        responseStream: false,
        requestType: react_render_pb.RenderRequest,
        responseType: react_render_pb.RenderResponse,
        requestSerialize: serialize_react_render_RenderRequest,
        requestDeserialize: deserialize_react_render_RenderRequest,
        responseSerialize: serialize_react_render_RenderResponse,
        responseDeserialize: deserialize_react_render_RenderResponse,
    },
};

exports.RenderServiceClient = grpc.makeGenericClientConstructor(RenderServiceService);