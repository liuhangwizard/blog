//react
import React, { useEffect, useState, useRef } from 'react';

//scss
import "src/style/main/preview-home.scss";

//hooks
import { useRequest } from '@umijs/hooks';

//store
import { useSelector, useDispatch } from 'react-redux';
import { mainSearch } from '../../../../store/dispatch/main';

//request
import { requestMainRecent, IVBRequestParamMainRecent } from '../../../../server/request/main';
import { clientAxios } from '../../../../server/request';



//component
import VBFocus from '../../../components/focus';
import VBToast from '../../../components/toast';
import InfiniteScroll from 'react-infinite-scroll-component';
import VBPreviewContent from '../content';
import NProgress from 'nprogress';


//interface
import { IVBStoreMain } from '../../../../store/data/define';
import { RouteConfigComponentProps } from 'react-router-config';
import { IVBActionDataMainSearchSet } from '../../../../store/action/data-main';




interface IVBMainPreviewHomeProps extends RouteConfigComponentProps<any>{

}


//search home分离 减少耦合
const onHomeList = (page: number,param:any) => {


    let data: IVBRequestParamMainRecent = {
        page,
        pageSize: 5,
        title:param?.title,
        catagory:param?.catagory,
        tag:param?.tag
    }

    return requestMainRecent(clientAxios, data).then(res => {
        const result = res?.data?.data;
        const total = result?.total ?? 0;
        const list = result?.list ?? [];

        return ({
            total,
            list
        })

    })
}

const VBMainPreviewSearch = (props:IVBMainPreviewHomeProps): JSX.Element => {


    //props
    const url=props.location.pathname;
    const docTitle=props.route.title;

    //hooks
    //默认已经加载了第一页 从第二页开始
    const search=useSelector((store:IVBStoreMain)=>store.search);
    const param=useSelector((store:IVBStoreMain)=>store.location.param);
    const dispatch=useDispatch();
    const searchList=Object.values(search.list);
    

    const { loading, run } = useRequest(() => onHomeList(search.page,param), {
        manual: true,
        onSuccess: (result, params) => {
            const newSearch:IVBActionDataMainSearchSet={
                list:[...searchList,...result.list],
                total:search.total,
                page:search.page+1
            }
            dispatch(mainSearch(newSearch));
            NProgress.done();
        },
        onError: (error, params) => NProgress.start()
    });

    useEffect(()=>{
        docTitle&&(document.title=docTitle);
    },[docTitle])

    const onLoad=()=>{
        NProgress.start();
        run();
    }

    //ui

    const contentUI = searchList.map((item, i) => (<VBPreviewContent key={i} content={item} />));

    const endMessage=search.total>0?"已经到达了世界终点...":"没有找到你想要的东西哦...";
   
    let showFocusCenter=null;
    let showFocsuType=null;
    let showTitle="全部文章";

    if(param.catagory){
        showFocusCenter=param.catagory;
        showFocsuType="catagory";
    }

    if(param.tag){
        showFocusCenter=param.tag;
    }

    if(param.title){
        showTitle=param.title+"的搜索结果"
    }


    return (
        <div className="preview-home vb-a-fadin">
            <VBFocus url={url} center={showFocusCenter} type={showFocsuType} title={showTitle}/>
            <InfiniteScroll
            hasMore={searchList.length < search.total}
            dataLength={searchList.length}
            loader={<VBToast content="正在加载..."/>}
            endMessage={<VBToast content={endMessage}/>}
            next={onLoad}>
                {contentUI}
            </InfiniteScroll>
        </div>
    )
}


export default VBMainPreviewSearch;