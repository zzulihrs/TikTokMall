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
        return total + (item.Price * item.Qty)
      }, 0)
    },
    totalQuantity: (state) => {
      return state?.items?.reduce((total, item) => total + item?.Qty, 0)
    },
    loading: (state) => state.loading,
    error: (state) => state.error
  },

  mutations: {
    SET_ITEMS(state, items) {
      state.items = items
    },
    ADD_ITEM(state, item) {
      const existing = state.items.find(i => i.Id === item.id)
      if (existing) {
        existing.Qty += item.Qty
      } else {
        state.items.push(item)
      }
    },
    UPDATE_QUANTITY(state, { id, Qty }) {
      const item = state.items.find(i => i.Id === id)
      if (item) {
        item.Qty = Qty
      }
    },
    REMOVE_ITEM(state, id) {
      state.items = state.items.filter(i => i.Id !== id)
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
        const { data } = await axios.get('/api/cart')
        commit('SET_ITEMS', data?.items || [])
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ITEMS', [])
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async addItem({ commit }, item) {
      try {
        commit('SET_LOADING', true)
        await axios.post('/api/cart', item)
        // commit('ADD_ITEM', item)
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async removeItem({ commit }, {Id}) {
      try {
        commit('SET_LOADING', true)
        commit('REMOVE_ITEM', Id)
        // await axios.delete(`http://localhost:8080/api/cart/${id}`)
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
        await axios.get(`/api/emptycart`)
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async updateQuantity({ commit }, { Id, Qty }) {
      try {
        commit('SET_LOADING', true)
        commit('UPDATE_QUANTITY', { Id, Qty })
        await axios.post(`/api/changeqty`, {
          product_id: Id,
          product_num: Qty,
        })
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    }
  }
}
