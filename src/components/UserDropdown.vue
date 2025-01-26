<template>
  <el-dropdown trigger="click" @command="handleCommand">
    <div class="user-avatar">
      <el-avatar :size="57" :src="user?.avatar"/>
    </div>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item command="orders">
          <el-icon><Tickets /></el-icon>
          <span>订单中心</span>
        </el-dropdown-item>
        <el-dropdown-item command="logout" divided>
          <el-icon><SwitchButton /></el-icon>
          <span>退出登录</span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup>
import { computed } from 'vue'
import { SwitchButton, Tickets } from '@element-plus/icons-vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

const store = useStore()
const router = useRouter()

const user = computed(() => store?.state?.auth?.user)

const handleCommand = (command) => {
  if (command === 'logout') {
    store.dispatch('auth/logout')
    ElMessage.success('已退出登录')
    router.push('/')
  } else if (command === 'orders') {
    router.push('/orders')
  }
}
</script>

<style scoped>
.user-avatar {
  cursor: pointer;
  padding: 0 12px;
}

.user-avatar:hover {
  background-color: rgba(0, 0, 0, 0.04);
}
</style>
