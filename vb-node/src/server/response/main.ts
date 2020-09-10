import { IVBMainPreface } from "../../store/data/define";

export interface IVBResponseMainCount{
    tag:number,
    record:number,
    article:number,
    back:string,
    avatar:string
}

export interface IVBResponseMainRecent{
    list:IVBMainPreface[],
    total:number
}