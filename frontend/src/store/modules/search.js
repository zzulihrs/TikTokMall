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

// 定义 actions，用于处理异步操作和业务逻辑, 也就是函数
const actions = {
  // 搜索产品的 action
  async searchProducts({ commit }, query) {
    commit('SET_LOADING', true); // 开始搜索时，设置加载状态为 true
    try {
      // 发送 GET 请求到指定的 URL，携带查询参数 q
      const response = await axios.get('api/search', {
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