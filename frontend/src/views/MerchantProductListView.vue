<template>
  <el-container style="height: 100vh;">
    <!-- 头部 -->
    <!-- <el-header style="height: 60px;">
        <MerchantHeader />
      </el-header> -->
    <!-- 主要内容区域 -->
    <el-main>
      <!-- 功能区 -->
      <Functions />
      <!-- 搜索条件 -->
      <el-form :inline="true" :model="searchQuery">

        <el-form-item label="名称：">
          <el-input v-model="searchQuery.name" placeholder="名称"></el-input>
        </el-form-item>
        <el-form-item label="类别" style="width: 200px">
          <el-select v-model="searchQuery.category_id" style="width:100%">
            <el-option v-for="category in categories" :key="category.id" :label="category.name" :value="category.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="库存<=">
          <el-input v-model.number="searchQuery.max_stock" />
        </el-form-item>
        <el-form-item label="最低价">
          <el-input v-model="searchQuery.min_price" placeholder="最低价"></el-input>
        </el-form-item>
        <el-form-item label="最高价">
          <el-input v-model="searchQuery.max_price" placeholder="最高价"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="clearSearchQuery">清空</el-button>
          <el-button type="primary" @click="searchProducts">查询</el-button>
        </el-form-item>
      </el-form>

      <!-- 商品列表 -->
      <el-table :data="products">
        <el-table-column prop="id" label="商品id" />
        <el-table-column prop="name" label="商品名称" />
        <el-table-column label="图片">
          <template #default="{ row }">
            <el-image :src="row.img_url" style="width: 50px; height: 50px" fit="cover" />
          </template>
        </el-table-column>
        <!-- <el-table-column label="类别">
          <template #default="{ row }">
            <span v-for="(category, index) in row.category" :key="index">
              {{ category.name }}
              <span v-if="index!== row.category.length - 1">, </span>
            </span>
          </template>
        </el-table-column> -->
        <el-table-column prop="price" label="价格" />
        <el-table-column prop="description" label="描述" />
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-button type="text" @click="editProduct(row)">编辑</el-button>
            <el-button type="text" @click="goToProductDetail(row)">详情页</el-button>
            <el-button type="text" @click="deleteProduct(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage4"
        :page-sizes="[10, 20, 30, 40]" :page-size="10" layout="total, sizes, prev, pager, next, jumper" :total="400">
      </el-pagination>

    </el-main> <!-- 修改这里：移除多余的闭合标签 -->
  </el-container>
</template>

<script setup>
import Functions from '@/components/merchant/Functions.vue';
import { computed, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';

const store = useStore();
const router = useRouter();

// 路由跳转
const goToProductDetail = (product) => router.push(`/products/${product.id}`);


// Vuex 数据
const searchQuery = computed(() => store.getters['merchant/getSearchQuery']);
const categories = computed(() => store.getters['category/categories']);
const products = computed(() => store.getters['merchant/getProducts']);
const currentPage = computed(() => store.getters['merchant/getCurrentPage']);
const pageSize = computed(() => store.getters['merchant/getPageSize']);
const totalProducts = computed(() => store.getters['merchant/getTotalProducts']);

// 操作方法
const clearSearchQuery = () => store.dispatch('merchant/SET_SEARCH_QUERY', { // 搜索条件
  name: '',
  category_id: 0,
  max_stock: '',
  min_price: '',
  max_price: '',
  page: 1,
  page_size: 10,
});

const searchProducts = () => {
  store.dispatch('merchant/ProductList');
};

// 分页处理
const handleSizeChange = (newSize) => store.dispatch('merchant/handleSizeChange', newSize);
const handleCurrentChange = (newPage) => store.dispatch('merchant/handleCurrentChange', newPage);
</script>

<style scoped>
.el-main {
  padding: 20px;
  background-color: #ffffff;
}
</style>
