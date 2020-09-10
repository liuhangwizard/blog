//react
import React, { useMemo, useRef, useEffect } from 'react';

//store
import { useSelector } from 'react-redux';

//scss
import 'src/style/article/directory.scss';

//component
import VBPanel from '../../../components/panel';


//interface
import { IVBStoreArticle } from '../../../../store/data/define';


const scrollToTaget=(content:string,hList:React.MutableRefObject<HTMLHeadingElement[]>)=>{
    for(let h of hList.current){

        if(h.textContent.indexOf(content)>-1){
            const toTop=h.offsetTop;
            window.scrollTo({
                top:toTop-60,
                behavior:"smooth",
            })
            break;
        }
    }
}


const VBArticleDirectory = (): JSX.Element => {

    //hooks
    const directory=useSelector((store:IVBStoreArticle)=>store.perface.directory)
    const artilceHList=useRef<HTMLHeadingElement[]>([]);
    

    useEffect(()=>{
        const article=document.getElementById('vb-a-main');
        const hList=Array.from<HTMLHeadingElement>(article.querySelectorAll('h1,h2,h3,h4,h5,h6'))
        artilceHList.current=hList;
    },[])



    //ui
    const directoryList=directory.indexOf(',')>-1?directory.split(','):[];
    const directoryUI=directoryList.map((title,index)=>(
        <p  key={index} className="d-title" onClick={e=>scrollToTaget(title,artilceHList)}>
            {title}
        </p>
    ));


    return (
        <VBPanel name="vb-article-directory" title="目录" icon="icon-navigate">
            {directoryUI}
        </VBPanel>
    )
}


export default VBArticleDirectory;