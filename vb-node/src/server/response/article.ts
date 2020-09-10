import { IVBResponseCommonPropose } from "./common";

export interface IVBResponseArticleInfo{
    id:number,
    title:string,
    image:string,
    status:string,
    summary:string,
    catagory:string,
    directory:string,
    create_date:string,
    tags:string,
    record:number,
    html:string,
    propose:IVBResponseCommonPropose[]
}
