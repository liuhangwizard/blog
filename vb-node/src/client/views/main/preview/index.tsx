//react
import React from 'react';

//scss
import 'src/style/main/preview.scss';
import 'src/style/common/animation.scss';

//router
import { RouteConfigComponentProps, renderRoutes } from 'react-router-config';

//component
//component header
import VBHeader from '../../../components/header';

//component link
import VBLink from '../../../components/link';

//component panel
import VBMainBoard from '../board';
import VBMainBasic from '../basic';
import VBMainPropose from '../propose';
import VBMainTag from '../tag';


interface IVBPreviewProps extends RouteConfigComponentProps {

}

const VBPreview = (props: IVBPreviewProps): JSX.Element => {

    const routes = props.route?.routes ?? [];

    return (
        <div className="vb-main">
            <VBHeader />
            <div className="vb-main-page vb-flexstart">
                <div className="page-preview">
                    <div className="vb-main-preview">
                        {renderRoutes(routes)}
                    </div>
                </div>
                <div className="page-panel">
                    <VBMainBoard />
                    <VBMainBasic />
                    <VBMainPropose />
                    <VBMainTag />
                </div>
            </div>
            <VBLink/>
        </div>
    )
}
export default VBPreview;