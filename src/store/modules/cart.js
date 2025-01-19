import axios from 'axios'

export default {
  namespaced: true,

  state: () => ({
    items: [],
    loading: false,
    error: null
  }),

  getters: {
    items: (state) => state.items,
  totalAmount: (state) => {
    return state.items.reduce((total, item) => {
      return total + (item.price * item.quantity)
    }, 0)
  },
  totalQuantity: (state) => {
    return state.items.reduce((total, item) => total + item.quantity, 0)
  },
    loading: (state) => state.loading,
    error: (state) => state.error
  },

  mutations: {
    SET_ITEMS(state, items) {
      state.items = items
    },
    ADD_ITEM(state, item) {
      const existing = state.items.find(i => i.id === item.id)
      if (existing) {
        existing.quantity += item.quantity
      } else {
        state.items.push(item)
      }
    },
    UPDATE_QUANTITY(state, { id, quantity }) {
      const item = state.items.find(i => i.id === id)
      if (item) {
        item.quantity = quantity
      }
    },
    REMOVE_ITEM(state, id) {
      state.items = state.items.filter(i => i.id !== id)
    },
    CLEAR_CART(state) {
      state.items = []
    },
    SET_LOADING(state, loading) {
      state.loading = loading
    },
    SET_ERROR(state, error) {
      state.error = error
    }
  },

  actions: {
    async fetchCart({ commit }) {
      try {
        commit('SET_LOADING', true)
        const { data } = await axios.get('http://localhost:8080/api/cart')
        commit('SET_ITEMS', data.items)
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async addItem({ commit }, item) {
      try {
        commit('SET_LOADING', true)
        commit('ADD_ITEM', item)
        await axios.post('http://localhost:8080/api/cart', item)
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async removeItem({ commit }, id) {
      try {
        commit('SET_LOADING', true)
        commit('REMOVE_ITEM', id)
        await axios.delete(`http://localhost:8080/api/cart/${id}`)
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async clearCart({ commit }) {
      try {
        commit('SET_LOADING', true)
        commit('CLEAR_CART')
        await axios.delete('http://localhost:8080/api/cart')
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async updateQuantity({ commit }, { id, quantity }) {
      try {
        commit('SET_LOADING', true)
        commit('UPDATE_QUANTITY', { id, quantity })
        await axios.put(`http://localhost:8080/api/cart/${id}`, { quantity })
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    }
  }
}
