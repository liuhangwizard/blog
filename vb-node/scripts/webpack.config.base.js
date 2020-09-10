//node 
const path = require('path');

//setting
const setting = require('./setting');

const baseConfig = {
    mode: setting.devMode ? "development" : "production",
    // devtool: setting.devMode ? "source-map" : "none",
    resolve: {
        extensions: ['.js', '.jsx', '.ts', '.tsx'],
        alias: {
            src: path.join(__dirname, "../", "src")
        }
    },
    module: {
        rules: [{
            test: /\.(js|jsx|ts|tsx)$/,
            loader: "babel-loader",
            exclude: /node_modules/,
            options: {
                babelrc: false,
                extends: './scripts/.babelrc'
            }
        }]
    }
}

module.exports = baseConfig;