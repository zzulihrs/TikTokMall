const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    port: 8000,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8080/', // 你要请求的接口的前缀
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          '^/api': '' // 重写路径
        }
      }
    }
  }
})
