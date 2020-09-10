//react
import React  from 'react';
import ReactDOM from 'react-dom';

//store
import { Provider } from 'react-redux';
import { GetStoreArticle } from '../../../store/data';
//component
import VBArticle from './page';




const clientStore=GetStoreArticle("client")

ReactDOM.hydrate(
    (<Provider store={clientStore}><VBArticle/></Provider>), 
    document.getElementById('wizard-article-LH')
);
