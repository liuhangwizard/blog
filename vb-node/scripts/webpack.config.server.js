//node
const path = require('path');
const merge = require('webpack-merge');
//const nodeExternals = require('webpack-node-externals'); //独立保留require进来的包 不打包进去
//const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;//分析包插件
const IgnorePlugin=require('webpack').IgnorePlugin;

//common
const base = require('./webpack.config.base.js');

const serverConfig = {
    target: "node",
    entry: "./src/server/index.ts",
    output: {
        filename: "server.js",
        path: path.join(__dirname, "../", "dist")
    },
    // externals:{
    //     'any-promise':'Promise'
    // },
    module: {
        rules: [{
            test: /\.scss$/,
            loader: "ignore-loader",
            exclude: /node_modules/,
        }]
    },
    plugins: [
        //new BundleAnalyzerPlugin()
       // new IgnorePlugin(/\/iconv-loader$/)
    ],
}

module.exports = merge(base, serverConfig);

//提示 /any-promise/register.js 24:18-41 出错的话
//可以去掉相关库的 var Promise=require('any-promise'); 因为node6以上已经支持默认promise了 所以不需要这个了