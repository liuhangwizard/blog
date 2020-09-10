//react
import React from 'react';

//store
import { useSelector } from 'react-redux';

//component
import VBSection from '../../../components/section';

//interface
import { IVBStoreMain } from '../../../../store/data/define';

// const recentList=[
//     {id:2,title:"Webpack从零开始的React同构实践",catagory:"前端",create_date:"3月前",record:1},
//     {id:2,title:"Nginx反向代理与负载均衡实践",catagory:"后台",create_date:"1年前",record:1},
//     {id:2,title:"React项目Fiber原理调试",catagory:"前端",create_date:"1年前",record:1},
//     {id:2,title:"Vue+Nuxt.js环境搭建,一个超级深的BUG解决和分析过程",catagory:"前端",create_date:"1年前",record:1},
//     {id:2,title:"自己动手写一个Webpack插件",catagory:"前端",create_date:"5月前",record:1},
// ];


const VBMainRecent = (): JSX.Element => {


    //hooks
    const propose=useSelector((store:IVBStoreMain)=>store.propose);
    const proposeList=Object.values(propose);


    return (
       <VBSection 
       name="文章" 
       icon="icon-article" 
       sectionList={proposeList}></VBSection>
    )
}


export default VBMainRecent;