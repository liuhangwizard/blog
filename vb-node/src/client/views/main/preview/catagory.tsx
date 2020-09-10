//react
import React, { useMemo, useState, useEffect } from 'react';

//scss
import 'src/style/main/preview-catagory.scss';

//hooks
import { useRequest } from '@umijs/hooks'

//store
import { mainRecent, mainSearch, mainLocation } from '../../../../store/dispatch/main';
//import { useHistory } from "react-router-dom";

//component
import VBFocus from '../../../components/focus';
import NProgress from 'nprogress';
//store
import { useSelector, useDispatch } from 'react-redux';


//request
import { requestMainRecent, IVBRequestParamMainRecent } from '../../../../server/request/main';
import { clientAxios } from '../../../../server/request';


//interface
import { IVBStoreMain, IVBMainCloudArchive, IVBMainLocation } from '../../../../store/data/define';
import { RouteConfigComponentProps } from 'react-router-config';



interface ICatagroyInfo {
    [id: string]: string[]
}

interface ICatagoryListProps {
    archiveList: IVBMainCloudArchive[],
    loading: boolean,
    activeIndex: number,
    onItemClick: (index: number, catagory: string) => void
}


interface IVBMainPreviewCatagoryProps extends RouteConfigComponentProps<any>{

}

const remoteCatagory = (catagory: string) => {

    const req: IVBRequestParamMainRecent = {
        page: 1,
        pageSize: 5,
        catagory
    }

    return requestMainRecent(clientAxios, req)
}

const CatagoryList = (props: ICatagoryListProps): JSX.Element => {

    //props
    const{archiveList,loading,activeIndex,onItemClick}=props;


    const info: ICatagroyInfo = {};
    Object.values(archiveList).forEach((archive, i) => {
        const catagory = archive.catagory;
        !info.hasOwnProperty(catagory) && (info[catagory] = []);
        info[catagory].push(archive.title);
    });

    const catagoryUI = Object.keys(info).map((catagory, i) => (
        <div key={i} className="c-item">
            <div className={"board-content i-center"+(loading&&(i!==activeIndex)?" i-center-shadow":"")}>
                <div 
                className="i-title"
                onClick={e=>!loading&&onItemClick(i,catagory)}>
                    {catagory}
                </div>
                <div className="i-count">{info[catagory].length}</div>
            </div>
        </div>
    ));

    return (
        <React.Fragment>{catagoryUI}</React.Fragment>
    )

}
const VBMainPreviewCatagory = (props:IVBMainPreviewCatagoryProps): JSX.Element => {

    //props
    const url=props.location.pathname;
    const docTitle=props.route.title;
    const focusName=props.route.focus;

    //hooks
    const archiveList = useSelector((store: IVBStoreMain) => store.cloud.archiveList);
    const location=useSelector((store:IVBStoreMain)=>store.location);
    const [active, setActive] = useState<number>(-1);
    const dispatch=useDispatch();


    const {loading, run } = useRequest(remoteCatagory, {
        manual: true,
        loadingDelay: 200,
        onSuccess: (res, params) => {
            if(res.data?.code===20000){
                const result=res.data.data;
                const list=result.list;
                const total=result.total;
                dispatch(mainLocation(location.current,{"catagory":params[0]}))
                dispatch(mainSearch({total,list,"page":2}))
                props.history.push({
                    pathname:'/search',
                    search:"?catagory="+params[0]
                })
            }
            setActive(-1);
            NProgress.done();
        },
        onError: (error, params) => { 
            setActive(-1);
            NProgress.done();
        }
    });


    const cUIList = useMemo(() => (
        <CatagoryList 
        loading={loading}
        activeIndex={active}
        onItemClick={(i,c)=>{
            NProgress.start();
            setActive(v=>i);
            run(c);
        }}
        archiveList={Object.values(archiveList)} />
    ), [archiveList,active,loading]);

    useEffect(()=>{
        document.title=docTitle;
    },[docTitle])


    return (
        <div className="prview-container  vb-a-fadin">
            <VBFocus url={url} center={focusName}/>
            <div className="preview-catagory">
                {cUIList}
            </div>
        </div>
    )
}

export default VBMainPreviewCatagory;