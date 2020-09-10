//interface
import { IVBAction } from '../../../action/define';
import { IVBReducerProcessList } from '../../define';
import { IVBActionDataArticlePerfaceSet } from '../../../action/data-article';
import { IVBArticlePerface } from '../../../data/define';

//action types
import * as Action from '../../../action/type-article'





//process
const PerfaceProcessSet=(state:IVBArticlePerface,data:IVBActionDataArticlePerfaceSet):IVBArticlePerface=>{
    return ({
        ...state,
        ...data,
    })
}


//reducer process
const processList:IVBReducerProcessList={
    [Action.ACTION_ARTICLE_PERFACE_SET]:PerfaceProcessSet,
}

//reducer
const VBReducerArticlePerface=(state:IVBArticlePerface,action:IVBAction<any>):string=>{

    //action
    const type=action.type;
    const data=action.data;

    //process
    return (processList.hasOwnProperty(type))?processList[type](state,data):state
}

export default VBReducerArticlePerface;
