//react
import React, { useEffect } from 'react';

//scss
import 'src/style/main/preview-archive.scss';

//component
import VBFocus from '../../../components/focus';

//hooks
import { useSelector } from 'react-redux';

//utils
import { formateDate } from '../../../../utils/date';

//config
import { serverConfig } from '../../../../server/config';

//interface
import { IVBStoreMain, IVBMainCloudArchive } from '../../../../store/data/define';
import { RouteConfigComponentProps } from 'react-router-config';


interface IVBMainPreviewArchiveProps extends RouteConfigComponentProps<any>{

}

interface IArchiveInfo{
    [id:string]:IVBMainCloudArchive[]
}

const VBMainPreviewArchive = (props:IVBMainPreviewArchiveProps): JSX.Element => {


    //props
    const url=props.location.pathname;
    const docTitle=props.route.title;
    const focusName=props.route.focus;

    //hooks
    const archiveList=useSelector((store:IVBStoreMain)=>store.cloud.archiveList);

    //ui
    const info:IArchiveInfo={};
    Object.values(archiveList).forEach((archive,i)=>{
        const date=new Date(archive.create_date).getFullYear();
        !info.hasOwnProperty(date)&&(info[date]=[]);
        info[date].push(archive);
    });

    const archiveUIList=Object.keys(info).reverse().map((year,yearIndex)=>{
        const current=info[year] as IVBMainCloudArchive[];
        const currentUI=current.map((article,articleIndex)=>(
            <li key={articleIndex} className="c-item">
                <a className="i-link" href={serverConfig.domain+"/article/"+article.id} target="_blank" >
                    <span className="i-date">{formateDate("YYYY-mm-dd",new Date(article.create_date))}</span>
                    <span className="i-title">{article.title}</span>              
                </a>
            </li>
        ))

        return (
            <div key={yearIndex} className="a-content common-item board-content">
                <p className="c-title">{year}</p>
                <ul className="c-list">
                    {currentUI}
                </ul>
            </div>
        )
    });

    useEffect(()=>{
        document.title=docTitle;
    },[docTitle])

    return (
        <div className="prview-container  vb-a-fadin">
            <VBFocus url={url} center={focusName}/>
            <div className="preview-archive">
                {archiveUIList}
            </div>
        </div>
    )
}





export default VBMainPreviewArchive;