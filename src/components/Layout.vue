<template>
  <el-container>
    <el-header class="header">
      <el-row justify="space-between" align="middle">
        <el-col :span="4">
          <div class="logo">电商商城</div>
        </el-col>

        <el-col :span="12">
          <el-row align="middle">
            <el-col :span="6">
              <el-menu
                mode="horizontal"
                :default-active="activeIndex"
                @select="handleSelect"
              >
                <el-menu-item index="1">首页</el-menu-item>
                <el-submenu index="2">
                  <template #title>商品分类</template>
                  <el-menu-item index="2-1">电子产品</el-menu-item>
                  <el-menu-item index="2-2">服装服饰</el-menu-item>
                  <el-menu-item index="2-3">家居用品</el-menu-item>
                  <el-menu-item index="2-4">图书音像</el-menu-item>
                  <el-menu-item index="2-5">食品饮料</el-menu-item>
                </el-submenu>
              </el-menu>
            </el-col>
            <el-col :span="18">
              <el-input
                v-model="searchQuery"
                placeholder="搜索商品..."
                class="search-bar"
                @keyup.enter="handleSearch"
              >
                <template #append>
                  <el-button :icon="Search" @click="handleSearch" />
                </template>
              </el-input>
            </el-col>
          </el-row>
        </el-col>

        <el-col :span="8">
          <div class="header-right">
            <el-menu mode="horizontal" :ellipsis="false">
              <el-menu-item v-if="!isAuthenticated" index="3" @click="goToLogin">登录</el-menu-item>
              <el-submenu v-else index="4">
                <UserDropdown></UserDropdown>
              </el-submenu>
              <el-menu-item index="6" @click="goToCart">
                <el-badge :value="cartCount" :max="99">
                  <el-icon><ShoppingCart /></el-icon>
                </el-badge>
              </el-menu-item>
            </el-menu>
          </div>
        </el-col>
      </el-row>
    </el-header>

    <el-main>
      <router-view />
    </el-main>

    <el-footer>
      <div class="footer-content">
        © 字节跳动青训营后端项目.
      </div>
    </el-footer>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import UserDropdown from './UserDropdown.vue'

// Verify router availability
onMounted(() => {
  if (!router) {
    console.error('Router is not available')
  }
})
import { ShoppingCart, Search } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const router = useRouter()
if (!router) {
  console.error('Router is not available')
}
const searchQuery = ref('')

onMounted(() => { // 最开始执行的，先获取数据
  console.log('组件已挂载');
  store.dispatch('search/searchProducts', "")
});

const handleSearch = async () => {
  const query = searchQuery.value.trim()
  if (true) { // 空白也可以搜索
    try {
      await store.dispatch('search/searchProducts', query)
      router.push({ path: '/search', query: { q: query } })
    } catch (error) {
      ElMessage.error('搜索失败：' + error.message)
    }
  }
}
const store = useStore()


const activeIndex = ref('1')
const userAvatar = ref('')
const userName = ref('Guest')
const cartCount = computed(() => store.getters['cart/totalQuantity'])

const isAuthenticated = computed(() => store.getters['auth/isAuthenticated'])

const handleSelect = (index) => {
  switch (index) {
    case '1':
      router.push('/')
      break
    case '2-1':
      router.push('/category/electronics')
      break
    case '2-2':
      router.push('/category/clothing')
      break
    case '2-3':
      router.push('/category/home')
      break
    case '2-4':
      router.push('/category/books')
      break
    case '2-5':
      router.push('/category/food')
      break
    case '4-1':
      router.push('/orders')
      break
  }
}

const goToLogin = () => {
  router.push('/login')
}

const goToRegister = () => {
  router.push('/register')
}

const handleLogout = () => {
  store.dispatch('auth/logout')
  router.push('/login')
}

const goToCart = () => {
  router.push('/cart')
}
</script>

<style scoped>
.header {
  position: sticky;
  top: 0;
  z-index: 1000;
  background: white;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
}

.header-right {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.username {
  margin-left: 8px;
}

.el-menu--horizontal .el-submenu > .el-submenu__title {
  display: flex;
  align-items: center;
}

.el-menu--horizontal .el-menu-item {
  display: flex;
  align-items: center;
  height: 60px;
}

.el-menu--horizontal .el-menu-item .el-icon {
  margin-right: 8px;
}

.el-submenu__title {
  display: flex;
  align-items: center;
  padding: 0 20px;
  font-weight: bold;
  color: #333;
}

.el-submenu .el-menu-item {
  min-width: 120px;
  padding-left: 40px !important;
}

.el-submenu .el-menu-item:hover {
  background-color: #f5f7fa;
}

.el-submenu .el-menu-item {
  min-width: 120px;
}

.el-menu--horizontal .el-submenu .el-submenu__title {
  height: 60px;
  line-height: 60px;
}

.footer-content {
  text-align: center;
  padding: 1rem 0;
}
</style>
