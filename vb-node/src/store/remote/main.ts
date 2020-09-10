//store
import { Dispatch } from "redux";

//request
import { requestMainCloud, requestMainCount, requestMainRecent, IVBRequestParamMainRecent } from "../../server/request/main"
import { mainCloud, mainCount, mainPropose, mainRecent, mainSearch } from "../dispatch/main";

//interface
import { AxiosInstance } from "axios";
import { IVBMainCloudArchive, IVBMainCloudTag } from "../data/define";
import { IVBResponseMainCount } from "../../server/response/main";
import { IVBResponseCommonPropose } from "../../server/response/common";


export const remoteMainCloud=(dispatch: Dispatch<any>, getStore: () => any, req: AxiosInstance)=>{

    const p=requestMainCloud(req);
    p.then((res)=>{
        if(res.data?.code===20000){
            const result=res.data?.data;

            const archive=result.archive as IVBMainCloudArchive[];
            const tag=result.tag as IVBMainCloudTag[];

            dispatch(mainCloud(tag,archive));
            
        }
    });
    return p;
}


export const remoteMainCount=(dispatch: Dispatch<any>, getStore: () => any, req: AxiosInstance)=>{

    const p=requestMainCount(req);
    p.then((res)=>{
        if(res.data?.code===20000){

            const result=res.data?.data;
            const info=result?.count as IVBResponseMainCount;
            const propose=result?.propose as IVBResponseCommonPropose;

            const {article,tag,record,avatar,back}=info;
            dispatch(mainCount(article,tag,record,avatar,back));
            dispatch(mainPropose(propose))
            
        }

    });
    return p;
}


export const remoteMainRecent=(path:string,param:any)=>(dispatch: Dispatch<any>, getStore: () => any, req: AxiosInstance)=>{

    let recentData:IVBRequestParamMainRecent={
        page:1,
        pageSize:5,
    }

    if(path==="/search"){
        recentData={
            ...recentData,
            title:param?.title,
            tag:param?.tag,
            catagory:param?.catagory
        }
    }

    const p=requestMainRecent(req,recentData);
    p.then((res)=>{

        if(res.data?.code===20000){
            const list=res.data.data.list;
            const total=res.data.data.total;

            if(path==="/search"){
                dispatch(mainSearch({total,list,page:2}))
                return;
            }
            dispatch(mainRecent({total,list,page:2}))
        }
    })

    return p;
}
