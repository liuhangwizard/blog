//interface
import { IVBAction } from '../../../action/define';
import { IVBReducerProcessList } from '../../define';
import { IVBActionDataArticleHTMLSet } from '../../../action/data-article';

//action types
import * as Action from '../../../action/type-article'



//process
const HTMLProcessSet=(state:string,data:IVBActionDataArticleHTMLSet):string=>{
    return data.html;
}


//reducer process
const processList:IVBReducerProcessList={
    [Action.ACTION_ARTICLE_HTML_SET]:HTMLProcessSet,
}

//reducer
const VBReducerArticleHTML=(state:string,action:IVBAction<any>):string=>{

    //action
    const type=action.type;
    const data=action.data;

    //process
    return (processList.hasOwnProperty(type))?processList[type](state,data):state
}

export default VBReducerArticleHTML;
