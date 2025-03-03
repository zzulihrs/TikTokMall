<template>
  <el-container>
    <el-main>
      <el-row>
        <el-col :span="24">
          <h2>{{ store?.state?.category?.category }} 类别商品</h2>
          <el-empty v-if="categoryItems?.length==0" description="暂无商品" />
          <el-row v-else>
            <el-col
                v-for="product in categoryItems"
                :key="product?.id"
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
import {ref, onMounted, watch} from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import axios from 'axios'
import { ElMessage } from 'element-plus'
import ProductCard from '@/components/ProductCard.vue'


const route = useRoute()
const store = useStore()
const category = ref('')
const categoryItems = ref([])

// Fetch category items
const fetchCategoryItems = async () => {
  try {

    // 获取到    /category/T-Shirt的T-Shirt

    // category.value = route?.params?.type // Get category from route params
    // console.log(route?.params?.type);
    const categoryName = store?.state?.category?.category
    let response;
    if (categoryName == 'All') {

    } else {
      response = await axios.get(`/api/category/${store?.state?.category?.category}`)
    }

    categoryItems.value = response?.data?.items || []
  } catch (err) {
    categoryItems.value = null
    ElMessage.error('获取商品列表失败：' + err.message)
  }
}

// vue3监听route?.params?.type
watch(() => store?.state?.category?.category, (newType, oldType) => {
  console.log('路由参数 type 变化了', newType, oldType);
  // 在这里执行相关操作，比如重新获取数据
  fetchCategoryItems()
});

onMounted(() => {
  fetchCategoryItems()
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
