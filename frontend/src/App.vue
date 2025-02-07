<template>
  <Layout />
</template>

<script setup>
import { onMounted } from 'vue'
import { useStore } from 'vuex'
import Layout from './components/Layout.vue'

const store = useStore()

// 初始化用户状态
onMounted(() => {
  const token = localStorage.getItem('token')
  const userInfo = JSON.parse(localStorage.getItem('userInfo') || '{}')
  
  if (token && userInfo) {
    store.commit('auth/SET_TOKEN', token)
    store.commit('auth/SET_USER_INFO', userInfo)
  }
})
</script>

<style>
/* Global styles can be added here */
</style>
