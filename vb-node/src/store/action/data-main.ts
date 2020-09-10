//interface
import { IVBMainCloudTag,IVBMainCloudArchive, IVBMainLocation} from '../data/define'
import { IVBResponseMainCount, IVBResponseMainRecent } from '../../server/response/main';
import { IVBResponseCommonPropose } from '../../server/response/common';

//设置cloud 包括 标签 分类 归档内容
export interface IVBActionDataMainCloudSet{
    tagList:IVBMainCloudTag[],
    archiveList:IVBMainCloudArchive[]
}

//设置当前路径
export interface IVBActionDataMainLocationSet extends IVBMainLocation{
    
}

//设置统计信息-基本信息
export interface IVBActionDataMainCountSet extends IVBResponseMainCount{
    
}

//设置统计信息-文章推荐
export interface IVBActionDataMainProposeSet extends IVBResponseCommonPropose{
    
}

//设置主文章
export interface IVBActionDataMainRecentSet extends IVBResponseMainRecent{
    page:number,
}

export interface IVBActionDataMainSearchSet extends IVBResponseMainRecent{
    page:number,
}