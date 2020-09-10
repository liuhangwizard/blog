//interface
import { IVBAction } from '../../../action/define';
import { IVBReducerProcessList } from '../../define';
import { IVBActionDataArticleProposeSet } from '../../../action/data-article';
import { IVBCommonConnect, IVBList } from '../../../data/define';

//action types
import * as Action from '../../../action/type-article'


//process
const PerfaceProposeSet=(state:IVBList<IVBCommonConnect>,data:IVBActionDataArticleProposeSet):IVBList<IVBCommonConnect>=>{
    
    let newState:IVBList<IVBCommonConnect>={}    
    
    for(let i of data.proposeList){
        newState={
            ...newState,
            [i.id]:i
        }
    }

    return newState

}

//reducer process
const processList:IVBReducerProcessList={
    [Action.ACTION_ARTICLE_PROPOSE_SET]:PerfaceProposeSet,
}


//reducer
const VBReducerArticlePropose=(state:IVBList<IVBCommonConnect>,action:IVBAction<any>):string=>{

    //action
    const type=action.type;
    const data=action.data;

    //process
    return (processList.hasOwnProperty(type))?processList[type](state,data):state
}

export default VBReducerArticlePropose;
