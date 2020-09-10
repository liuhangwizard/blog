//node
import fs from 'fs';

//server
import koa from 'koa';
import koaRouter from 'koa-router';
import koaStatic from 'koa-static';
import koaProxy from 'koa2-proxy-middleware';
//import koaParse from 'koa-bodyparser';

//gzip尝试让nginx来实现
//import koaCompress from 'koa-compress';


//content
import { GetMainHTMLContent,GetArticleHTMLContent,GetMainHTMLNodeStream } from '../common/html-content';


//setting
import setting from '../../scripts/setting';
import { serverConfig } from './config';

//store
import { GetStoreArticle, GetStoreMain } from '../store/data';

//store init
import { InitArticleData } from './init-store/article'
import { InitMainData } from './init-store/main'
import { convertServer } from '../utils/convert';


//server
const main = new koa();
const router = new koaRouter();

//html
const HTMLTemplateMain = fs.readFileSync(setting.directory.client+"/index.html").toString();
const HTMLTemplateArticle=fs.readFileSync(setting.directory.client+"/article.html").toString();


//routes
const name="星与微光-刘行的博客";
const mainRoutes:string | RegExp | (string | RegExp)[]=[
    "/",
    "/catagory",
    "/archive",
    "/search"
];

const mainRoutesTitle={
    "/":"首页 | "+name,
    "/catagory":"分类 | "+name,
    "/archive":"归档 | "+name,
    "/search":"搜索 | "+name,
};

router.get(mainRoutes, async (ctx) => {



    //data
    const serverStore=GetStoreMain("server");

    const dataList=InitMainData(serverStore,ctx.path,convertServer(ctx.query));
    await Promise.all(dataList).catch((error)=>console.log("main error",error));


    //html
    const mainTitle=mainRoutesTitle[ctx.path];
    const contentMain=GetMainHTMLContent(serverStore);
    let html=HTMLTemplateMain.replace(`<div id="wizard-main-LH"></div>`,`<div id="wizard-main-LH">${contentMain}</div>`);
    html=html.replace(`<title>INDEX_TITLE</title>`,`<title>${mainTitle}</title>`);
    html=html.replace(`INITIAL_DATA`,`window.INITIAL_DATA=${JSON.stringify(serverStore.getState())}</script>`);
    ctx.body = html;
    ctx.type = "html";

});

const articleRoutes:string | RegExp | (string | RegExp)[]=[
    "/article/:id"
];

router.get(articleRoutes,async (ctx)=>{



    //id
    const id=Number(ctx.params.id);
    if(isNaN(id)){
        ctx.redirect("/404");
        return;
    }

    //data
    const articleStore=GetStoreArticle("server");
    const dataList=InitArticleData(articleStore,id);
    await Promise.all(dataList).catch((error)=>console.log("article error",error));


    //html
    const articleStoreData=articleStore.getState();
    const articleTitle=(articleStoreData?.perface?.title??"正文")+" | "+name;
    const contentArticle=GetArticleHTMLContent(articleStore);
    let html=HTMLTemplateArticle.replace(`<div id="wizard-article-LH"></div>`,`<div id="wizard-article-LH">${contentArticle}</div>`);
    html=html.replace(`<title>ARTICLE_TITLE</title>`,`<title>${articleTitle}</title>`);
    html=html.replace(`INITIAL_DATA`,`window.INITIAL_DATA=${JSON.stringify(articleStoreData)}</script>`);
    ctx.body = html;
    ctx.type = "html";

});

// router.get('/node', async (ctx) => {
//     ctx.res.write("<!DOCTYPE html><html><head><title>My Page</title></head><body>");
//     ctx.res.write('<div id="wizard-main-LH">'); 
//     const stream=GetMainHTMLNodeStream();
//     //const html=HTMLTemplateMain.replace(`<div id="wizard-main-LH"></div>`,`    <div id="wizard-main-LH">${content}</div>`);
//     //ctx.body = html;
//     stream.pipe(ctx.res, { end: false });
//     stream.on('end', () => {
//         ctx.res.write("</div></body></html>");
//         ctx.res.end();
//     });
// });

router.get('/404',async (ctx)=>{
    ctx.body = "<div>你来到了没有知识的荒芜之地</div>";
    ctx.type = "html";
});


//gzip
// main.use(koaCompress({
//     filter (content_type) {
//         return /text/i.test(content_type)
//     },
//     threshold: 2048,
//     gzip: {
//       flush: require('zlib').Z_SYNC_FLUSH
//     },
//     deflate: {
//       flush: require('zlib').Z_SYNC_FLUSH,
//     },
//     br: false
// }));


//反向代理
//doc
//https://www.npmjs.com/package/http-proxy-middleware
const proxyOptions: koaProxy.Koa2ProxyMiddlewareConfig = {
    targets:{
        "/api/(.*)":{
            target:serverConfig.serverBaseURL,
            changeOrigin: true,
            pathRewrite: {
                '/api': '/',
            },
            logLevel:"silent"
        }
    },
};

main.use(koaProxy(proxyOptions));
//main.use(koaParse({enableTypes:['json', 'form', 'text']}));
main.use(router.routes());
main.use(koaStatic(setting.directory.static, {}));

//start
main.listen(setting.port.server, () => console.log("server on port", setting.port.server));