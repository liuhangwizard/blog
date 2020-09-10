//utils
export interface IVBList<T>{
    [id:string]:T
}

//common setting
export interface IVBCommonSetting{
    name:string
}



//common connect
export interface IVBCommonConnect{
    id:number,
    title:string,
    catagory:string,
    create_date:string,
    record:number
}

//main
//路由地址 包括首页 归档 分类 搜索[包括参数]
export interface IVBMainLocation{
    current:string
    param:any
}



//统计面板 包括头像和背景
export interface IVBMainCount{
    article:number,
    tag:number,
    record:number,
    avatar:string,
    back:string,
}

//巫师基本信息
export interface IVBMainBasic{
    email:string,
    sina:string,
    wechat:string,
    github:string,
    zhihu:string,
}

//云标签信息
export interface IVBMainCloudTag{
    id:number,
    name:string,
}
//云分类信息(暂未使用)
// export interface IVBMainCloudCatagory{
//     id:number,
//     content:string
// }
//云归档信息
export interface IVBMainCloudArchive extends IVBCommonConnect{

}

//云信息
export interface IVBMainCloud{
    tagList:IVBList<IVBMainCloudTag>,
    // catagoryList:IVBList<IVBMainCloudCatagory>
    archiveList:IVBList<IVBMainCloudArchive>
}

//id 标题 日期 作者 查看次数 图片 简介 (首页的)描述 分类 标签合集
export interface IVBMainPreface{
    id:number
    title:string,
    author:string,
    catagory:string,
    tags:string,//用逗号隔开
    image:string,
    summary:string,
    record:number,
    create_date:string,
}

//page不同
export interface IVBMainRecent{
    list:IVBList<IVBMainPreface>
    total:number,
    page:number,
}

export interface IVBMainSearch extends IVBMainRecent{

}

//main data
export interface IVBStoreMain{
    setting:IVBCommonSetting,
    location:IVBMainLocation,
    count:IVBMainCount,
    basic:IVBMainBasic,
    cloud:IVBMainCloud,
    propose:IVBList<IVBCommonConnect>,
    recent:IVBMainRecent
    search:IVBMainSearch
}

//article
//文章简介
export interface IVBArticlePerface extends IVBMainPreface{
    directory:string
}
//article data
export interface IVBStoreArticle{
    setting:IVBCommonSetting,
    perface:IVBArticlePerface,
    html:string,
    propose:IVBList<IVBCommonConnect>
}