//interface
import { IVBReducerList } from '../define';
import { IVBStoreArticle } from '../../data/define';
import { IVBAction } from '../../action/define';

//action types
import * as Action from '../../action/type-article';

//reducer
import VBReducerArticleHTML from './html';
import VBReducerArticlePerface from './perface';
import VBReducerArticlePropose from './propose';

//reducer list
const reducerList:IVBReducerList={
    [Action.GROUP_ARTICLE_HTML]:VBReducerArticleHTML,
    [Action.GROUP_ARTICLE_PERFACE]:VBReducerArticlePerface,
    [Action.GROUP_ARTICLE_PROPOSE]:VBReducerArticlePropose,
}


//entry
const VBReducerArticle=(store:IVBStoreArticle|undefined,action:IVBAction<any>):IVBStoreArticle=>{
    
    const state=store as any;
    const group=action.group;
    
    //new
    const newStore=(group&& action.type)?{
        ...state,
        [group]:reducerList[group](state[group],action)
    }:store;

    return newStore;
}


export default VBReducerArticle;