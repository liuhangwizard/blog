//react
import React, { useEffect } from 'react';

//scss
import "src/style/article/detail.scss";

//store
import { useSelector } from 'react-redux';

//request
import { clientAxios } from '../../../../server/request';
import { requestArticleRecord, IVBRequestParamArticleRecord } from '../../../../server/request/article';

//component
import VBFocus from '../../../components/focus';


//interface
import { IVBStoreArticle } from '../../../../store/data/define';





const VBArticleDetail = (): JSX.Element => {

    const contentHTML=useSelector((store:IVBStoreArticle)=>store.html);
    const perface=useSelector((store:IVBStoreArticle)=>store.perface);

    useEffect(()=>{
        var IPInfo = {"cip": "未知", "cid": "未知", "cname": "未知"};
        IPInfo={
            ...(window as any).returnCitySN
        }
        const recordData:IVBRequestParamArticleRecord={
            client:navigator?.userAgent??"未知",
            blog_id:perface.id,
            ip:IPInfo.cip,
            city:IPInfo.cname,
            city_id:IPInfo.cid
        }
        requestArticleRecord(clientAxios,recordData).then(res=>{
            //console.log("响应",res.data);
        }).catch(e=>{});

    },[])



    return (
        <div className="vb-article-detail">
            <VBFocus url="/article" title={perface.title}/>
            <div id="vb-a-main" className="text-content board-content" dangerouslySetInnerHTML={{__html:contentHTML}}/>
        </div>
    )
}

export default VBArticleDetail;