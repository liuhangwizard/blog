//react
import React, { Fragment } from 'react';

//scss
import "src/style/article/center.scss";

//component
//component detail
import VBArticleDetail from '../detail';

//component panel
import VBArticlePropose from '../propose';
import VBArticleDirectory from '../directory';


const VBArticleCenter = (): JSX.Element => {

    

    return (
        <div className="vb-article-center vb-flexstart">
            <div className="center-detail">
                <VBArticleDetail/>
            </div>
            <div className="center-panel">
                <VBArticleDirectory/>
                <VBArticlePropose/>
            </div>
        </div>
    )
}


export default VBArticleCenter;