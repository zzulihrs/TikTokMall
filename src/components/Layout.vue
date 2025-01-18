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
                <template #title>
                  <el-avatar :size="32" :src="userAvatar" />
                </template>
                <el-menu-item index="4-1">我的订单</el-menu-item>
                <el-menu-item index="4-2" @click="handleLogout">退出登录</el-menu-item>
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
        © 2024 电商商城. 版权所有.
      </div>
    </el-footer>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

// Verify router availability
onMounted(() => {
  if (!router) {
    console.error('Router is not available')
  }
})
import { ShoppingCart, Search } from '@element-plus/icons-vue'

const router = useRouter()
if (!router) {
  console.error('Router is not available')
}
const searchQuery = ref('')

const handleSearch = () => {
  if (searchQuery.value.trim()) {
    router.push({ path: '/search', query: { q: searchQuery.value } })
  }
}
const store = useStore()

const activeIndex = ref('1')
const userAvatar = ref('')
const userName = ref('Guest')
const cartCount = ref(0)

const isAuthenticated = computed(() => store.getters.isAuthenticated)

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
  }
}

const goToLogin = () => {
  router.push('/login')
}

const goToRegister = () => {
  router.push('/register')
}

const handleLogout = () => {
  store.dispatch('logout')
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

.footer-content {
  text-align: center;
  padding: 1rem 0;
}
</style>
