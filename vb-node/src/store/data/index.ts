
//store
import { createStore, applyMiddleware } from 'redux'
import { vbMain, vbArticle } from './init';

//reducer
import VBReducerMain from '../reducer/main';
import VBReducerArtilce from '../reducer/article';

//request
import { clientAxios, serverAxios } from '../../server/request';

//middleware
import thunk from 'redux-thunk';
import { composeWithDevTools } from 'redux-devtools-extension';



//interface
import { IVBStoreMain, IVBStoreArticle } from './define';



type TVBStoreType="client"|"server"

const storeUtil=(type:TVBStoreType,init:IVBStoreMain|IVBStoreArticle)=>{
    const thunkMiddleware=(type==="client")?thunk.withExtraArgument(clientAxios):thunk.withExtraArgument(serverAxios);
    const initStore=(type==="client")?(window as any).INITIAL_DATA:init;

    //const initStore=tempMainData as any;

    return {
        thunkMiddleware,
        initStore
    }
}

export const GetStoreMain = (type:TVBStoreType) => {

    const utils=storeUtil(type,vbMain);
    const store = createStore(VBReducerMain, utils.initStore, composeWithDevTools(applyMiddleware(utils.thunkMiddleware)));
    
    return store;

}

export const GetStoreArticle = (type:TVBStoreType) => {

    const utils=storeUtil(type,vbArticle);
    const store = createStore(VBReducerArtilce, utils.initStore, composeWithDevTools(applyMiddleware(utils.thunkMiddleware)));
    
    return store;
}


