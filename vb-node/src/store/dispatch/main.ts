
//store
import { VBDispatch } from '.';

//interface
import { IVBAction } from '../action/define';
import { IVBMainCloudTag, IVBMainCloudArchive } from '../data/define';
import { 
    IVBActionDataMainCloudSet, 
    IVBActionDataMainCountSet, 
    IVBActionDataMainLocationSet, 
    IVBActionDataMainProposeSet, 
    IVBActionDataMainRecentSet,
    IVBActionDataMainSearchSet,  } from '../action/data-main';

//action
import * as Action from '../action/type-main';
import { IVBResponseCommonPropose } from '../../server/response/common';



//设置主页cloud
export const mainCloud=(tagList:IVBMainCloudTag[],archiveList:IVBMainCloudArchive[])=>{

    const action: IVBAction<IVBActionDataMainCloudSet> = {
        type: Action.ACTION_MAIN_CLOUD_SET,
        group: Action.GROUP_MAIN_CLOUD,
        data:{
            tagList,
            archiveList
        }
    };

    return VBDispatch(action);
}


//设置主页文章
export const mainRecent=(recent:IVBActionDataMainRecentSet)=>{
    
    const action:IVBAction<IVBActionDataMainRecentSet>={
        type:Action.ACTION_MAIN_RECENT_SET,
        group:Action.GROUP_MAIN_RECENT,
        data:recent
    }

    return VBDispatch(action);
}

//设置搜索结果
export const mainSearch=(search:IVBActionDataMainSearchSet)=>{

    const action:IVBAction<IVBActionDataMainSearchSet>={
        type:Action.ACTION_MAIN_SEARCH_SET,
        group:Action.GROUP_MAIN_SEARCH,
        data:search
    }

    return VBDispatch(action);

}

//设置主页统计-基本信息
export const mainCount=(article:number,tag:number,record:number,avatar:string,back:string)=>{

    const action:IVBAction<IVBActionDataMainCountSet>={
        type:Action.ACTION_MAIN_COUNT_SET,
        group:Action.GROUP_MAIN_COUNT,
        data:{
            article,
            tag,
            record,
            avatar,
            back
        }
    }

    return VBDispatch(action);
}

//设置主页统计-最近文章
export const mainPropose=(propose:IVBResponseCommonPropose)=>{

    const action:IVBAction<IVBActionDataMainProposeSet>={
        type:Action.ACTION_MAIN_PROPOSE_SET,
        group:Action.GROUP_MAIN_PROPOSE,
        data:propose
    }
    
    return VBDispatch(action);
}


//设置主页当前的路由
export const mainLocation=(current:string,param:Object)=>{

    const action:IVBAction<IVBActionDataMainLocationSet>={
        type:Action.ACTION_MAIN_LOCATION_SET,
        group:Action.GROUP_MAIN_LOCATION,
        data:{
            current,
            param
        }
    }


    return VBDispatch(action);
}