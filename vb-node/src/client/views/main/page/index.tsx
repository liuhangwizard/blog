//react
import React from 'react';


//store
import { useSelector } from 'react-redux';


//router
import { StaticRouter, BrowserRouter,Router} from 'react-router-dom';
import { renderRoutes } from 'react-router-config';
import routes from '../../../routes/main';
import VBHistory from '../../../routes/history';


//interface
import { IVBStoreMain } from '../../../../store/data/define';


const VBMain=():JSX.Element=>{

    //hooks
    const location=useSelector((store:IVBStoreMain)=>store.location.current);


    //platform
    const isClient=typeof document==="object";

    //ui
    const routerUI:JSX.Element=isClient?
    (
        <BrowserRouter>
            <Router history={VBHistory as any}>
                {renderRoutes(routes)}
            </Router>
        </BrowserRouter>
    ):
    (
        <StaticRouter location={location} context={{}} >
            {renderRoutes(routes)}
        </StaticRouter>
    )

    return routerUI;

}

export default VBMain