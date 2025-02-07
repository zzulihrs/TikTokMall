export default {
  namespaced: true,

  state: () => ({
    user: JSON.parse(localStorage.getItem('user') || 'null'),
    token: localStorage.getItem('token') || null
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    user: (state) => state.user
  },

  mutations: {
    SET_USER(state, user) {
      state.user = user
      localStorage.setItem('user', JSON.stringify(user))
    },
    SET_TOKEN(state, token) {
      state.token = token
      localStorage.setItem('token', token)
    },
    CLEAR_AUTH(state) {
      state.user = null
      state.token = null
      localStorage.removeItem('user')
      localStorage.removeItem('token')
    }
  },

  actions: {
    login({ commit }, userData) {
      commit('SET_USER', userData)
      commit('SET_TOKEN', userData.token)
    },
    logout({ commit }) {
      commit('CLEAR_AUTH')
    },
    initialize({ commit }) {
      const user = localStorage.getItem('user')
      const token = localStorage.getItem('token')
      if (user) {
        commit('SET_USER', JSON.parse(user))
      }
      if (token) {
        commit('SET_TOKEN', token)
      }
    }
  }
}
