//interface
import { AxiosInstance } from "axios";


//获取文章正文内容
export interface IVBRequestParamArticleInfo{
    id:number
}
export const requestArticleInfo=(request:AxiosInstance,data:IVBRequestParamArticleInfo)=>{
    return request({
        url:"/blog/artilceinfo",
        method:"post",
        data
    })
}


export interface IVBRequestParamArticleRecord{
    blog_id:number,
    ip:string,
    city:string,
    city_id:string,
    client:string
}
export const requestArticleRecord=(request:AxiosInstance,data:IVBRequestParamArticleRecord)=>{
    return request({
        url:"/recordblog/add",
        method:"post",
        data
    })
}
