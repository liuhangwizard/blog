//react
import React, { useMemo } from 'react';

//scss
import 'src/style/main/tag.scss';

//config
import { serverConfig } from '../../../../server/config';

//store
import { useSelector } from 'react-redux';

//component
import VBPanel from '../../../components/panel';

//interface
import { IVBStoreMain } from '../../../../store/data/define';


const VBMainTag = (props:any): JSX.Element => {


    //hooks
    const tagList=useSelector((store:IVBStoreMain)=>store.cloud.tagList);

    //ui
    const tagListUI=useMemo(()=>Object.values(tagList).map((tag,i)=>(
        <a href={serverConfig.domain+"/search?tag="+tag.name} className="tag-item" key={"tag-id-"+tag.id}>{tag.name}</a>
    )),[tagList]);


    return (
        <VBPanel name="vb-main-tag" title="标签" icon="icon-tag">
            {tagListUI}
        </VBPanel>
    )
}


export default VBMainTag;