<template>
  <el-container class="home-container">
    <el-main>
      <el-row :gutter="20">
        <el-col
          v-for="product in products"
          :key="product.id"
          :xs="24" :sm="12" :md="8" :lg="6"
        >
          <product-card :product="product" />
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script setup>
import {computed, onMounted, ref} from 'vue'
import ProductCard from '@/components/ProductCard.vue'
import store from "@/store";
import {ElMessage} from "element-plus";


// 计算属性products返回store?.state?.search?.searchResults
const products = computed(() => store?.state?.search?.searchResults)

onMounted(() => {
  handleSearch()
})

const handleSearch = async () => {
  const query = ''
  if (true) { // 空白也可以搜索
    try {
      await store.dispatch('search/searchProducts', query)
      // router.push({ path: '/search', query: { q: query } })
    } catch (error) {
      // ElMessage.error('搜索失败：' + error.message)
    }
  }
}

</script>

<style scoped>
.home-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.el-main {
  padding: 0;
}
</style>
