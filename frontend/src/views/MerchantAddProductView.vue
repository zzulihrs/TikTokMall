<template>
  <el-container style="height: 100vh;">
    <!-- 主要内容区域 -->
    <el-main>
      <Functions />
      <el-row :gutter="20">
        <el-col :span="20">
          <h2>添加商品</h2>
          <el-form ref="addProductFormRef" :model="product" label-width="120px">
            <el-form-item label="商品名称">
              <el-input v-model="product.name" placeholder="请输入商品名称" />
            </el-form-item>
            <el-form-item label="商品类别">
              <el-select v-model="product.category" placeholder="请选择商品类别">
                <el-option v-for="category in categories" :key="category.id" :label="category.name"
                  :value="category.id" />
                <!-- 可根据实际情况添加更多类别 -->
              </el-select>
            </el-form-item>
            <el-form-item label="商品价格">
              <el-input v-model.number="product.price" placeholder="请输入商品价格" />
            </el-form-item>
            <el-form-item label="商品库存">
              <el-input v-model.number="product.stock" placeholder="请输入商品库存" />
            </el-form-item>
            <el-form-item label="商品描述">
              <el-input v-model="product.description" placeholder="请输入商品描述" type="textarea" />
            </el-form-item>
            <el-form-item label="商品图片">
              <el-input v-model="product.picture" placeholder="请输入商品图片链接" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="addProduct">添加</el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </el-main>
  </el-container>

</template>

<script setup>
import Functions from '@/components/merchant/Functions.vue';
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import MerchantLeft from '../components/merchant/Left.vue';
import { ElMessage } from 'element-plus';
import axios from 'axios';

const router = useRouter();
const store = useStore();
const addProductFormRef = ref(null);

// Vuex 数据
const categories = computed(() => store.getters['category/categories']);

const product = ref({
  name: '',
  category: '',
  price: 0,
  stock: 0,
  description: '',
  picture: ''
});

const addProduct = async () => {
  // TODO：添加商品
  try {
    await axios.post('/api/merchant/products', product.value);
    ElMessage.success('商品添加成功');
    router.push('/merchant/products');
  } catch (error) {
    ElMessage.error('商品添加失败：' + error.message);
  }
};

const resetForm = () => {
  product.value = {
    name: '',
    category: '',
    price: 0,
    stock: 0,
    description: '',
    picture: ''
  };
};
</script>

<style scoped>
.el-main {
  padding: 20px;
  background-color: #ffffff;
}
</style>
