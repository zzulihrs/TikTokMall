import { createStore } from 'vuex'
import cart from './modules/cart'

export default createStore({
  modules: {
    cart
  }
})
