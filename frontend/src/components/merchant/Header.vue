<template>
  <el-header class="header" style="background-color: #f5f7fa; padding: 20px;">
    <el-row justify="space-between" align="middle">
      <el-col :span="4">
        <div class="logo">电商商城 - 店家管理</div>
      </el-col>
      <el-col :span="12">
      </el-col>
      <el-col :span="8">
        <div class="header-right">
          <el-menu mode="horizontal" :ellipsis="false">
            <el-menu-item index="5" v-show="isAuthenticated" @click="goToMerchant">我是店家</el-menu-item>
            <el-menu-item v-if="!isAuthenticated" index="3" @click="goToLogin">登录</el-menu-item>
            <el-menu-item v-else index="4">
              <UserDropdown></UserDropdown>
            </el-menu-item>
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
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { Search, ShoppingCart } from '@element-plus/icons-vue'
import UserDropdown from '@/components/UserDropdown.vue'

const router = useRouter()
const store = useStore()
const searchQuery = ref('')
const activeIndex = ref('1')
const cartCount = computed(() => store.getters['cart/totalQuantity'])
const isAuthenticated = computed(() => store.getters['auth/isAuthenticated'])

const handleSearch = async () => {
  const query = searchQuery.value.trim()
  if (query) {
    try {
      await store.dispatch('search/searchProducts', query)
      router.push({ path: '/search', query: { q: query } })
    } catch (error) {
      ElMessage.error('搜索失败：' + error.message)
    }
  }
}

const handleSelect = (index) => {
  switch (index) {
    case '1':
      router.push('/')
      break
    case '2-1':
      router.push('/category/T-Shirt')
      break
    case '2-2':
      router.push('/category/Sticker')
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

const goToMerchant = () => {
  router.push('/merchant/product/list')
}

const goToLogin = () => {
  router.push('/login')
}

const goToCart = () => {
  router.push('/cart')
}
</script>

<style scoped>
.header {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.logo {
  font-size: 24px;
  font-weight: bold;
}

.search-bar {
  width: 100%;
}
</style>
