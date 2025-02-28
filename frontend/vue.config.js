const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    port: 8000,
    proxy: {
      '/api': {
        // 通过 VUE_ENV 来判断是否是线上环境
        target: process.env.VUE_ENV === 'online' 
          ? 'http://62.234.20.127:8080/' // 线上服务器地址
          : 'http://localhost:8080/',
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          '^/api': '' // 重写路径
        }
      }
    },
    client: {
      overlay: false // 自适应窗口大小时webpack会报错
    },
  }
})
