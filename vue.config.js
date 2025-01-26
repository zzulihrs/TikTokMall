const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080/', // 你要请求的接口的前缀
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          '^/api': '' // 重写路径
        }
      }
    }
  }
})
