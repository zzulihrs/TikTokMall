import axios from 'axios';
import { ElMessage } from 'element-plus';

const state = {
  searchQuery: '',
  categoryFilter: '',
  products: [],
  currentPage: 1,
  pageSize: 10,
  totalProducts: 0
};

const mutations = {
  SET_SEARCH_QUERY(state, query) {
    state.searchQuery = query;
  },
  SET_CATEGORY_FILTER(state, filter) {
    state.categoryFilter = filter;
  },
  SET_PRODUCTS(state, products) {
    state.products = products;
  },
  SET_CURRENT_PAGE(state, page) {
    state.currentPage = page;
  },
  SET_PAGE_SIZE(state, size) {
    state.pageSize = size;
  },
  SET_TOTAL_PRODUCTS(state, total) {
    state.totalProducts = total;
  }
};

const actions = {
  async fetchProducts({ commit, state }) {
    try {
      // TODO: 店家商品查询（条件：名称，类别，库存，价格）
      const response = await axios.post(`/api/merchant/products?page=${state.currentPage}&limit=${state.pageSize}&search=${state.searchQuery}&category=${state.categoryFilter}`);
      commit('SET_PRODUCTS', response.data.items);
      commit('SET_TOTAL_PRODUCTS', response.data.total);
    } catch (error) {
      ElMessage.error('获取商品列表失败：' + error.message);
    }
  },
  searchProducts({ commit, dispatch }) {
    commit('SET_CURRENT_PAGE', 1);
    dispatch('fetchProducts');
  },
  handleSizeChange({ commit, dispatch }, newSize) {
    commit('SET_PAGE_SIZE', newSize);
    dispatch('fetchProducts');
  },
  handleCurrentChange({ commit, dispatch }, newPage) {
    commit('SET_CURRENT_PAGE', newPage);
    dispatch('fetchProducts');
  }
};

const getters = {
  getSearchQuery: state => state.searchQuery,
  getCategoryFilter: state => state.categoryFilter,
  getProducts: state => state.products,
  getCurrentPage: state => state.currentPage,
  getPageSize: state => state.pageSize,
  getTotalProducts: state => state.totalProducts
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
};
