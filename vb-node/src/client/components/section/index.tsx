//react
import React, { useMemo } from 'react';

//scss
import 'src/style/common/section.scss';

//component
import VBPanel from '../panel';

import { serverConfig } from '../../../server/config';

//utils
import { formatDateTime } from '../../../utils/date';
import { formatRecord } from '../../../utils/record';


//interface
import { IVBCommonConnect } from '../../../store/data/define';




interface IVBSectionData extends IVBCommonConnect{

}



interface IVBSectionProps{
    name:string,
    icon:string,
    sectionList:IVBSectionData[]
}


const VBSection = (props:IVBSectionProps): JSX.Element => {


    const{name,icon,sectionList}=props;


    const sectionListUI:JSX.Element[]=useMemo(()=>sectionList.map((item,index)=>(

        <div key={index} className="recent-item">
            <p className="r-title"><a href={serverConfig.domain+"/article/"+item.id} target="_blank">{item.title}</a></p>
            <p className="r-basic">
                <a className="b-item iconfont icon-view"><span>{formatRecord(item.record)}</span></a>
                <a className="b-item iconfont icon-category"><span>{item.catagory}</span></a>
                <a className="b-item iconfont icon-date"><span>{formatDateTime(item.create_date)}</span></a>
            </p>
        </div>
        
    )),[sectionList]);


    return (
        <VBPanel name="vb-section" title={name} icon={icon}>
             {sectionListUI}
        </VBPanel>
    )
}


export default VBSection;