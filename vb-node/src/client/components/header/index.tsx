
//react
import React, { useState, useMemo } from 'react';

//scss
import 'src/style/common/index.scss';

//config
import { serverConfig } from '../../../server/config';

//hooks
import { useSelector } from 'react-redux';

//routes
import { Link } from 'react-router-dom';

//utils
import { convertContent } from '../../../utils/convert';



//interface
import { IVBStoreMain, IVBStoreArticle } from '../../../store/data/define';




const onInputEnter=(e: React.KeyboardEvent<HTMLInputElement>,onSearch:()=>void)=>{
  
    var keyCode = null;
    if(e.which)  
    {
        keyCode = e.which;
    }  
    else if(e.keyCode)   
    {
        keyCode = e.keyCode;
    }  

    if(keyCode == 13)   
    {  
        onSearch();
        return false;  
    }  
    return true;  
}

const VBHeader = (): JSX.Element => {

    //hooks
    const name = useSelector<IVBStoreMain | IVBStoreArticle>((store) => store.setting.name);
    const [search,setSearch]=useState<string>("");
    const [focus, setFocus] = useState<Boolean>(false);


    //ui
    //focus
    const searchFocus = "h-search" + (focus ? " h-focus" : "");
    const searchStatus = (status: boolean) => setFocus(f => status);

    //link
    const linkUI = useMemo(()=>name === "article"?
    (
        <React.Fragment>
            <span className="h-logo"><a href={serverConfig.domain}>星与微光</a></span>
            <ul className="h-tabs">
                <li className="h-item"><a href={serverConfig.domain}>首页</a></li>
                <li className="h-item"><a href={serverConfig.domain+"/catagory"}>分类</a></li>
                <li className="h-item"><a href={serverConfig.domain+"/archive"}>归档</a></li>
            </ul>
        </React.Fragment>
    ): 
    (
        <React.Fragment>
            <span className="h-logo"><Link to="/">星与微光</Link></span>
            <ul className="h-tabs">
                <li className="h-item"><Link to="/">首页</Link></li>
                <li className="h-item"><Link to="/catagory">分类</Link></li>
                <li className="h-item"><Link to="/archive">归档</Link></li>
            </ul>
        </React.Fragment>

    ),[name]);

    //search
    const onSearchClick=()=>{

        window.location.href=serverConfig.domain+"/search?title="+convertContent(search)
    }


    //render
    return (
        <header className="vb-header vb-flexbetween">
            <div className="h-main">
                {linkUI}
            </div>
            <div className={searchFocus}>
                <input
                className="s-content"
                placeholder="搜索文章"
                value={search}
                onKeyPress={e=>onInputEnter(e,onSearchClick)}
                onChange={e=>setSearch(e.target.value)}
                onFocus={()=>searchStatus(true)}
                onBlur={()=>searchStatus(false)}></input>
                <button onClick={onSearchClick} className="s-main">
                    <i className="iconfont icon-search"></i>
                </button>
            </div>
        </header>
    )
}

export default VBHeader;