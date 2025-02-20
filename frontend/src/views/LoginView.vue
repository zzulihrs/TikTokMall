<template>
  <el-container class="auth-container">
    <el-card class="auth-card">
      <h2 class="auth-title">登录</h2>
      <el-form
        :model="form"
        :rules="rules"
        ref="loginForm"
        @submit.prevent="handleLogin"
      >
        <el-form-item prop="email">
          <el-input
            v-model="form.email"
            placeholder="邮箱"
            prefix-icon="el-icon-user"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="密码"
            prefix-icon="el-icon-lock"
            show-password
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            native-type="submit"
            class="auth-button"
            :loading="loading"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
      <div class="auth-footer">
        没有账号？<el-link type="primary" @click="goToRegister">立即注册</el-link>
      </div>
    </el-card>
  </el-container>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'
import axios from "axios";

const store = useStore()

const router = useRouter()

const form = ref({
  email: '',
  password: ''
})

const rules = ref({
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ]
})

const loading = ref(false)


const handleLogin = async () => {

  try {
    const response = await axios.post('/api/auth/login', {
      email: form.value.email,
      password: form.value.password,
    });

    if(response?.status === 200)  {
      const userData = {
        email: response.data.email,
        password: form.value.password,
        avatar: response?.data?.avatar || 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png',
        token: response.data.token || 'test-token',
      };

      await store.dispatch('auth/login', userData);
      loading.value = true;

      ElMessage.success('登录成功');
      await router.push('/');
      await store.dispatch('cart/fetchCart')

    } else {
      loading.value = false;
      ElMessage.error(response?.data);
      console.error('Login failed:', response?.data);
    }

  } catch (error) {
    loading.value = false;
    ElMessage.error('登录失败');
    console.error('Login failed:', error);
  }
};


const goToRegister = () => {
  router.push('/register')
}
</script>

<style scoped>
.auth-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.auth-card {
  width: 400px;
  padding: 20px;
}

.auth-title {
  text-align: center;
  margin-bottom: 20px;
}

.auth-button {
  width: 100%;
}

.auth-footer {
  text-align: center;
  margin-top: 15px;
}
</style>
