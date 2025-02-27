import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
      path: "/",
      redirect: "/home"
  },
  {
    path: '/home',
    name: 'home',
    component: () => import('../views/HomeView.vue')
  },
  {
    path: '/products/:id',
    name: 'product',
    component: () => import('../views/ProductView.vue'),
    props: true
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/LoginView.vue'),
    meta: { requiresGuest: true }
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/RegisterView.vue'),
    meta: { requiresGuest: true }
  },
  {
    path: '/category',
    name: 'category',
    component: () => import('../views/CategoryView.vue'),
    props: true
  },
  {
    path: '/search',
    name: 'search',
    component: () => import('../views/SearchView.vue'),
    props: route => ({ query: route.query.q })
  },
  {
    path: '/cart',
    name: 'cart',
    component: () => import('../views/CartView.vue')
  },
  {
    path: '/payment',
    name: 'payment',
    component: () => import('../views/PaymentView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/orders',
    name: 'orders',
    component: () => import('../views/OrdersView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/merchant',
    name: 'merchant',
    redirect: '/merchant/product/list',
    // component: () => import('../views/MerchantProductListView.vue'),
    // TODO: reuiresMerchant 店家权限校验
    meta: { requiresMerchant: true }, // 假设需要店家权限
    // 在这个路由中，添加子路由
    children: [
      {
        path: 'product/list',
        name:'merchantProductsList',
        component: () => import('../views/MerchantProductListView.vue')
      },
      {
        path: 'product/add',
        name:'merchantAddProduct',
        component: () => import('../views/MerchantAddProductView.vue')
      },
      {
        path: 'product/detail',
        name:'merchantProductDetail',
        component: () => import('../views/MerchantProductDetail.vue')
      },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

export default router
