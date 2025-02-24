<template>
  <el-container>
    <el-header class="header">
      <el-row justify="space-between" align="middle">
        <el-col :span="4">
          <div class="logo">电商商城</div>
        </el-col>

        <el-col :span="12">
          <el-row align="middle">
            <el-col :span="12">
              <el-menu
                mode="horizontal"
                :default-active="activeIndex"
                @select="handleSelect"
              >
                <el-menu-item index="1">首页</el-menu-item>
                <el-sub-menu index="category">
                  <template #title>商品类别</template>
                  <el-menu-item index="2-1">T-Shirt</el-menu-item>
                  <el-menu-item index="2-2">Sticker</el-menu-item>
                </el-sub-menu>
              </el-menu>
            </el-col>
            <el-col :span="12">
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

        <!-- 当前页，个人头像左边，加一个我是店家的超链接 -->
        <el-col :span="8">
          <div class="header-right">
            <el-menu mode="horizontal" :ellipsis="false">
            <el-menu-item index="5" v-show="isAuthenticated" @click="goToMerchant">我的商品</el-menu-item>
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

    <el-main>
      <router-view />
    </el-main>

    <el-footer>
      <div class="footer-content">
        © 字节跳动青训营后端项目.
        <el-link type="primary" @click="openMerchantDialog" style="margin-left: 10px">
          成为商家
        </el-link>
      </div>

    </el-footer>
    <ChatDialog/>

    <!-- 成为商家弹窗 -->
    <el-dialog
      v-model="merchantDialogVisible"
      title="成为商家"
      width="400px"
    >
      <el-form :model="merchantForm" label-width="80px">
        <el-form-item label="兑换码">
          <el-input v-model="merchantForm.code" placeholder="请输入商家兑换码" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="merchantDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleBecomeMerchant">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import UserDropdown from './UserDropdown.vue'
import ChatDialog from './ChatDialog.vue'
// Verify router availability
onMounted(() => {
  if (!router) {
    console.error('Router is not available')
  }
})
import { ShoppingCart, Search } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

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
const goToMerchant = () => {
  // 商家权限认证，无误后，去 merchant 页面
  const id = computed(() => store.getters['merchant/getMerchantId']).value
  console.log('merchantId:', id)
  if (id > 0) {
    router.push('merchant')
  } else {
    ElMessage.warning("你还不是店家")
  }
}

const merchantDialogVisible = ref(false)
const merchantForm = reactive({
  code: ''
})

// 打开成为商家弹窗
const openMerchantDialog = () => {
  merchantDialogVisible.value = true
}

// 处理成为商家
const handleBecomeMerchant = async () => {
  try {
    const response = await axios.post('/api/merchant/register', {
      code: merchantForm.code
    })

    ElMessage.success('注册成功，您已成为商家！')
    merchantDialogVisible.value = false

    // 刷新页面或更新状态
    window.location.reload()
  } catch (error) {
    ElMessage.error('注册失败：' + error.message)
  }
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
  padding: 20px 0;
}

.el-link {
  font-size: 14px;
}
</style>
