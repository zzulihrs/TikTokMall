<template>
  <el-container class="auth-container">
    <el-card class="auth-card">
      <h2 class="auth-title">注册</h2>
      <el-form
        :model="form"
        :rules="rules"
        ref="registerForm"
        @submit.prevent="handleRegister"
      >
<!--        <el-form-item prop="username">-->
<!--          <el-input-->
<!--            v-model="form.username"-->
<!--            placeholder="用户名"-->
<!--            prefix-icon="el-icon-user"-->
<!--          />-->
<!--        </el-form-item>-->
        <el-form-item prop="email">
          <el-input
            v-model="form.email"
            placeholder="邮箱"
            prefix-icon="el-icon-message"
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
        <el-form-item prop="confirmPassword">
          <el-input
            v-model="form.confirmPassword"
            type="password"
            placeholder="确认密码"
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
            注册
          </el-button>
        </el-form-item>
      </el-form>
      <div class="auth-footer">
        已有账号？<el-link type="primary" @click="goToLogin">立即登录</el-link>
      </div>
    </el-card>
  </el-container>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import axios from "axios";

const router = useRouter()

const form = ref({
  username: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const validatePassword = (rule, value, callback) => {
  if (value !== form.value.password) {
    callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}

const rules = ref({
  // username: [
  //   { required: true, message: '请输入用户名', trigger: 'blur' }
  // ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入密码', trigger: 'blur' },
    { validator: validatePassword, trigger: 'blur' }
  ]
})

const loading = ref(false)

const handleRegister = async () => {
  loading.value = true;

  try {
    const response = await axios.post('/api/auth/register', {
      email: form.value.email,
      password: form.value.password,
      confirmPassword: form.value.confirmPassword
    });

    if(response?.status == 200) {
      const userData = {
        email: response.data.email,
        password: form.value.password,
        avatar: response?.data?.avatar || 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png',
        token: response.data.token,
      };

      store.dispatch('auth/login', userData);
      loading.value = false;

      ElMessage.success('注册成功');
      // 跳转login页面
      router.push('/login');
    } else {
      loading.value = false;
      ElMessage.error('注册失败: ', response);
      console.error('Register failed:', response);
    }

  } catch (error) {
    loading.value = false;
    ElMessage.error('注册失败');
    console.error('Register failed:', error);
  }
};


const goToLogin = () => {
  router.push('/login')
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
