<template>
  <router-link :to="`/products/${product.id}`" class="product-link">
    <el-card class="product-card">
      <el-image
        :src="product?.picture"
        fit="cover"
        class="product-image"
      />
      <div class="product-info">
        <h3 class="product-name">{{ product.name }}</h3>
        <p class="product-description">{{ product.description }}</p>
        <div class="product-footer">
          <span class="product-price">${{ product.price.toFixed(2) }}</span>
          <el-button
            type="primary"
            size="small"
            @click.prevent="addToCart"
          >
            加入购物车
          </el-button>
        </div>
      </div>
    </el-card>
  </router-link>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'

const router = useRouter()
const store = useStore()

const props = defineProps({
  product: {
    type: Object,
    required: true
  }
})

const addToCart = (event) => {
  event.preventDefault()
  event.stopPropagation()
  store.dispatch('cart/addItem', {
    ...props.product,
    quantity: 1
  })
  ElMessage.success('已加入购物车')
}
</script>

<style scoped>
.product-link {
  text-decoration: none;
  display: block;
}

.product-card {
  margin-bottom: 20px;
  transition: transform 0.2s;
}

.product-card:hover {
  transform: translateY(-5px);
}

.product-image {
  width: 100%;
  height: 200px;
}

.product-info {
  padding: 15px;
}

.product-name {
  margin: 0 0 10px;
  font-size: 1.2rem;
  color: #333;
}

.product-description {
  color: #666;
  margin: 0 0 15px;
}

.product-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.product-price {
  font-size: 1.1rem;
  font-weight: bold;
  color: #409eff;
}
</style>
