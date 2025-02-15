import { createStore } from 'vuex'
import auth from './modules/auth'
import cart from './modules/cart'
import search from './modules/search'

export default createStore({
  modules: {
    auth,
    cart,
    search
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
