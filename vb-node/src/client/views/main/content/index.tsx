//react
import React, { useMemo } from 'react';

//scss
import 'src/style/main/content.scss';

//store
import { useSelector } from 'react-redux';

//config
import { serverConfig } from '../../../../server/config';

//utils
import { formatRecord } from '../../../../utils/record';
import { formateDate } from '../../../../utils/date';

//interface
import { IVBMainPreface, IVBStoreMain } from '../../../../store/data/define';




interface IVBMainContentProps {
    content: IVBMainPreface
}

//preview
const VBPreviewContent = (props: IVBMainContentProps): JSX.Element => {

    //props
    const { id,title, create_date, author, record, image, summary, catagory, tags } = props.content;




    //data
    //avatar
    
    const avatar = serverConfig.domainStatic+"/blog/image/avatar.jpg";
    //const tempImage="/static/abc.jpg";
    const visitShow=formatRecord(record);

    const imageAddress=serverConfig.domainStatic+"/"+image;

    //render
    const tagList=tags.split(',');
    const tagUIList=useMemo(()=>tagList.map((item,index)=>(
        <a key={index}>#{item??"通用"}</a>
    )),[tagList]);


    //render final
    return (
        <div className="vb-preview-content common-item board-content">
            <div className="c-author">
                <div className="au-avatar"><img className="a-image" src={avatar}></img></div>
                <div className="au-basic">
                    <span className="b-name">LiuHang</span>
                    <div className="b-date">{formateDate("YYYY-mm-dd",new Date(create_date))}</div>
                </div>
            </div>
            <div className="c-article vb-flexbetween">
                <a href={serverConfig.domain+"/article/"+id} target="_blank">
                    <div className="a-title">{title}</div>
                </a>
                <div className="a-info">
                    <a className="i-item iconfont icon-view"><span>{visitShow}</span></a>
                    <a className="i-item iconfont icon-category"><span>{catagory}</span></a>
                </div>
            </div>
            <div className="c-main">
                <div className="m-back" style={{ backgroundImage: "url(" + imageAddress + ")" }}></div>
                <div className="m-text">
                    <div className="t-content">{summary}</div>
                    <div className="t-tag vb-flexstartrow"><i className="iconfont icon-tag"></i>{tagUIList}</div>
                </div>
            </div>
        </div>
    )
}


export default VBPreviewContent;