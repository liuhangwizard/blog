//store
import { Store } from "redux";

//interface
import { IVBStoreArticle,IVBStoreMain } from "../../store/data/define";

export const promiseList=(store: Store<IVBStoreArticle|IVBStoreMain,any>,dataList:Array<any>)=>{
    
    
    const resultList:Promise<any>[]=[];

    for(let i of dataList){
        const pro=new Promise((resolve,reject)=>{
            const p=store.dispatch(i) as any;
            if(p?.finally){
                p.finally(()=>resolve());
                return;
            } 
            resolve();
        })
        resultList.push(pro);   
    }

    return resultList;
}