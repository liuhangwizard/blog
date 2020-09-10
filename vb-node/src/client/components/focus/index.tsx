//react
import React  from 'react';

//store
import { useSelector } from 'react-redux';

//scss
import 'src/style/common/focus.scss';

//config
import { serverConfig } from '../../../server/config';

//component
import { Link } from 'react-router-dom';

//interface
import { IVBStoreMain, IVBStoreArticle } from '../../../store/data/define';


//title 定义
//NULL 不显示
//"" 显示为 '全部文章'

interface IVBFocusProps{
    url:string,
    center?:string,
    type?:string,
    title?:string
}

interface IFocusContentProps{
    name?:string,
    url?:string,
    content?:string
    type?:string
}

const FocusHome=(props:IFocusContentProps):JSX.Element=>{

    //props
    const{ name }=props;

    //ui
    const homeUI=name==="article"
    ?( <a className="f-link" href={serverConfig.domain}>首页</a> ) 
    :( <Link className="f-link" to="/">首页</Link> )

    return (
        <span>{homeUI}</span>
    );
}

const FocusCenter=(props:IFocusContentProps):JSX.Element|null=>{

    //props
    const{ url,name,content,type }=props;

    let showURL="/catagory";
    let showHref=serverConfig.domain+showURL;

    if(type==="tag"){
        showURL="/tag"
        showHref=serverConfig.domain+showURL;
    }



    //ui
    const centerUI:JSX.Element=name==="article"
    ?( <a className="f-link" href={showHref}>{content}</a> )
    :( <Link className="f-link" to={showURL}>{content}</Link> )

    return (
        <React.Fragment>
            <span> &gt; </span>
            <span>{centerUI}</span>
        </React.Fragment>
    )

}

const FocusTitle=(props:IFocusContentProps):JSX.Element|null=>{

    //props
    const{ content }=props;

    //最后一栏无链接
    return (
        <React.Fragment>
            <span> &gt; </span>
            <a>{content}</a>
        </React.Fragment>
    )
}




const VBFocus=(props:IVBFocusProps):JSX.Element=>{

    //props
    const { url,center,type,title}=props;

    //hooks
    const name = useSelector((store:IVBStoreMain | IVBStoreArticle) => store.setting.name);



    //ui
    const centerUI=(center&&center!=="")?(
        (type&&type!=="")
        ?(<FocusCenter name={name} url={url} content={center} type={type}/>)
        :(<FocusTitle content={center}/>)
    ):null;
    const titleUI=(title&&title!=="")?(<FocusTitle content={title}/>):null



    return(
        <div className="vb-focus board-content">
            <span className="f-toast">当前位置:</span>
            <FocusHome name={name}/>
            {centerUI}
            {titleUI}
        </div>
    )
}


export default VBFocus;