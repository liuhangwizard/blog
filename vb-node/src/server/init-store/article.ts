//store
import { Store} from "redux";

//interface
import { IVBStoreArticle } from "../../store/data/define";

//utils
import { promiseList } from "./promise";

//remote
import { remoteArticleInfo } from "../../store/remote/article";


export const InitArticleData=(store: Store<IVBStoreArticle,any>,artilceId:number)=>{

    const dataList=[
        remoteArticleInfo(artilceId)
    ];

    return promiseList(store,dataList);
}