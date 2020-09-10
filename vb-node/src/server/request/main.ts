//interface
import { AxiosInstance } from "axios";

//interface
//获取主页文章信息
export interface IVBRequestParamMainRecent{
    page:number,
    pageSize:number,
    title?:string,
    catagory?:string,
    tag?:string
}
export const requestMainRecent=(request:AxiosInstance,data:IVBRequestParamMainRecent)=>{
    return request({
        url:"/blog/mainrecent",
        method:"post",
        data
    })
}

//获取主页云信息
export const requestMainCloud=(request:AxiosInstance)=>{
    return request({
        url:"/blog/maincloud",
        method:"post",
    })
}

//获取主页统计
export const requestMainCount=(request:AxiosInstance)=>{
    return request({
        url:"/recordblog/maincount",
        method:"post",
    })
}

