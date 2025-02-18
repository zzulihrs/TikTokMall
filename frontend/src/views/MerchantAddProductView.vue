<template>
  <el-container style="height: 100vh;">
    <!-- 左侧导航栏 -->
    <el-aside width="200px">
      <MerchantLeft />
    </el-aside>
    <!-- 右侧主体内容 -->
    <el-container>
      <!-- 主要内容区域 -->
      <el-main>
        <el-row :gutter="20">
          <el-col :span="24">
            <h2>添加商品</h2>
            <el-form ref="addProductFormRef" :model="product" label-width="120px">
              <el-form-item label="商品名称">
                <el-input v-model="product.name" placeholder="请输入商品名称" />
              </el-form-item>
              <el-form-item label="商品类别">
                <el-select v-model="product.category" placeholder="请选择商品类别">
                  <el-option label="全部" value="" />
                  <el-option label="T-Shirt" value="T-Shirt" />
                  <el-option label="Sticker" value="Sticker" />
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
  </el-container>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import MerchantLeft from '../components/merchant/Left.vue';
import { ElMessage } from 'element-plus';
import axios from 'axios';

const router = useRouter();
const store = useStore();
const addProductFormRef = ref(null);

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
.el-header,
.el-footer {
  background-color: #f5f7fa;
  padding: 20px;
}

.el-aside {
  background-color: #2196F3;
  padding: 20px;
}

.el-main {
  padding: 20px;
  background-color: #ffffff;
}
</style>
