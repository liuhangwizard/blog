//router
import { RouteConfig } from 'react-router-config'


//component
import VBPage from '../views/main/preview';
import VBMainPreviewHome from '../views/main/preview/home'
import VBMainPreviewCatagory from '../views/main/preview/catagory';
import VBMainPreviewArchive from '../views/main/preview/archive';
import VBMainPreviewSearch from '../views/main/preview/search';



//interface
interface IVBMainRoute extends RouteConfig {
    title?:string,
    focus?:string,
}


//routes
const name="星与微光-刘行的博客";
const routes: IVBMainRoute[] = [
    {
        path: "/",
        key: "main",
        exact: false,
        component: VBPage,
        routes: [
            {
                path: "/",
                key: "home",
                title:"首页 | "+name,
                foucs:"首页",
                exact: true,
                component: VBMainPreviewHome,
            },
            {
                path: "/search",
                key: "search",
                title:"搜索 | "+name,
                focus:"搜索",
                exact: true,
                component: VBMainPreviewSearch
                //复用组件
            },
            {
                path: "/catagory",
                key: "catagory",
                title:"分类 | "+name,
                focus:"分类",
                exact: true,
                component: VBMainPreviewCatagory
            },
            {
                path: "/archive",
                key: "archive",
                title:"归档 | "+name,
                focus:"归档",
                exact: true,
                component: VBMainPreviewArchive
            }
        ]
    },
];

export default routes;