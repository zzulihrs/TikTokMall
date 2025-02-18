<template>
  <el-container style="height: 100vh;">
    <!-- 左侧导航栏 -->
    <el-aside width="200px">
      <MerchantLeft />
    </el-aside>
    <!-- 右侧主体内容 -->
    <el-container>
      <!-- 头部 -->
      <!-- <el-header style="height: 60px;">
        <MerchantHeader />
      </el-header> -->
      <!-- 主要内容区域 -->
      <el-main>
        <!-- 搜索条件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <h2>店家商品管理</h2>
            <el-row :gutter="20">
              <el-col :span="3">
                <span>名称：</span><el-input v-model="searchQuery" placeholder="名称" />
              </el-col>
              <el-col :span="3">
                <span>类别：</span>
                <el-select v-model="categoryFilter" placeholder="类别">
                  <el-option label="全部" value="" />
                  <el-option label="T-Shirt" value="T-Shirt" />
                  <el-option label="Sticker" value="Sticker" />
                  <!-- 可根据实际情况添加更多类别 -->
                </el-select>
              </el-col>
              <el-col :span="3">
                <span>库存<= </span><el-input v-model.number="stockFilter" placeholder="库存<=" @input="validatePositiveNumber(stockFilter)" />
              </el-col>
              <el-col :span="3">
                <span>最低价：</span><el-input v-model.number="minPriceFilter" placeholder="最低价" @input="validateNumber(minPriceFilter)" />
              </el-col>
              <el-col :span="3">
                <span>最高价：</span><el-input v-model.number="maxPriceFilter" placeholder="最高价" @input="validateNumber(maxPriceFilter)" />
              </el-col>
              <el-col :span="3">
                <el-button @click="clearFilters">清空</el-button>
                <el-button type="primary" @click="searchProducts">查询</el-button>
              </el-col>
            </el-row>
          </el-col>
        </el-row>
        <!-- 商品列表 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-table :data="products" style="width: 100%">
              <el-table-column prop="name" label="商品名称" />
              <el-table-column label="图片">
                <template #default="{ row }">
                  <el-image :src="row.picture" style="width: 50px; height: 50px" fit="cover" />
                </template>
              </el-table-column>
              <el-table-column prop="category" label="类别" />
              <el-table-column prop="price" label="价格" />
              <el-table-column prop="description" label="描述" />
              <el-table-column label="操作">
                <template #default="{ row }">
                  <el-button type="text" @click="editProduct(row)">编辑</el-button>
                  <el-button type="text" @click="deleteProduct(row)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="currentPage"
              :page-sizes="[10, 20, 30]"
              :page-size="pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="totalProducts"
              style="margin-top: 20px; text-align: center;"
            >
            </el-pagination>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { Search } from '@element-plus/icons-vue';
import { useStore } from 'vuex';
import MerchantLeft from '../components/merchant/Left.vue';

const router = useRouter();
const store = useStore();

const searchQuery = computed({
  get: () => store.getters['merchant/getSearchQuery'],
  set: (value) => store.commit('merchant/SET_SEARCH_QUERY', value)
});

const categoryFilter = computed({
  get: () => store.getters['merchant/getCategoryFilter'],
  set: (value) => store.commit('merchant/SET_CATEGORY_FILTER', value)
});

const products = computed(() => store.getters['merchant/getProducts']);
const currentPage = computed(() => store.getters['merchant/getCurrentPage']);
const pageSize = computed(() => store.getters['merchant/getPageSize']);
const totalProducts = computed(() => store.getters['merchant/getTotalProducts']);

const stockFilter = ref('');
const minPriceFilter = ref('');
const maxPriceFilter = ref('');

const validatePositiveNumber = (value) => {
  if (value < 0) {
    stockFilter.value = '';
  }
};

const validateNumber = (value) => {
  if (isNaN(value)) {
    if (value === '') return;
    if (value.includes('.')) {
      value = parseFloat(value);
      if (isNaN(value)) {
        if (value === '') return;
        value = '';
      }
    } else {
      value = parseInt(value);
      if (isNaN(value)) {
        if (value === '') return;
        value = '';
      }
    }
  }
};

const clearFilters = () => {
  // TODO：条件初始化
  searchQuery.value = '';
  categoryFilter.value = '';
  stockFilter.value = '';
  minPriceFilter.value = '';
  maxPriceFilter.value = '';
  searchProducts();
};

const fetchProducts = () => store.dispatch('merchant/fetchProducts');
const searchProducts = () => {
  store.commit('merchant/SET_CURRENT_PAGE', 1);
  store.dispatch('merchant/searchProducts', {
    searchQuery: searchQuery.value,
    categoryFilter: categoryFilter.value,
    stockFilter: stockFilter.value,
    minPriceFilter: minPriceFilter.value,
    maxPriceFilter: maxPriceFilter.value
  });
};
const handleSizeChange = (newSize) => store.dispatch('merchant/handleSizeChange', newSize);
const handleCurrentChange = (newPage) => store.dispatch('merchant/handleCurrentChange', newPage);

const addProduct = () => {
  router.push('/merchant/products/add'); // 假设添加商品页面的路由
};

const editProduct = (product) => {
  router.push(`/merchant/products/edit/${product.id}`); // 假设编辑商品页面的路由
};

const deleteProduct = async (product) => {
  // TODO：删除商品逻辑。参数（商品ID，店家userID）
  try {
    await axios.delete(`/api/merchant/products/${product.id}`);
    ElMessage.success('删除商品成功');
    fetchProducts();
  } catch (error) {
    ElMessage.error('删除商品失败：' + error.message);
  }
};

fetchProducts();
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
