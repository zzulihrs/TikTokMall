import { createStore } from 'vuex'
import auth from './modules/auth'
import cart from './modules/cart'
import search from './modules/search'

const store = createStore({
  modules: {
    auth,
    cart,
    search
  }
})

// 初始化用户状态
store.dispatch('auth/initialize')

export default store
