//react
import React from 'react';

//scss
import 'src/style/common/link.scss';


const VBLink=():JSX.Element=>{


    return(
        <div className="vb-link">
            <div className="l-info">
                <span>版权所有© 前端@LiuHang</span>
                <span> | </span>
                <a href="http://beian.miit.gov.cn">
                    <span>渝ICP备20006339号</span>
                </a>
            </div>
            <div className="l-power">
                <span>Powerd by ReactSSR Redux Koa Node.js Golang MySQL</span>
            </div>
        </div>
    )
}


export default VBLink;