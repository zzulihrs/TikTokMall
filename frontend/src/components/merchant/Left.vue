<template>
  <el-aside width="200px" style="height: 100vh; overflow-y: auto;">
    <el-menu
      mode="vertical"
      :default-active="activeMenu"
      @select="handleMenuSelect"
      background-color="#2196F3" 
      text-color="#FFFFFF"
      active-text-color="#FFEB3B"
      style="height: 100%;"
    >
      <el-menu-item index="1" @click="handleMenuSelect('1')">商品列表</el-menu-item>
      <el-menu-item index="2" @click="handleMenuSelect('2')">添加商品</el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
// import { Search, ShoppingCart } from '@element-plus/icons-vue'
// import UserDropdown from '@/components/UserDropdown.vue'

const router = useRouter()
const store = useStore()
const activeMenu = ref('1')

const handleMenuSelect = (key) => {
  activeMenu.value = key
  if (key === '1') {
    // 商品列表页面逻辑
    // 这里可以触发父组件的方法来加载商品列表
    router.push('/merchant/product/list')
  } else if (key === '2') {
    // 其他页面逻辑，可根据需求添加
    router.push('/merchant/product/add')
  }
}
</script>

<style scoped>
.el-aside {
  background-color: #2196F3; /* 设置侧边栏背景色为蓝色 */
  padding: 20px;
}

.el-menu {
  border-right: none;
}

.el-menu-item {
  transition: background-color 0.3s; /* 添加过渡效果 */
}

.el-menu-item:hover {
  background-color: #1976D2; /* 鼠标悬停时的背景色 */
}
</style>
