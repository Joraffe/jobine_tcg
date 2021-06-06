const HtmlWebpackPlugin = require('html-webpack-plugin');
const path = require('path');

module.exports = {
  mode: "development",
  entry: path.resolve(__dirname, 'client', 'index.js'),
  output: {
    path: path.resolve(__dirname, 'dist', 'static'),
    filename: 'bundle.js',
  },
  resolve: {
    modules: [path.resolve(__dirname, 'client', 'modules'), 'node_modules'],
  },
  plugins: [new HtmlWebpackPlugin({
    'publicPath': 'static',
  })],
};
