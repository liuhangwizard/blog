//interface
import { IVBAction } from '../../../action/define';
import { IVBMainLocation } from '../../../data/define';
import { IVBReducerProcessList } from '../../define';
import { IVBActionDataMainLocationSet } from '../../../action/data-main'

//action types
import * as Action from '../../../action/type-main';

//process
const LoctaionProcessSet=(state:IVBMainLocation,data:IVBActionDataMainLocationSet):IVBMainLocation=>{
    

    return ({
        "current":data.current,
        "param":data.param,
    })
}


//reducer process
const processList:IVBReducerProcessList={
    [Action.ACTION_MAIN_LOCATION_SET]:LoctaionProcessSet,
}

//reducer
const VBReducerMainLocation=(state:IVBMainLocation,action:IVBAction<any>):IVBMainLocation=>{

    //action
    const type=action.type;
    const data=action.data;

    //process
    return (processList.hasOwnProperty(type))?processList[type](state,data):state
}

export default VBReducerMainLocation;
