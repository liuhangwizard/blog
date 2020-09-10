import { IVBStoreMain, IVBStoreArticle } from "./define";


export const vbMain: IVBStoreMain = {
    setting:{
        name:"main",
    },
    location:{
        current:"",
        param:{},
    },
    count:{
        article:-1,
        tag:-1,
        record:-1,
        avatar:"",
        back:""
    },
    basic: {
        email: "854732152@qq.com",
        zhihu: "西门不吹",
        github: "yechongwizard",
        sina:"星与微光",
        wechat:"星与微光",
    },
    cloud: {
        tagList:{},
        archiveList:{}
    },
    propose:{},
    recent: {
        list:{},
        total:0,
        page:0,
    },
    search:{
        list:{},
        total:0,
        page:0,
    }
};

export const vbArticle:IVBStoreArticle={
    setting:{
        name:"article"
    },
    perface: {
        id: -1,
        author: "",
        image: "",
        summary: "",
        title: "",
        catagory:"",
        tags:"",
        record: 0,
        create_date: "",
        directory:"",
    },
    html: "",
    propose:{}
}

