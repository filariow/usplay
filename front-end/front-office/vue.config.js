module.exports = {
  transpileDependencies: ["vuetify"],
  configureWebpack: {
    devtool: "source-map"
  },
  devServer: {
    public: "0.0.0.0",
    host: "0.0.0.0",
    hot: true,
    open: true,
    proxy: {
      "/bookmastergrpc.ActivityTypeSvc": {
        target: "http://localhost:9091/bookmastergrpc.ActivityTypeSvc",
        changeOrigin: true
      },
    }
  }
};
