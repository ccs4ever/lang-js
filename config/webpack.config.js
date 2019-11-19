const path = require('path');

module.exports = {
  entry: path.resolve(__dirname, '..', 'src', 'index.js'),
  output: {
    path: path.resolve(__dirname, '..', 'dist'),
    filename: 'bundle.js'
  },
  module: {
      rules: [
          {
              test: /\.go/,
              use: ['golang-wasm-async-loader']
          }
      ]
  },
  node: {
      fs: 'empty'
  }
};
