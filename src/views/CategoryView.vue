<template>
  <el-container>
    <el-main>
      <el-row :gutter="20">
        <el-col :span="24">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>{{ categoryName }}</el-breadcrumb-item>
          </el-breadcrumb>
        </el-col>
        
        <el-col :span="24">
          <el-row :gutter="20">
            <el-col 
              v-for="product in products" 
              :key="product.id"
              :xs="24" :sm="12" :md="8" :lg="6"
            >
              <product-card :product="product" />
            </el-col>
          </el-row>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import ProductCard from '@/components/ProductCard.vue'

const route = useRoute()
const products = ref([])

const categoryName = computed(() => {
  switch (route.params.type) {
    case 'electronics':
      return '电子产品'
    case 'clothing':
      return '服装服饰'
    case 'home':
      return '家居用品'
    case 'books':
      return '图书音像'
    default:
      return '商品分类'
  }
})

// Mock data - replace with API call
const mockProducts = [
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
]

onMounted(() => {
  // TODO: Fetch products by category from API
  products.value = mockProducts
})
</script>

<style scoped>
.el-breadcrumb {
  margin-bottom: 20px;
}

.el-row {
  margin-bottom: 20px;
}
</style>
