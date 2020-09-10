//interface
import { IVBAction } from '../../../action/define';
import { IVBCommonConnect,IVBList } from '../../../data/define';
import { IVBReducerProcessList } from '../../define';
import { IVBActionDataMainProposeSet } from '../../../action/data-main'

//action types
import * as Action from '../../../action/type-main';

//process
const ProposeProcessSet=(state:IVBList<IVBCommonConnect>,data:IVBActionDataMainProposeSet):IVBList<IVBCommonConnect>=>{
    
    
    let newProposeList:IVBList<IVBCommonConnect>={}     
    for(let i of data){
        newProposeList={
            ...newProposeList,
            [i.id]:i
        }
    }

    return newProposeList;
}


//reducer process
const processList:IVBReducerProcessList={
    [Action.ACTION_MAIN_PROPOSE_SET]:ProposeProcessSet,
}

//reducer
const VBReducerMainPropose=(state:IVBList<IVBCommonConnect>,action:IVBAction<any>):IVBList<IVBCommonConnect>=>{

    //action
    const type=action.type;
    const data=action.data;

    //process
    return (processList.hasOwnProperty(type))?processList[type](state,data):state
}

export default VBReducerMainPropose;
