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
          <el-table-column prop="orderNumber" label="订单号" width="180" />
          <el-table-column prop="createTime" label="下单时间" width="180" />
          <el-table-column prop="totalAmount" label="总金额" width="120" />
          <el-table-column prop="status" label="状态" width="120">
            <template #default="{ row }">
              <el-tag :type="statusType(row.status)">
                {{ statusText(row.status) }}
              </el-tag>
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
  </el-container>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useStore } from 'vuex'

const store = useStore()
const loading = ref(true)
const orders = ref([])

const statusText = (status) => {
  const statusMap = {
    1: '待付款',
    2: '已付款',
    3: '已发货',
    4: '已完成',
    5: '已取消'
  }
  return statusMap[status] || '未知状态'
}

const statusType = (status) => {
  const typeMap = {
    1: 'warning',
    2: '',
    3: 'info',
    4: 'success',
    5: 'danger'
  }
  return typeMap[status] || ''
}

const viewDetail = (order) => {
  // TODO: Implement order detail view
  console.log('View order detail:', order)
}

onMounted(async () => {
  try {
    // TODO: Fetch orders from API
    orders.value = [
      {
        orderNumber: '202401180001',
        createTime: '2024-01-18 10:00:00',
        totalAmount: 299.0,
        status: 2
      },
      {
        orderNumber: '202401170001',
        createTime: '2024-01-17 15:30:00',
        totalAmount: 599.0,
        status: 3
      }
    ]
  } catch (error) {
    console.error('Failed to fetch orders:', error)
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
</style>
