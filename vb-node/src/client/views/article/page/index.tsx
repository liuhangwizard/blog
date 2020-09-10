//react
import React, { Fragment } from 'react';


//component
import VBLink from '../../../components/link';
import VBHeader from '../../../components/header';
import VBArticleCenter from '../center';

//scss
import "src/style/article/index.scss";


const VBArticle = (): JSX.Element => {
    return (
        <div className="vb-article">
            <VBHeader/>
            <VBArticleCenter/>
            <VBLink/>
        </div>
    )
}


export default VBArticle;