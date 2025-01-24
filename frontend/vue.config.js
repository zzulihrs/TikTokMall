const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8000/', // 你要请求的接口的前缀
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          '^/api': '' // 重写路径
        }
      }
    }
  }
})
