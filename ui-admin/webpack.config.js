const path = require('path');
const package = require('./package.json')

module.exports = {
    entry: package.main,
    output: {
        path: path.resolve(__dirname,"../web/public/"),
        filename: 'admin.js',
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
                    if (req.originalUrl === '/public/admin') {
                        return '/admin';
                    }
                }
            }
        }
    }
}