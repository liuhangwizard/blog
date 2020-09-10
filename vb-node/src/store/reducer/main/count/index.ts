//interface
import { IVBAction } from '../../../action/define';
import { IVBMainCount } from '../../../data/define';
import { IVBReducerProcessList } from '../../define';
import { IVBActionDataMainCountSet } from '../../../action/data-main'

//action types
import * as Action from '../../../action/type-main';

//process
const CountProcessSet=(state:IVBMainCount,data:IVBActionDataMainCountSet):IVBMainCount=>{
    

    return ({
        ...data
    })
}


//reducer process
const processList:IVBReducerProcessList={
    [Action.ACTION_MAIN_COUNT_SET]:CountProcessSet,
}

//reducer
const VBReducerMainCount=(state:IVBMainCount,action:IVBAction<any>):IVBMainCount=>{

    //action
    const type=action.type;
    const data=action.data;

    //process
    return (processList.hasOwnProperty(type))?processList[type](state,data):state
}

export default VBReducerMainCount;
