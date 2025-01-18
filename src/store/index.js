import { createStore } from 'vuex'
import auth from './modules/auth'
import cart from './modules/cart'
import search from './modules/search'

export default createStore({
  modules: {
    auth,
    cart,
    search
  }
})
