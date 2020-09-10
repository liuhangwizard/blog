//react
import React from 'react';
import ReactDOMServer from 'react-dom/server'

//store
import { Provider } from 'react-redux';
import { Store } from 'redux';

//component
import VBMain from '../../client/views/main/page';
import VBArticle from '../../client/views/article/page';


//interface
import { IVBStoreArticle, IVBStoreMain } from '../../store/data/define';



//article
export const GetArticleHTMLContent = (store:Store<IVBStoreArticle, any>): string => {
    
    return ReactDOMServer.renderToString(
        <Provider store={store}>
            <VBArticle />
        </Provider>
    );
}


export const GetArticleHTMLNodeStream = (): NodeJS.ReadableStream => {

    return ReactDOMServer.renderToNodeStream(<VBArticle />);
}

//main
export const GetMainHTMLContent = (store:Store<IVBStoreMain,any>): string => {

    return ReactDOMServer.renderToString(
        <Provider store={store}>
            <VBMain />
        </Provider>
    );
}


export const GetMainHTMLNodeStream = (): NodeJS.ReadableStream => {

    return ReactDOMServer.renderToNodeStream(<VBMain />);
}

