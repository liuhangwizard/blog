//react
import React from 'react';

//scss
import 'src/style/main/board.scss';

//store
import { useSelector } from 'react-redux';

//event
import { disableImageDrag } from '../../../../utils/image';

//utils
import { formatRecord } from '../../../../utils/record';

//interface
import { IVBStoreMain } from '../../../../store/data/define';


const VBMainBoard = (): JSX.Element => {


    //hooks
    const count=useSelector((store:IVBStoreMain)=>store.count);
    const {article,tag,record}=count;


    return (
        <div className="board-content vb-main-board">
            <div className="b-back">
            </div>
            <div className="b-wizard">
                <div className="w-main">
                    <img onDragStart={disableImageDrag} className="m-avatar" src="https://static.lhwizard.com/blog/image/avatar.jpg"></img>
                    <div className="m-name">@LiuHang</div>
                </div>
            </div>
            <div className="b-data">
                <a className="d-item ">
                    <div className="i-value">{article}</div>
                    <div className="i-name">文章</div>
                </a>
                <a className="d-item">
                    <div className="i-value">{tag}</div>
                    <div className="i-name">标签</div>
                </a>
                <a className="d-item">
                    <div className="i-value">{formatRecord(record)}</div>
                    <div className="i-name">阅读</div>
                </a>
            </div>
        </div>
    )
}

export default VBMainBoard;