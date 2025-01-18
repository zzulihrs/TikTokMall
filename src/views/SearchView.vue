<template>
  <el-container>
    <el-main>
      <el-row :gutter="20">
        <el-col :span="24">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>搜索结果</el-breadcrumb-item>
          </el-breadcrumb>
        </el-col>
        
        <el-col :span="24">
          <el-row v-if="loading" justify="center">
            <el-col :span="24" class="text-center">
              <el-icon class="is-loading" :size="32">
                <Loading />
              </el-icon>
            </el-col>
          </el-row>
          
          <template v-else>
            <el-row v-if="products.length > 0">
              <el-col 
                v-for="product in products" 
                :key="product.id"
                :xs="24" :sm="12" :md="8" :lg="6"
              >
                <product-card :product="product" />
              </el-col>
            </el-row>
            
            <el-empty v-else description="没有找到相关商品" />
          </template>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { Loading } from '@element-plus/icons-vue'
import ProductCard from '@/components/ProductCard.vue'

const route = useRoute()
const products = ref([])
const loading = ref(false)

// Mock search function - replace with API call
const searchProducts = async (query) => {
  loading.value = true
  // Simulate network delay
  await new Promise(resolve => setTimeout(resolve, 500))
  
  // Mock data
  return [
    {
      id: 1,
      name: '智能手机',
      price: 2999,
      image: 'https://via.placeholder.com/300',
      description: '最新款智能手机'
    },
    {
      id: 2,
      name: '笔记本电脑',
      price: 5999,
      image: 'https://via.placeholder.com/300',
      description: '高性能笔记本电脑'
    }
  ].filter(p => p.name.includes(query))
}

watch(
  () => route.query.q,
  async (newQuery) => {
    if (newQuery) {
      products.value = await searchProducts(newQuery)
      loading.value = false
    }
  },
  { immediate: true }
)
</script>

<style scoped>
.el-breadcrumb {
  margin-bottom: 20px;
}

.text-center {
  text-align: center;
}
</style>
