//node
const path = require('path');

//plugins
const merge = require('webpack-merge');
const HTMLWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const FileWebpackPlugin = require('../scripts/plugins/FileWebpackPlugin');

//
//const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;

//setting
const base = require('./webpack.config.base.js');
const setting = require('./setting');


//config
const clientConfig = {
    entry: {
        main: "./src/client/views/main/index.tsx",
        article: "./src/client/views/article/index.tsx"
    },
    output: {
        publicPath:"/",
        filename: "js/[name]_bundle.js",
        path: path.join(__dirname, "../", "build")
    },
    module: {
        rules: [
            {
                test: /\.scss$/,
                use: ExtractTextPlugin.extract({
                    fallback: 'style-loader',
                    use: ['css-loader', 'sass-loader']
                })
            },
            {
                test: /\.css$/,
                use: ExtractTextPlugin.extract({
                    fallback: 'style-loader',
                    use: ['css-loader']
                })
            },
        ]
    },
    plugins: [
        new HTMLWebpackPlugin({
            filename: "index.html",
            chunks: ['main', 'commons'],
            template: path.join(__dirname, "../", "public", "main.html")
        }),
        new HTMLWebpackPlugin({
            filename: "article.html",
            chunks: ['article', 'commons'],
            template: path.join(__dirname, "../", "public", "article.html")
        }),
        new ExtractTextPlugin({
            filename: 'css/[name].css'
        }),
        new FileWebpackPlugin({
            log: true,
            fileList: {
                "js/commons_bundle.js": [path.join(__dirname, "../", "dist")],
                "js/main_bundle.js": [path.join(__dirname, "../", "dist")],
                "js/article_bundle.js": [path.join(__dirname, "../", "dist")],
                "css/commons.css": [path.join(__dirname, "../", "dist")],
                "css/main.css": [path.join(__dirname, "../", "dist")],
                "css/article.css": [path.join(__dirname, "../", "dist")],
            }
        }),
        //new BundleAnalyzerPlugin()
    ],
    optimization: {
        splitChunks: {
            cacheGroups: {
                commons: {
                    name: "commons",
                    chunks: "all",
                    minChunks: 2,
                }
            }
        }
    },
    devServer: {
        contentBase: path.join(__dirname, "../", "build"),
        //historyApiFallback: true, 研究下这个
        port: setting.port.client,
        hot: true,
        overlay: true //全屏显示错误
    }
}

module.exports = merge(base, clientConfig);