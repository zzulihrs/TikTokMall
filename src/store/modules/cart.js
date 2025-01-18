export default {
  namespaced: true,
  
  state: () => ({
    items: []
  }),

  getters: {
    items: (state) => state.items,
    totalAmount: (state) => {
      return state.items.reduce((total, item) => {
        return total + (item.price * item.quantity)
      }, 0)
    }
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
    REMOVE_ITEM(state, id) {
      state.items = state.items.filter(i => i.id !== id)
    },
    CLEAR_CART(state) {
      state.items = []
    }
  },

  actions: {
    addItem({ commit }, item) {
      commit('ADD_ITEM', item)
    },
    removeItem({ commit }, id) {
      commit('REMOVE_ITEM', id)
    },
    clearCart({ commit }) {
      commit('CLEAR_CART')
    }
  }
}
