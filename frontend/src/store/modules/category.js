export default {
    namespaced: true,

    state: () => ({
        category: 'All',
      categories: [
        {
          id: 0,
          name: 'All',
          description: 'All'
        },
        {
          id: 1,
          name: 'T-Shirt',
          description: 'T-Shirt'
        },
        {
          id: 2,
          name: 'Sticker',
          description: 'Sticker'
        }
      ]
    }),

    getters: {
        // 获取搜索结果
        categories: state => state.categories,
        category: state => state.category
    },

    mutations: {
        Update_Category(state, payload) {
            state.category = payload;
        }
    },

    actions: {

    }
  }
