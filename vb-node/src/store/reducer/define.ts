import { IVBAction } from "../action/define";

export interface IVBReducerList{
    [id:string]:(state:any,action:IVBAction<any>)=>any
}

export interface IVBReducerProcessList{
    [id:string]:(state:any,data:any)=>any
} 