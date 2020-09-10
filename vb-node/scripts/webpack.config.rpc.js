//node
const path = require('path');
const merge = require('webpack-merge');
//const nodeExternals = require('webpack-node-externals'); //独立保留require进来的包 不打包进去
//common
const base = require('./webpack.config.base.js');

const rpcConfig = {
    target: "node",
    entry: "./src/rpc/index.ts",
    output: {
        filename: "rpc_server.js",
        path: path.join(__dirname, "../", "dist")
    },
    module: {
        rules: [{
            test: /\.scss$/,
            loader: "ignore-loader",
            exclude: /node_modules/,
        }]
    }
}

module.exports = merge(base, rpcConfig);