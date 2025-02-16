<template>
  <el-container>
    <el-main>
      <el-card>
        <template #header>
          <div class="card-header">
            <span>我的订单</span>
          </div>
        </template>

        <el-table
            :data="orders"
            style="width: 100%"
            v-loading="loading"
            empty-text="暂无订单"
        >
          <el-table-column label="订单号" width="180">
            <template #default="{ row }">
              {{ formatOrderNumber(row.orderNumber) }}
            </template>
          </el-table-column>
          <el-table-column prop="createTime" label="下单时间" width="180" />
          <el-table-column label="总金额" width="120">
            <template #default="{ row }">
              ¥{{ row.totalAmount.toFixed(2) }}
            </template>
          </el-table-column>
          <el-table-column label="操作" width="120">
            <template #default="{ row }">
              <el-button
                  type="text"
                  size="small"
                  @click="viewDetail(row)"
              >
                查看详情
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </el-main>

    <!-- 详情弹窗 -->
    <el-dialog
        v-model="dialogVisible"
        title="订单详情"
        width="600px"
    >
      <div v-if="currentOrder" class="order-detail">
        <div class="detail-item">
          <span class="label">订单号：</span>
          <span class="value">{{ currentOrder.orderNumber }}</span>
        </div>
        <div class="detail-item">
          <span class="label">下单时间：</span>
          <span class="value">{{ currentOrder.createTime }}</span>
        </div>
        <div class="detail-item">
          <span class="label">总金额：</span>
          <span class="value">¥{{ currentOrder.totalAmount.toFixed(2) }}</span>
        </div>
        <div class="detail-item">
          <span class="label">商品列表：</span>
          <el-table :data="currentOrder.items" style="width: 100%">
            <el-table-column prop="ProductName" label="商品名称" />
            <el-table-column label="图片" width="100">
              <template #default="{ row }">
                <el-image
                    :src="row.Picture"
                    style="width: 50px; height: 50px"
                    fit="cover"
                />
              </template>
            </el-table-column>
            <el-table-column prop="Qty" label="数量" width="80" />
            <el-table-column prop="Cost" label="金额" width="100">
              <template #default="{ row }">
                ¥{{ row.Cost.toFixed(2) }}
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </el-dialog>
  </el-container>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import {ElMessage} from "element-plus";

const loading = ref(true)
const orders = ref([])
const dialogVisible = ref(false) // 控制弹窗显示
const currentOrder = ref(null) // 当前选中的订单

// 格式化订单号
const formatOrderNumber = (orderNumber) => {
  const firstDashIndex = orderNumber.indexOf('-')
  if (firstDashIndex !== -1) {
    return `${orderNumber.substring(0, firstDashIndex)}...`
  }
  return orderNumber
}

// 查看详情
const viewDetail = (order) => {
  currentOrder.value = order
  dialogVisible.value = true
}

onMounted(async () => {
  try {
    const response = await axios.get('/api/order')
    orders.value = response.data.orders
      .map(order => ({
        orderNumber: order?.OrderId,
        createTime: order?.CreatedDate,
        totalAmount: order?.Cost,
        items: order?.Items // 商品列表
      }))
      // 按照下单时间降序排列
      .sort((a, b) => {
        const timeA = new Date(a.createTime).getTime()
        const timeB = new Date(b.createTime).getTime()
        return timeB - timeA // 降序排列
      })
  } catch (error) {
    console.error('获取订单失败:', error)
    // ElMessage.error('获取订单失败')
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.card-header {
  font-size: 18px;
  font-weight: bold;
}

.order-detail {
  padding: 10px;
}

.detail-item {
  margin-bottom: 20px;
}

.label {
  font-weight: bold;
  margin-right: 10px;
}

.value {
  color: #666;
}
</style>