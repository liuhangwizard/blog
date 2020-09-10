//react
import React, { useEffect, useState, useRef } from 'react';

//scss
import "src/style/main/preview-home.scss";

//hooks
import { useRequest } from '@umijs/hooks';

//store
import { useSelector, useDispatch } from 'react-redux';
import { mainRecent } from '../../../../store/dispatch/main';

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
import { IVBActionDataMainRecentSet } from '../../../../store/action/data-main';




interface IVBMainPreviewHomeProps extends RouteConfigComponentProps<any>{

}


const onHomeList = (page: number) => {


    let data: IVBRequestParamMainRecent = {
        page,
        pageSize: 5,
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

const VBMainPreviewHome = (props:IVBMainPreviewHomeProps): JSX.Element => {


    //props
    const url=props.location.pathname;
    const docTitle=props.route.title;

    //hooks
    //默认已经加载了第一页 从第二页开始 store的page已经初始化为了2
    const recent = useSelector((store: IVBStoreMain) => store.recent);
    const dispatch=useDispatch();
    const recentList=Object.values(recent.list);

    const { loading, run } = useRequest(() => onHomeList(recent.page), {
        manual: true,
        onSuccess: (result, params) => {
            const newRecent:IVBActionDataMainRecentSet={
                list:[...recentList,...result.list],
                total:recent.total,
                page:recent.page+1
            }
            dispatch(mainRecent(newRecent));
            NProgress.done();
        },
        onError: (error, params) => NProgress.done()
    });

    useEffect(()=>{
        docTitle&&(document.title=docTitle);
    },[docTitle])


    const onLoad=()=>{
        NProgress.start();
        run();
    }

    //ui 
    const contentUI = recentList.map((item, i) => (<VBPreviewContent key={i} content={item} />));


    return (
        <div className="preview-home vb-a-fadin">
            <VBFocus url={url}  title="全部文章"/>
            <InfiniteScroll
            hasMore={recentList.length < recent.total}
            dataLength={recentList.length}
            loader={<VBToast content="正在加载..."/>}
            endMessage={<VBToast content="已经到达了世界终点..."/>}
            next={onLoad}>
                {contentUI}
            </InfiniteScroll>
        </div>
    )
}


export default VBMainPreviewHome;