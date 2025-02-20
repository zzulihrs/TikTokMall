export default {
    namespaced: true,
  
    state: () => ({
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
    },
  
    mutations: {
    },
  
    actions: {
    }
  }
  