<template>
  <el-dropdown trigger="click" @command="handleCommand">
    <div class="user-avatar">
      <el-avatar :size="57" :src="user?.avatar"/>
    </div>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item command="edit">
          <el-icon><Edit /></el-icon>
          <span>个人信息</span>
        </el-dropdown-item>
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

  <!-- 编辑用户信息弹窗 -->
  <el-dialog
    v-model="dialogVisible"
    title="修改个人信息"
    width="400px"
  >
    <el-form :model="editForm" label-width="80px">
      <el-form-item label="头像">
        <el-upload
          class="avatar-uploader"
          action="/api/uploadImage"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="editForm.avatar" :src="editForm.avatar" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="用户名">
        <el-input v-model="editForm.username" />
      </el-form-item>
      <el-form-item label="邮箱">
        <el-input v-model="editForm.email" type="email" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSave">保存</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { computed, ref, reactive } from 'vue'
import { SwitchButton, Tickets, Edit, Plus } from '@element-plus/icons-vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import axios from "axios";
import Cookies from 'js-cookie';

const store = useStore()
const router = useRouter()
const dialogVisible = ref(false)
const editForm = reactive({
  avatar: '',
  username: '',
  email: ''
})

const user = computed(() => store?.state?.auth?.user)

const handleCommand = (command) => {
  if (command === 'logout') {
    axios.post('/api/auth/logout');
    store.dispatch('cart/fetchCart')
    Cookies.remove("cloudwego-shop")
    store.dispatch('auth/logout')
    ElMessage.success('已退出登录')
    router.push('/')
  } else if (command === 'orders') {
    router.push('/orders')
  } else if (command === 'edit') {
    openEditDialog()
  }
}

// 打开编辑弹窗
const openEditDialog = async () => {
  // 获取用户信息
  const response = await axios.get('/api/user/query')
  console.log(response?.data?.data);
  store.commit('auth/SET_USER', response?.data?.data)

  editForm.avatar = user.value?.avator || ''
  editForm.username = user.value?.username || ''
  editForm.email = user.value?.email || ''
  dialogVisible.value = true
}

// 头像上传前的验证
const beforeAvatarUpload = (file) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isJpgOrPng) {
    ElMessage.error('头像只能是 JPG 或 PNG 格式!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('头像大小不能超过 2MB!')
    return false
  }
  return true
}

// 头像上传成功的回调
const handleAvatarSuccess = (res) => {
  editForm.avatar = res.url
}

// 保存用户信息
const handleSave = async () => {
  try {
    await axios.post('/api/user/update', {
      avator: editForm?.avatar,
      username: editForm?.username,
    })

    // 更新 store 中的用户信息
    store.commit('auth/SET_USER', {
      ...user.value,
      avatar: editForm.avatar,
      username: editForm.username,
    })

    ElMessage.success('保存成功')
    dialogVisible.value = false
  } catch (error) {
    ElMessage.error('保存失败：' + error.message)
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

.avatar-uploader {
  text-align: center;
  margin-bottom: 20px;
}

.avatar-uploader .avatar {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
}

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 50%;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
  width: 100px;
  height: 100px;
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100px;
  height: 100px;
  text-align: center;
  line-height: 100px;
}

.dialog-footer {
  margin-top: 20px;
}

.el-form-item {
  margin-bottom: 20px;
}
</style>
