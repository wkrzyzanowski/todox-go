module.exports = {
  outputDir: '../webapp/dist/public',
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        secure: false,
        changeOrigin: true,
        pathRewrite: { '^/api/': '' },
        logLevel: 'debug',
      },
    },
  },
};
