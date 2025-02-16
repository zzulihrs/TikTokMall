import { createStore } from 'vuex'
import auth from './modules/auth'
import cart from './modules/cart'
import search from './modules/search'
import merchant from './modules/merchant'

const store = createStore({
  modules: {
    auth,
    cart,
    search,
    merchant
  }
})

// 初始化用户状态
store.dispatch('auth/initialize')

export default store
