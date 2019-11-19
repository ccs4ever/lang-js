const path = require('path');

module.exports = {
  mode: 'production',
  entry: path.resolve(__dirname, 'src', 'index.jsx'),
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js'
  },
  resolve: {
    extensions: [".go", ".jsx", ".js", ".json"]
  },
  devtool: "source-map",
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        use: ["source-map-loader", "babel-loader"]
      },
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
