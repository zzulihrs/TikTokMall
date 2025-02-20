<template>
  <el-container style="height: 100vh;">
    <el-main>
      <Functions />
      <el-row :gutter="20">
        <el-col :span="24">
          <h2 class="product-title">商品详情</h2>
          <!-- {{ product }} -->
          <!-- 商品信息卡片 -->
          <el-card shadow="hover" class="product-card">
            <!-- 走马灯图片展示 -->
            <el-carousel v-if="product.slider_img.length > 0" height="400px">
              <el-carousel-item v-for="(img, index) in product.slider_img" :key="index">
                <el-image :src="img" fit="cover" style="width: 100%; height: 100%" v-if="img" />
                <div v-else class="empty-slider">暂无图片</div>
              </el-carousel-item>
            </el-carousel>
            <el-row :gutter="20" class="product-info">
              <!-- 左侧商品信息 -->
              <el-col :span="14">
                <h1 class="product-name">{{ product.name }}</h1>
                <div class="price-section">
                  <span class="price-label">价格：</span>
                  <span class="price">¥{{ product.price }}</span>
                </div>
                <div class="stock-section">
                  <el-tag type="success" size="large">库存：{{ product.stock }}件</el-tag>
                </div>
                <div class="category-section">
                  <el-tag v-for="category in product.category" :key="category.id" class="category-tag">
                    {{ category.name }}
                  </el-tag>
                </div>
              </el-col>

              <!-- 右侧商品描述 -->
              <el-col :span="10">
                <div class="description-box">
                  <h3>商品描述</h3>
                  <p class="description-content">{{ product.description }}</p>
                </div>
              </el-col>
            </el-row>
          </el-card>
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
import { ElMessage } from 'element-plus';
import axios from 'axios';

const router = useRouter();
const store = useStore();

// Vuex 数据
const categories = computed(() => store.getters['category/categories']);
const product = computed(() => store.getters['merchant/getProductDetail']);

</script>

<style scoped>
.product-title {
  margin-bottom: 30px;
  color: #333;
}

.product-card {
  margin-top: 20px;
  padding: 20px;
}

.product-name {
  font-size: 28px;
  color: #1a1a1a;
  margin-bottom: 20px;
}

.price-section {
  margin: 25px 0;
}

.price-label {
  font-size: 18px;
  color: #666;
}

.price {
  font-size: 32px;
  color: #f56c6c;
  font-weight: bold;
}

.stock-section {
  margin: 20px 0;
}

.category-tag {
  margin-right: 10px;
  margin-bottom: 10px;
  font-size: 14px;
}

.description-box {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  min-height: 300px;
}

.description-content {
  line-height: 1.8;
  color: #666;
  font-size: 15px;
}

.empty-slider {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #999;
  background: #f5f7fa;
}
</style>