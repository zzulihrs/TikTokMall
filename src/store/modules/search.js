import axios from 'axios'

const state = {
  searchResults: [],
  isLoading: false,
  error: null
}

const mutations = {
  SET_SEARCH_RESULTS(state, results) {
    state.searchResults = results
  },
  SET_LOADING(state, isLoading) {
    state.isLoading = isLoading
  },
  SET_ERROR(state, error) {
    state.error = error
  }
}

const actions = {
  async searchProducts({ commit }, query) {
    commit('SET_LOADING', true)
    try {
      const response = await axios.get('http://localhost:8080/search', {
        params: {
          q: query
        }
      })
      commit('SET_SEARCH_RESULTS', response?.data?.items || [])
      commit('SET_ERROR', null)
    } catch (error) {
      commit('SET_ERROR', error.response?.data?.message || error.message)
      commit('SET_SEARCH_RESULTS', [])
    } finally {
      commit('SET_LOADING', false)
    }
  }
}

const getters = {
  searchResults: state => state.searchResults,
  isLoading: state => state.isLoading,
  error: state => state.error
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
