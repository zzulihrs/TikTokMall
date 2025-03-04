import router from '@/router';
import axios from 'axios';
import { ElMessage } from 'element-plus';

const state = {
  // 店家 id
  id: localStorage.getItem('merchantId') || '',
  name: localStorage.getItem('merchantName') || '',
  // 列表页
  searchQuery: { // 搜索条件
    name: '',
    category_id: '',
    max_stock: '',
    min_price: '',
    max_price: '',
  },
  products: [], // 商品列表
  currentPage: 1,
  pageSize: 10,
  totalCount: 0, // 符合条件的商品总数
  // 详情页
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
  },
  categories: []
};

const mutations = {
  SET_MERCHANT_ID(state, id) {
    state.id = id;
    localStorage.setItem('merchantId', state.id)
  },
  SET_SEARCH_QUERY(state, query) {
    state.searchQuery = query;
  },
  SET_CATEGORY_FILTER(state, filter) {
    state.categoryFilter = filter;
  },
  SET_TOTAL_COUNT(state, count) {
    state.totalCount = count;
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
  handleCurrentChange({ commit, dispatch}, page) {
    commit('SET_CURRENT_PAGE', page);
    dispatch('ProductList');
  },
  handleSizeChange({ commit, dispatch }, size) {
    commit('SET_PAGE_SIZE', 0);
    commit('SET_PAGE_SIZE', size);
    dispatch('ProductList');
  },
  // 列表
  async ProductList({ commit, state }) {
    const data = Object.entries({
      name: state.searchQuery.name,
      category_id: state.searchQuery.category_id,
      max_stock: state.searchQuery.max_stock,
      min_price: state.searchQuery.min_price,
      max_price: state.searchQuery.max_price,
      page_no: state.currentPage,
      page_size: state.pageSize,
      merchant_id: state.id,
    }).reduce((acc, [key, value]) => {
      // 过滤空字符串、null 和 undefined
      if (value !== '' && value !== null && value !== undefined && value != 0) {
        acc[key] = value;
      }
      return acc;
    }, {});

    try {
      console.log("merchant/product/list, data:", data, "\n")
      const response = await axios.post('/api/merchant/product/list', data);
      console.log('merchant/product/list: ', response.data)
      if (+response.data.code !== 200) {
        ElMessage.error('获取商品列表失败：' + response.data.message);
        return;
      }
      commit('SET_PRODUCTS', response.data.data.items);
      commit('SET_TOTAL_COUNT', response.data.data.total);
    } catch (error) {
      ElMessage.error('获取商品列表失败：' + error.message);
    }
  },
  // 商品商品详情
  async GetProductMerchant({ commit }, { product_id }) {
    const data = Object.entries({
      pid: product_id,
      mid: state.id
    })
    try {
      console.log("merchant/product/detail, data:", data, "\n")
      const response = await axios.post('/api/merchant/product/detail', data);
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
  async MerchantAddProduct({ state }, { product }) {
    const data = Object.entries({
      mid: state.id,
      name: product.name,
      price: product.price,
      stock: product.stock,
      description: product.description,
      img_url: product.img_url,
      slider_imgs: product.slider_imgs,
      categories: product.categories
    })
    try {
      console.log('merchant/product/add, data: ', data , "\n")
      const response = await axios.post('/api/merchant/product/add', data);
      console.log('merchant/product/add: ', response.data)
      if (+response.data.code !== 200) {
        ElMessage.error('添加商品失败：' + response.data.message);
        return;
      }
    } catch (error) {
      ElMessage.error('添加商品失败：' + error.message);
    }
  },

  // 权限认证
  async MerchantAuth({commit, state}) {

    // if (state.id > 0) {
    //   router.push('/merchant');
    //   return;
    // }
    try {
      const response = await axios.get('/api/merchant/auth');
      console.log('merchant/auth: ', response)
      if (+response.data.code !== 200) {
        // ElMessage.error('店家权限认证' + response?.data?.message);
        return;
      }
      commit('SET_MERCHANT_ID', response?.data?.merchant_info?.id)
    } catch (error) {
      commit('SET_MERCHANT_ID', 0)
      // ElMessage.error('店家权限认证' + error.message);
    }
  },
  async updateProduct({ state }, product) {
    try {
      const response = await axios.put(`/api/merchant/product/${product.id}`, {
        mid: state.id,
        name: product.name,
        price: product.price,
        stock: product.stock,
        description: product.description,
        img_url: product.img_url,
        categories: product.categories
      })

      if (response.data.code !== 200) {
        throw new Error(response.data.message)
      }

      return response.data
    } catch (error) {
      throw new Error(error.message)
    }
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
  // 店家id
  getMerchantId: state => state.id,
  gettotalCount: state => state.totalCount
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
};
