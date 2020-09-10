//interface
import { IVBAction } from '../../../action/define';
import { IVBMainPreface, IVBList, IVBMainRecent } from '../../../data/define';
import { IVBReducerProcessList } from '../../define';
import { IVBActionDataMainRecentSet } from '../../../action/data-main'

//action types
import * as Action from '../../../action/type-main';

//process
const RecentProcessSet=(state:IVBMainRecent,data:IVBActionDataMainRecentSet):IVBMainRecent=>{
            
    const{list,total,page}=data;
    let newState:IVBMainRecent={
        list:{},
        total:total,
        page:page
    }
    for(let i of list){
        newState.list={
            ...newState.list,
            [i.id]:i
        }
    }

    return newState;
}


//reducer process
const processList:IVBReducerProcessList={
    [Action.ACTION_MAIN_RECENT_SET]:RecentProcessSet,

    //search和recent复用
    [Action.ACTION_MAIN_SEARCH_SET]:RecentProcessSet
}

//reducer
const VBReducerMainRecent=(state:IVBMainRecent,action:IVBAction<any>):IVBMainRecent=>{

    //action
    const type=action.type;
    const data=action.data;

    //process
    return (processList.hasOwnProperty(type))?processList[type](state,data):state
}

export default VBReducerMainRecent;
