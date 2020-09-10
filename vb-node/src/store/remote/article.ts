//store
import { Dispatch } from "redux"

//request
import { requestArticleInfo, IVBRequestParamArticleInfo } from "../../server/request/article"

//store
import { articleHTML, articleInfo, articlePropose } from "../dispatch/article"


//interface
import { AxiosInstance } from "axios"
import { IVBCommonConnect } from "../data/define"
import { IVBResponseArticleInfo } from '../../server/response/article'



export const remoteArticleInfo = (id: number) => (dispatch: Dispatch<any>, getStore: () => any, req: AxiosInstance) => {

    const data: IVBRequestParamArticleInfo = {
        id: id
    }
    const p=requestArticleInfo(req, data);
    p.then(res=>{
        if(res.data?.code!==20000){
            return;
        }
        const info=res.data.data.info as IVBResponseArticleInfo
        const propose=res.data.data.propose as IVBCommonConnect[]

        //article
        dispatch(articleInfo(info.id,info.title,"liuhang",info.catagory,info.tags,info.image,info.summary,info.directory,info.record,info.create_date));  
        dispatch(articleHTML(info.html))
        dispatch(articlePropose(propose)); 

    });
    return p
}
