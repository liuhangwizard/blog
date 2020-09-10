//config
import setting from '../../scripts/setting';

//rpc
import * as grpc from '@grpc/grpc-js';
import reactMessage from '../protos/react_render_pb';
import reactService from '../protos/react_render_grpc_pb';

//react


function getReactHTMLContent(call: any, callback: any) {

    //console.log("V",call);
    //content
    const content="<div>Temp</div>"
    //reply
    const reply = new reactMessage.RenderResponse();
    reply.setHtml(content);
    callback(null, reply);
}


function RPCServer() {

    const rpcAddress=setting.IP+":"+setting.port.rpc;
    const server = new grpc.Server();
    server.addService(reactService.RenderServiceService as any, { "getReactHTMLContent": getReactHTMLContent });
    server.bindAsync(rpcAddress, grpc.ServerCredentials.createInsecure(),(port,error)=>{
        console.log("rpc server on ",rpcAddress);
        server.start();
    });

}

RPCServer();
