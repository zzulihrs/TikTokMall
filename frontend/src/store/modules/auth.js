export default {
  namespaced: true,

  state: () => ({
    user: null,
    token: null
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    user: (state) => state.user
  },

  mutations: {
    SET_USER(state, user) {
      state.user = user
    },
    SET_TOKEN(state, token) {
      state.token = token
    },
    CLEAR_AUTH(state) {
      state.user = null
      state.token = null
    }
  },

  actions: {
    login({ commit }, userData) {
      commit('SET_USER', userData)
      localStorage.setItem('user', JSON.stringify(userData))
      localStorage.setItem('token', userData.token)// 假设token存储在userData中
      // 修改state.token
      commit('SET_TOKEN', userData.token)
    },
    logout({ commit }) {
      commit('CLEAR_AUTH')
      localStorage.removeItem('user')
    },
    initialize({ commit }) {
      const user = localStorage.getItem('user')
      if (user) {
        commit('SET_USER', JSON.parse(user))
      }
    }
  }
}
