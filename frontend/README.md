# Mall Frontend 项目说明

## 项目目录结构

```
mall-frontend/
├── public/                # 静态资源文件（不会被webpack处理）
│   ├── favicon.ico        # 网站图标
│   └── index.html         # 项目入口HTML文件
├── src/
│   ├── assets/            # 静态资源（图片、字体等）
│   ├── components/        # 公共组件
│   │   ├── Layout.vue     # 页面布局组件
│   │   ├── ProductCard.vue # 商品卡片组件
│   │   └── UserDropdown.vue # 用户下拉菜单组件
│   ├── router/            # 路由配置
│   │   └── index.js       # 路由配置文件
│   ├── store/             # Vuex状态管理
│   │   ├── index.js       # Vuex主文件
│   │   └── modules/       # 模块化状态管理
│   │       ├── auth.js    # 认证相关状态
│   │       ├── cart.js    # 购物车相关状态
│   │       └── search.js  # 搜索相关状态
│   ├── views/             # 页面视图组件
│   │   ├── HomeView.vue   # 首页
│   │   ├── CartView.vue   # 购物车页面
│   │   ├── ProductView.vue # 商品详情页
│   │   └── ...            # 其他页面
│   ├── App.vue            # 根组件
│   └── main.js            # 项目入口文件
├── .env                   # 环境变量配置文件
├── package.json           # 项目依赖配置
└── README.md              # 项目说明文档
```

## 开发环境配置
           
node版本：14.17.5
1. 安装依赖
```bash
npm install
```

2. 启动开发服务器
```bash
npm run serve 
```

3. 构建生产环境
```bash
npm run build
```


## Vue3 核心概念对比

| 前端概念       | 后端对应概念       | 说明                             |
|----------------|-------------------|----------------------------------|
| Vuex Store     | 全局变量/函数     | 用于存储全局共享状态             |
| Computed       | 计算属性          | 实时获取最新的数据             |
| Watch          | 观察者模式        | 监听数据变化，类似事件监听       |
| Props          | 方法参数          | 父组件向子组件传递数据           |
| Emit           | 事件触发          | 子组件向父组件传递事件           |
| Router         | 路由控制器        | 处理页面跳转，类似后端的路由     |
| Composition API| 服务层           | 将逻辑抽离，类似后端的Service层  |
|onMounted       | 生命周期钩子      | 组件挂载后执行，类似后端的初始化 |



## 后端接口对接示例
                         
后端需要把html渲染，修改成返回json数据
比如app/frontend/handler/product/product_servics.go

```go 
//c.HTML(consts.StatusOK, "search", resp)
c.JSON(consts.StatusOK, resp)
```


### 情况1: 不需要与其他页面有数据交互，只需要调用接口，对接字段

参考：https://github.com/zzulihrs/TikTokMall/commit/ff6c8f8156bfa0b3833e2857cc9ac3a094f81307
```vue
Index: src/views/ProductView.vue
IDEA additional info:
Subsystem: com.intellij.openapi.diff.impl.patch.CharsetEP
<+>UTF-8
===================================================================
diff --git a/src/views/ProductView.vue b/src/views/ProductView.vue
--- a/src/views/ProductView.vue	(revision db6645178153eafa0b4c469ca6266005ecbd0394)
+++ b/src/views/ProductView.vue	(revision ff6c8f8156bfa0b3833e2857cc9ac3a094f81307)
@@ -4,18 +4,18 @@
       <el-row :gutter="20">
         <el-col :span="12">
           <el-image
-            :src="product.image"
+            :src="product?.picture"  
                与后端属性对接
             fit="contain"
             class="product-main-image"
           />
         </el-col>
@@ -42,16 +42,45 @@
+import axios from "axios";
 
 const Qty = ref(1)
+const loading = ref(true)
+
建立函数，获取商品信息
+const fetchProduct = async () => {
+  try {
+    console.log('fetchProduct');
+    axios.defaults.baseURL = '/api';
+
     // 发送请求获取商品信息
+    const response = await axios.get(`/product?id=${route.params.id}`)
+
+    axios.defaults.baseURL = '/';
+
     // 不清楚返回数据结构，可以输出控制台查看
+    console.log(response);
+
+    if (!response?.status) {
+      throw new Error('Failed to fetch product')
+    }
+    const data = response?.data?.item
+    product.value = {
+      ...data,
+      price: parseFloat(data?.price)
+    }
+  } catch (error) {
+    ElMessage.error('Failed to load product details')
+    console.error(error)
+  } finally {
+    loading.value = false
+  }
+}
 
 onMounted(() => { // 初始就执行
-  // Fetch product details from API
-  // using productId from route.params.id
+  fetchProduct()
 })
 </script>
```

### 情况2: 页面数据会更新，或者与其它页面有数据交互

参考搜索功能
这个时候需要把数据存放在store里面，然后通过store来修改和获取数据，这样就可以实现数据共享了。
通过计算属性实现数据更新
```vue
 // 计算属性products返回store?.state?.search?.searchResults
const products = computed(() => store?.state?.search?.searchResults)
```

```js
// 导入 axios 库，用于发送 HTTP 请求
import axios from 'axios';

// 定义 Vuex 模块的初始状态
const state = {
  searchResults: [], // 存储搜索结果的数组
  isLoading: false, // 表示是否正在加载数据的布尔值
  error: null // 存储错误信息的变量，初始值为 null
};

// 定义 mutations，用于直接修改 state
const mutations = {
  // 设置搜索结果
  SET_SEARCH_RESULTS(state, results) {
    state.searchResults = results; // 将传入的 results 赋值给 state 的 searchResults
  },
  // 设置加载状态
  SET_LOADING(state, isLoading) {
    state.isLoading = isLoading; // 将传入的 isLoading 赋值给 state 的 isLoading
  },
  // 设置错误信息
  SET_ERROR(state, error) {
    state.error = error; // 将传入的 error 赋值给 state 的 error
  }
};

// 定义 actions，用于处理异步操作和业务逻辑
const actions = {
  // 搜索产品的 action
  async searchProducts({ commit }, query) {
    commit('SET_LOADING', true); // 开始搜索时，设置加载状态为 true
    try {
      // 发送 GET 请求到指定的 URL，携带查询参数 q
      const response = await axios.get('http://localhost:8080/search', {
        params: {
          q: query // 查询参数
        }
      });
      // 使用可选链操作符安全地获取响应数据中的 items，如果不存在则返回空数组
      commit('SET_SEARCH_RESULTS', response?.data?.items || []);
      commit('SET_ERROR', null); // 清除错误信息
    } catch (error) {
      // 捕获错误，设置错误信息
      commit('SET_ERROR', error.response?.data?.message || error.message);
      commit('SET_SEARCH_RESULTS', []); // 清空搜索结果
    } finally {
      commit('SET_LOADING', false); // 搜索结束后，设置加载状态为 false
    }
  }
};

// 定义 getters，用于从 state 中派生出一些状态
const getters = {
  // 获取搜索结果
  searchResults: state => state.searchResults,
  // 获取加载状态
  isLoading: state => state.isLoading,
  // 获取错误信息
  error: state => state.error
};

// 导出 Vuex 模块
export default {
  namespaced: true, // 开启命名空间，避免命名冲突
  state,
  mutations,
  actions,
  getters
};
```

