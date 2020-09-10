//interface
import { IVBReducerList } from '../define';
import { IVBStoreMain } from '../../data/define';
import { IVBAction } from '../../action/define';

//action types
import * as Action from '../../action/type-main'

//reducer
import VBReducerMainCloud from './cloud';
import VBReducerMainCount from './count';
import VBReducerMainRecent from './recent';
import VBReducerMainPropose from './propose';
import VBReducerMainLocation from './location';


//reducer list
const reducerList:IVBReducerList={
    [Action.GROUP_MAIN_CLOUD]:VBReducerMainCloud,
    [Action.GROUP_MAIN_COUNT]:VBReducerMainCount,
    [Action.GROUP_MAIN_PROPOSE]:VBReducerMainPropose,
    [Action.GROUP_MAIN_LOCATION]:VBReducerMainLocation,

    //recent search 结构相同 复用相同的reducer
    [Action.GROUP_MAIN_RECENT]:VBReducerMainRecent,
    [Action.GROUP_MAIN_SEARCH]:VBReducerMainRecent,
}



//entry
const VBReducerMain=(store:IVBStoreMain|undefined,action:IVBAction<any>):IVBStoreMain=>{
    
    const state=store as any;
    const group=action.group;

    
    //new
    const newStore=(group && action.type)?{
        ...state,
        [group]:reducerList[group](state[group],action)
    }:store;

    return newStore;
}


export default VBReducerMain;