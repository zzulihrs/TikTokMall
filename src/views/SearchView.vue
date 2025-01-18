<template>
  <el-container>
    <el-main>
      <el-row :gutter="20">
        <el-col :span="24">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>搜索结果</el-breadcrumb-item>
          </el-breadcrumb>
        </el-col>

        <el-col :span="24">
          <el-row v-if="loading" justify="center">
            <el-col :span="24" class="text-center">
              <el-icon class="is-loading" :size="32">
                <Loading />
              </el-icon>
            </el-col>
          </el-row>

          <template v-else>
            <el-row v-if="products?.length > 0">
              <el-col
                v-for="product in products"
                :key="product.id"
                :xs="24" :sm="12" :md="8" :lg="6"
              >
                <product-card :product="product" />
              </el-col>
            </el-row>

            <el-empty v-else description="没有找到相关商品" />
          </template>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script setup>
import {computed, ref} from 'vue'
import { useRoute } from 'vue-router'
import { Loading } from '@element-plus/icons-vue'
import ProductCard from '@/components/ProductCard.vue'
import store from "@/store";

const route = useRoute()
const loading = ref(false)

// 计算属性products返回store?.state?.search?.searchResults
const products = computed(() => store?.state?.search?.searchResults)

</script>

<style scoped>
.el-breadcrumb {
  margin-bottom: 20px;
}

.text-center {
  text-align: center;
}
</style>
