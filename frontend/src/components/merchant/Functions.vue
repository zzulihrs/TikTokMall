<template>
    <div style="display: flex; align-items: center; margin-bottom: 20px;">
        <div style="margin-right: 20px;">
            <el-button type="primary" @click="goToMerchantProductList">商品列表</el-button>
        </div>
        <div>
            <el-button type="primary" @click="goToMerchantProduct">添加商品</el-button>
        </div>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
const router = useRouter();
// 路由跳转
const goToMerchantProductList = () => router.push('/merchant/product/list');
const goToMerchantProduct = () => router.push('/merchant/product/add');
</script>

<style scoped></style>