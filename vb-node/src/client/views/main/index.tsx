//react
import React  from 'react';
import ReactDOM from 'react-dom';

//store
import { Provider } from 'react-redux';
import { GetStoreMain } from '../../../store/data';

//component
import VBMain from './page';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css'

//init
NProgress.configure({showSpinner:false})
NProgress.start();


//store render
const store = GetStoreMain("client");
ReactDOM.hydrate((
    <Provider store={store}>
        <VBMain/>
    </Provider>
),document.getElementById('wizard-main-LH'));

//loader

window.onload=(e:Event)=>NProgress.done()


export default VBMain;



