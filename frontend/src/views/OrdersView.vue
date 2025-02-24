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
          <el-table-column label="订单状态" width="120">
            <template #default="{ row }">
              <el-tag :type="getStatusType(row.status)">
                {{ getStatusText(row.status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="200">
            <template #default="{ row }">
              <el-button
                  type="primary"
                  size="small"
                  @click="viewDetail(row)"
              >
                查看详情
              </el-button>
              <el-button
                  v-if="row.status === 0"
                  type="primary"
                  size="small"
                  @click="handlePay(row)"
              >
                支付
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
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import axios from 'axios'

const router = useRouter()

const loading = ref(true)
const orders = ref([])
const dialogVisible = ref(false) // 控制弹窗显示
const currentOrder = ref(null) // 当前选中的订单

// 处理支付
const handlePay = async (order) => {
  try {
    // 调用支付接口
    const response = await axios.post('/api/alipay', {
      transaction_id: order?.orderNumber,
      total_amount: order?.totalAmount.toFixed(2),
    })
    // 把transaction_id, total_amount拼接到alipay上

    if (response?.data?.PayUrl) {
      // 获取支付表单
      // 打开PayUrl
      window.open(response.data?.PayUrl, '_blank')

    } else {
      ElMessage.error('支付失败：' + response.data.message)
    }
  } catch (error) {
    ElMessage.error('支付发起失败：' + error.message)
  }
}

// 处理支付结果
const handlePayResult = async (event) => {
  // 移除事件监听
  window.removeEventListener('message', handlePayResult)

  if (event.data.paid) {
    // 支付成功，跳转到成功页面
    router.push('/pay-success')
  } else {
    ElMessage.error('支付失败')
  }
}

// 支付成功页面组件
const PaySuccess = {
  template: `
    <div class="pay-success">
      <el-result
        icon="success"
        title="支付成功"
        sub-title="您的订单已支付成功"
      >
        <template #extra>
          <el-button type="primary" @click="$router.push('/orders')">
            返回订单列表
          </el-button>
        </template>
      </el-result>
    </div>
  `
}

// 在路由配置中添加支付成功页面
router.addRoute({
  path: '/pay-success',
  component: PaySuccess
})

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

// 获取订单状态对应的标签类型
const getStatusType = (status) => {
  const statusMap = {
    0: 'warning',   // 待支付
    1: 'success',   // 支付成功
    2: 'danger',    // 支付失败
    3: 'info'       // 取消订单
  }
  return statusMap[status] || 'info'
}

// 获取订单状态对应的文字
const getStatusText = (status) => {
  const statusMap = {
    0: '待支付',
    1: '支付成功',
    2: '支付失败',
    3: '已取消'
  }
  return statusMap[status] || '未知状态'
}

onMounted(async () => {
  try {
    const response = await axios.get('/api/order')
    orders.value = response.data.orders
      .map(order => ({
        orderNumber: order?.OrderId,
        createTime: order?.CreatedDate,
        totalAmount: order?.Cost,
        status: order?.OrderStatus,  // 添加状态字段
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
.el-button {
  margin-right: 5px;
}
.el-button:last-child {
  margin-right: 0;
}
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

.pay-success {
  padding: 40px;
  text-align: center;
}
</style>