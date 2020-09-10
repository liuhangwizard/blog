
//interface
import { IVBCommonConnect,IVBArticlePerface } from '../data/define'


export interface IVBActionDataArticleHTMLSet{
    html:string
}


export interface IVBActionDataArticleProposeSet{
    proposeList:IVBCommonConnect[]
}

export interface IVBActionDataArticlePerfaceSet extends IVBArticlePerface{

}