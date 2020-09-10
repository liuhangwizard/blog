//react
import React from 'react';

//store
import { useSelector } from 'react-redux';

//component
import VBSection from '../../../components/section';


//interface
import { IVBStoreArticle } from '../../../../store/data/define';




const VBArticlePropose = (): JSX.Element => {

    //hooks
    const propose=useSelector((store:IVBStoreArticle)=>store.propose)
    const proposeList=Object.values(propose);

    //ui
    const proposeUI=proposeList.length>0?(
        <VBSection 
        name="推荐" 
        icon="icon-article" 
        sectionList={proposeList}/>
    ):null;

    return proposeUI;
}


export default VBArticlePropose;