//react
import React  from 'react';

//scss
import 'src/style/common/toast.scss';

interface IVBToastProps{
    content:string
}

const VBToast = (props:IVBToastProps): JSX.Element => {

    const{ content }=props;

    return (
        <div className="board-content vb-c-toast">
            <span>{content}</span>
        </div>
    )
}

export default VBToast