//store
import { Store} from "redux";


//remote dispatch
import { remoteMainCloud, remoteMainCount,remoteMainRecent} from "../../store/remote/main";
import { mainLocation } from "../../store/dispatch/main";

//utils
import { promiseList } from "./promise";

//interface
import { IVBStoreMain } from "../../store/data/define";


export const InitMainData=(store: Store<IVBStoreMain,any>,path:string,param:Object)=>{


    let dataList=[
        remoteMainCloud,
        remoteMainCount,
        remoteMainRecent("/",param),
        mainLocation(path,param),
    ];

    if(path==="/search"){
        dataList.push(remoteMainRecent(path,param),);
    }


    return promiseList(store,dataList);
}