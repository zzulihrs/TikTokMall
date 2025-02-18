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
  },
  state: {
    chatHistory: []
  },
  mutations: {
    addMessage(state, message) {
      state.chatHistory.push(message)
    },
    updateLastResponse(state, response) {
      state.chatHistory[state.chatHistory.length - 1].response += ` ${response}`
    }
  },
  actions: {
    addMessage({ commit }, message) {
      commit('addMessage', message)
    },
    updateLastResponse({ commit }, response) {
      commit('updateLastResponse', response)
    }
  },
  getters: {
    chatHistory: state => state.chatHistory
  }
})

// 初始化用户状态
store.dispatch('auth/initialize')

export default store
