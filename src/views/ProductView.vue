<template>
  <el-container class="product-container">
    <el-main>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-image
            :src="product?.picture"
            fit="contain"
            class="product-main-image"
          />
        </el-col>
        <el-col :span="12">
          <div class="product-details">
            <h1 class="product-title">{{ product?.name }}</h1>
            <p class="product-price">${{ product?.price?.toFixed(2) }}</p>
            <el-divider />
            <p class="product-description">{{ product?.description }}</p>

            <div class="product-actions">
              <el-input-number
                v-model="Qty"
                :min="1"
                :max="10"
                size="large"
              />
              <el-button
                type="primary"
                size="large"
                @click="addToCart"
              >
                Add to Cart
              </el-button>
            </div>
          </div>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import axios from "axios";
import {useStore} from "vuex";

const route = useRoute()
const store = useStore()

let product = ref({
  id: route.params.id,
  name: '',
  description: '',
  Price: 0,
  image: ''
})
let Qty = ref(1)
const loading = ref(true)
let Id = ref(0)

const fetchProduct = async () => {
  try {
    console.log('fetchProduct');
    axios.defaults.baseURL = '/api';

    const response = await axios.get(`/product?id=${route.params.id}`)

    axios.defaults.baseURL = '/';

    console.log(response);

    if (!response?.status) {
      throw new Error('Failed to fetch product')
    }
    const data = response?.data?.item
    product.value = {
      ...data,
      Price: parseFloat(data?.Price)
    }
    Id = data?.id

  } catch (error) {
    ElMessage.error('Failed to load product details')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const addToCart = (event) => {
  event.preventDefault()
  event.stopPropagation()
  store.dispatch('cart/addItem', {
    product_id: Id,
    product_num: Qty.value,
    Id: Id,
    Qty: Qty.value,
  })
  ElMessage.success('已加入购物车')
}

onMounted(() => {
  fetchProduct()
})
</script>

<style scoped>
.product-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.product-main-image {
  width: 100%;
  height: 500px;
  background: #f5f5f5;
  border-radius: 4px;
}

.product-details {
  padding: 20px;
}

.product-title {
  font-size: 2rem;
  margin-bottom: 10px;
}

.product-price {
  font-size: 1.5rem;
  color: #409eff;
  font-weight: bold;
  margin-bottom: 20px;
}

.product-description {
  color: #666;
  line-height: 1.6;
  margin-bottom: 30px;
}

.product-actions {
  display: flex;
  gap: 20px;
  margin-top: 30px;
}
</style>
