const path = require('path');

module.exports = {
  mode: "development",
  entry: path.resolve(__dirname, 'client', 'index.js'),
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
  },
  resolve: {
    modules: [path.resolve(__dirname, 'client', 'modules'), 'node_modules'],
  },
};