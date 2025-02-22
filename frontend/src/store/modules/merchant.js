import axios from 'axios';
import { ElMessage } from 'element-plus';

const state = {
  // 列表页
  searchQuery: { // 搜索条件
    name: '',
    category_id: '',
    max_stock: '',
    min_price: '',
    max_price: '',
    page: 1,
    page_size: 10,
  },
  products: [], // 商品列表
  currentPage: 1,
  pageSize: 10,
  // 详情页
  aproductDetail: {
    id: '',
    name: '',
    description: "",
    stock: '',
    price: '',
    img_url: '',
    catetory: [
      {
        id: '',
        name: '',
        description: '',
      }
    ],
    slider_img: []
  },
  productDetail: {
    category: [
      {
        description: "Sticker",
        id: 2,
        name: "Sticker"
      }
    ],
    description: "The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ",
    id: 2,
    img_url: "/static/image/mouse-pad.jpeg",
    name: "Mouse-Pad",
    price: 8.8,
    slider_img: [
      "/static/image/mouse-pad.jpeg"
    ],
    stock: 100000
  }
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

  // 详情页
  SET_PRODUCT_DETAIL(state, product) {
    state.productDetail = product;
  },
};

const actions = {
  async ProductList({ commit, state }) {
    const config = {
      headers: {
        'Authorization': 'Hello World'
      }
    }

    const data = Object.entries({
      name: state.searchQuery.name,
      category_id: state.searchQuery.category_id,
      max_stock: state.searchQuery.max_stock,
      min_price: state.searchQuery.min_price,
      max_price: state.searchQuery.max_price,
      page_no: state.currentPage,
      page_size: state.pageSize,
      merchant_id: 1,
    }).reduce((acc, [key, value]) => {
      // 过滤空字符串、null 和 undefined
      if (value !== '' && value !== null && value !== undefined && value != 0) {
        acc[key] = value;
      }
      return acc;
    }, {});

    try {
      console.log("headers:", config.headers, "\ndata:", data, "\n")
      const response = await axios.post('/api/merchant/product/list', data, config);
      console.log('merchant/product/list: ', response.data)
      if (+response.data.code !== 200) {
        ElMessage.error('获取商品列表失败：' + response.data.message);
        return;
      }
      commit('SET_PRODUCTS', response.data.data);
    } catch (error) {
      ElMessage.error('获取商品列表失败：' + error.message);
    }
  },

  async GetProductMerchant({ commit }, { product_id, merchant_id }) {
    const config = {
      headers: {
        'Authorization': 'Hello World'
      }
    }
    const data = Object.entries({
      pid: product_id,
      mid: merchant_id
    })
    try {
      // console.log("headers:", config.headers, "\ndata:", data, "\n")
      const response = await axios.post('/api/merchant/product/detail', data, config);
      console.log('merchant/product/detail: ', response.data)
      if (+response.data.code !== 200) {
        ElMessage.error('获取商品详情失败：' + response.data.message);
        return;
      }
      commit('SET_PRODUCTS', response.data.data);
    } catch (error) {
      ElMessage.error('获取商品详情失败：' + error.message);
    }
  },

  // 添加商品
  async MerchantAddProduct({ commit, product }) {
    const config = {
      headers: {
        'Authorization': 'Hello World'
      }
    }
    const data = Object.entries({
      mid: 100,
      name: product.name,
      price: product.price,
      stock: product.stock,
      description: product.description,
      img_url: product.img_url,
      slider_imgs: product.slider_imgs,
      categories: product.categories
    })
    try {
      const response = await axios.post('/api/merchant/product/add', data, config);
      console.log('merchant/product/add: ', response.data)
      if (+response.data.code !== 200) {
        ElMessage.error('添加商品失败：' + response.data.message);
        return;
      }
    } catch (error) {
      ElMessage.error('添加商品失败：' + error.message);
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
  getTotalProducts: state => state.totalProducts,
  // 商品详情
  getProductDetail: state => state.productDetail,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
};
