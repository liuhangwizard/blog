//react
import React, { useMemo } from 'react';

//component

//scss
import 'src/style/common/panel.scss';



interface IVBPanelProps {
    name: string,
    title: string,
    icon: string,
    children: JSX.Element | JSX.Element[]
}

const VBPanel = (props: IVBPanelProps): JSX.Element => {

    const { name, title, icon, children } = props;

    const titleUI: JSX.Element = (title !== "NULL" ? (
        <div className="panel-title">
            <i className={"iconfont "+icon}></i>
            <span>{title}</span>
        </div>
    ) : null);


    return (
        <div className={"vb-panel board-content " + name}>
            {titleUI}
            <div className="panel-center">
                {children}
            </div>
        </div>
    )
}


export default VBPanel;