//interface
import { IVBAction } from '../../../action/define';
import { IVBMainCloud, IVBList, IVBMainCloudTag, IVBMainCloudArchive } from '../../../data/define';
import { IVBReducerProcessList } from '../../define';
import { IVBActionDataMainCloudSet } from '../../../action/data-main'

//action types
import * as Action from '../../../action/type-main';

//process
const CloudProcessSet=(state:IVBMainCloud,data:IVBActionDataMainCloudSet):IVBMainCloud=>{
    

    //tag
    let newTagList:IVBList<IVBMainCloudTag>={}     
    for(let i of data.tagList){
        newTagList={
            ...newTagList,
            [i.id]:i
        }
    }

    //catagory
    // let newCatagoryList:IVBList<IVBMainCloudCatagory>={}
    // for(let j of data.catagoryList){
    //     newCatagoryList={
    //         ...newCatagoryList,
    //         [j.id]:j
    //     }
    // }

    //archive
    let newArchiveList:IVBList<IVBMainCloudArchive>={}
    for(let k of data.archiveList){
        newArchiveList={
            ...newArchiveList,
            [k.id]:k
        }
    }

    
    return ({
        tagList:newTagList,
        archiveList:newArchiveList
    })
}


//reducer process
const processList:IVBReducerProcessList={
    [Action.ACTION_MAIN_CLOUD_SET]:CloudProcessSet,
}

//reducer
const VBReducerMainCloud=(state:IVBMainCloud,action:IVBAction<any>):IVBMainCloud=>{

    //action
    const type=action.type;
    const data=action.data;

    //process
    return (processList.hasOwnProperty(type))?processList[type](state,data):state
}

export default VBReducerMainCloud;
