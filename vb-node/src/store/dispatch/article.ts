
//store
import { VBDispatch } from '.';

//action
import * as Action from '../action/type-article'


//interface
import { IVBAction } from '../action/define';
import { IVBCommonConnect } from '../data/define';
import { 
    IVBActionDataArticlePerfaceSet, 
    IVBActionDataArticleProposeSet, 
    IVBActionDataArticleHTMLSet 
} from '../action/data-article';



//设置文章信息
export const articleInfo = (id:number,title:string,author:string,catagory:string,tags:string,image:string,summary:string,directory:string,record:number,create_date:string) => {

    const action: IVBAction<IVBActionDataArticlePerfaceSet> = {
        type: Action.ACTION_ARTICLE_PERFACE_SET,
        group: Action.GROUP_ARTICLE_PERFACE,
        data: {
            id:id,
            title:title,
            author:author,
            catagory:catagory,
            tags:tags,
            image:image,
            summary:summary,
            record:record,
            directory:directory,
            create_date:create_date
        }
    };


    return VBDispatch(action);

}

//设置文章推荐信息
export const articlePropose = (p:IVBCommonConnect[])=> {

    const action: IVBAction<IVBActionDataArticleProposeSet> = {
        type: Action.ACTION_ARTICLE_PROPOSE_SET,
        group: Action.GROUP_ARTICLE_PROPOSE,
        data:{
            proposeList:p
        }
    };

    return VBDispatch(action);

}


//设置文章html
export const articleHTML = (html: string) => {

    const action: IVBAction<IVBActionDataArticleHTMLSet> = {
        type: Action.ACTION_ARTICLE_HTML_SET,
        group: Action.GROUP_ARTICLE_HTML,
        data: {
            html: html
        }
    };

    return VBDispatch(action);

}
