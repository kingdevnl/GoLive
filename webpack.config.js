const path = require('path');

module.exports = {
    entry: './js/golive.ts',
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
        ],
    },
    resolve: {
        extensions: ['.tsx', '.ts', '.js'],
    },
    output: {
        filename: 'golive.js',
        path: path.resolve(__dirname, './js/dist'),
    },
};