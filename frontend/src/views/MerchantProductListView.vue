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
      <el-table :data="products || []">
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

      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="[10, 20, 30, 40]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="totalCount">
      </el-pagination>

      <!-- 编辑弹窗 -->
      <el-dialog
        v-model="editDialogVisible"
        title="编辑商品"
        width="1200px"
        :close-on-click-modal="false"
        destroy-on-close
        @closed="resetForm"
      >
        <el-form
          ref="editFormRef"
          :model="editForm"
          label-width="120px"
          :size="formSize"
        >
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="商品名称" prop="name">
                <el-input v-model="editForm.name" placeholder="请输入商品名称" clearable />
              </el-form-item>
            </el-col>

            <el-col :span="12">
              <el-form-item label="商品类别" prop="categories">
                <el-select
                  v-model="editForm.categories"
                  multiple
                  placeholder="请选择商品类别"
                  style="width: 100%"
                  clearable
                  collapse-tags
                  collapse-tags-tooltip
                >
                  <el-option
                    v-for="category in categories.slice(1)"
                    :key="category.id"
                    :label="category.name"
                    :value="category.id"
                  />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>

          // ... 其他表单项保持不变 ...

        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="closeDialog">取消</el-button>
            <el-button type="primary" @click="handleSubmit">保存修改</el-button>
          </div>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script setup>
import Functions from '@/components/merchant/Functions.vue';
import { computed, reactive, ref, nextTick } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { ElMessage } from 'element-plus'
import axios from "axios";

const store = useStore();
const router = useRouter();
const formSize = ref('default')

const editDialogVisible = ref(false)
const editFormRef = ref(null)
const editForm = reactive({
  id: '',
  name: '',
  categories: [],
  price: 0,
  stock: 0,
  description: '',
  img_url: ''
})

// 使用计算属性获取数据
const products = computed(() => store.state.merchant.products || [])
const categories = computed(() => store.state.category.categories || [])

// 路由跳转
const goToProductDetail = (product) => router.push(`/products?id=${product.id}`);


// Vuex 数据
const searchQuery = computed(() => store.getters['merchant/getSearchQuery']);
const currentPage = computed(() => store.getters['merchant/getCurrentPage']);
const pageSize = computed(() => store.getters['merchant/getPageSize']);
const totalCount = computed(() => store.getters['merchant/gettotalCount']);

// 操作方法
const clearSearchQuery = () => store.dispatch('merchant/SET_SEARCH_QUERY', { // 搜索条件
  name: '',
  category_id: 0,
  max_stock: '',
  min_price: '',
  max_price: '',
  page: 1,
  page_size: 20,
});

const searchProducts = () => {
  store.dispatch('merchant/ProductList');
};

// 分页处理
const handleSizeChange = (newSize) => {
  store.dispatch('merchant/handleSizeChange', newSize);
}
const handleCurrentChange = (newPage) => {
  store.dispatch('merchant/handleCurrentChange', newPage);
}


// 删除商品
const deleteProduct = async (row) => {
  await axios.post('/api/merchant/product/delete', {
   mid: store.state.merchant.id,
   pid: row.id,
  });
  ElMessage.success('删除成功');

  resetForm()
}

// 处理编辑按钮点击
const editProduct = async (row) => {
  // 先重置表单
  resetForm()

  // 等待下一个 DOM 更新周期
  await nextTick()

  // 复制数据到表单
  Object.assign(editForm, {
    id: row.id,
    name: row.name || '',
    categories: Array.isArray(row.categories) ? [...row.categories] : [],
    price: Number(row.price) || 0,
    stock: Number(row.stock) || 0,
    description: row.description || '',
    img_url: row.img_url || ''
  })

  // 打开弹窗
  editDialogVisible.value = true
}

// 重置表单
const resetForm = () => {
  Object.assign(editForm, {
    id: '',
    name: '',
    categories: [],
    price: 0,
    stock: 0,
    description: '',
    img_url: ''
  })
  if (editFormRef.value) {
    editFormRef.value.resetFields()
  }
}

// 关闭弹窗
const closeDialog = () => {
  editDialogVisible.value = false
  resetForm()
}

// 提交表单
const handleSubmit = async () => {
  if (!editFormRef.value) return

  try {
    await editFormRef.value.validate()
    await store.dispatch('merchant/updateProduct', { ...editForm })
    ElMessage.success('商品更新成功')
    closeDialog()
    // 刷新商品列表
    await store.dispatch('merchant/fetchProducts')
  } catch (error) {
    ElMessage.error('更新失败：' + (error.message || '请检查表单填写是否正确'))
  }
}
</script>

<style scoped>
.el-main {
  padding: 20px;
  background-color: #ffffff;
}

/* 添加过渡效果 */
.el-dialog {
  transition: all 0.3s ease;
}

.dialog-fade-enter-active,
.dialog-fade-leave-active {
  transition: opacity 0.3s ease;
}

.dialog-fade-enter-from,
.dialog-fade-leave-to {
  opacity: 0;
}

/* 优化表单样式 */
:deep(.el-form-item__label) {
  font-weight: 500;
}

:deep(.el-input-number) {
  width: 100%;
}

:deep(.el-select) {
  width: 100%;
}
</style>
