const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const package = require('./package.json')

module.exports = {
    entry: package.main,
    output: {
        path: path.resolve(__dirname,"../web/public/"),
        filename: 'ui-browser.js',
        publicPath: "/public"
    },
    module: {
        rules: [
            {
                test: /\.js?$/,
                exclude: /node_modules/,
                use: 'babel-loader'
            },
            {
                test: /\.s?css$/,
                use: [
                    "style-loader", "css-loader", "sass-loader"
                ]
            }
        ]
    },
    devServer: {
        port: 8084,
        inline: true,
        publicPath:"http://localhost:8080/public",
        historyApiFallback: true,
        // overlay: true,
        proxy: {
            "*": {
                context: ['/public'],
                target: 'http://localhost:8080',
                // // pathRewrite: {"^/public" : ''},
                // secure: false,
                ws: true, /// auto update
                bypass: function (req, res, proxyOptions) {
                    console.log(req.originalUrl)
                    if (req.originalUrl === '/public/ui-browser') {
                        return '/ui-browser';
                    }
                }
            },
            '/logger': {
                target: 'ws://localhost:8080/',
                ws: false,
                secure: false,
                sockPort: 8080,
                onProxyReq: function (proxyReq) {
                    proxyReq.setHeader("Accept-Encoding", '') // disable gzip
                    proxyReq.setHeader("Access-Control-Allow-Origin","*")
                }
             }
        }
    }
}