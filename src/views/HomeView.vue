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
import { ref } from 'vue'
import ProductCard from '@/components/ProductCard.vue'

const products = ref([
  {
    id: 1,
    name: 'Product 1',
    description: 'Description for product 1',
    price: 19.99,
    image: 'https://via.placeholder.com/400'
  },
  {
    id: 2,
    name: 'Product 2',
    description: 'Description for product 2',
    price: 29.99,
    image: 'https://via.placeholder.com/400'
  },
  {
    id: 3,
    name: 'Product 3',
    description: 'Description for product 3',
    price: 39.99,
    image: 'https://via.placeholder.com/400'
  },
  {
    id: 4,
    name: 'Product 4',
    description: 'Description for product 4',
    price: 49.99,
    image: 'https://via.placeholder.com/400'
  }
])
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
