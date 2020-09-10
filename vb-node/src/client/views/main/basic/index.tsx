//react
import React, { useMemo } from 'react';


//scss
import 'src/style/main/basic.scss';
import VBPanel from '../../../components/panel';

//interface

interface IVBBasicInfo {
    icon: string,
    title: string,
    link: string,
    color: string
}


const basicInfos: IVBBasicInfo[] = [
    { color: "#f6c343", icon: "icon-email", title: "854732152@qq.com", link: "" },
    { color: "#eb3323", icon: "icon-weibo", title: "面向信仰编程", link: "" },
    { color: "#73df81", icon: "icon-wechat", title: "liuhang8147", link: "" },
    { color: "#222222", icon: "icon-github", title: "ScourgeWizard", link: "" },
    { color: "#3486f7", icon: "icon-zhihu", title: "ScourgeWizard", link: "" },
];


const VBMainBasic = (): JSX.Element => {


    const itemUIList = useMemo(() => basicInfos.map((item, index) => (
        <div key={index} className="b-item">
            <i className={"i-icon iconfont " + item.icon} style={{color:item.color}}></i>
            <a className="i-title">{item.title}</a>
        </div>
    )), []);

    return (
        <VBPanel name="vb-main-basic" title="NULL" icon="">
            {itemUIList}
        </VBPanel>
    )
}


export default VBMainBasic;