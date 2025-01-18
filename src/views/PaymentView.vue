<template>
  <el-container>
    <el-main>
      <el-card>
        <el-steps :active="activeStep" finish-status="success">
          <el-step title="确认订单" />
          <el-step title="支付" />
          <el-step title="完成" />
        </el-steps>

        <div v-if="activeStep === 0">
          <h3>订单详情</h3>
          <el-table :data="cartItems" stripe>
            <el-table-column prop="name" label="商品名称" />
            <el-table-column prop="price" label="单价" />
            <el-table-column prop="quantity" label="数量" />
            <el-table-column label="小计">
              <template #default="{row}">
                {{ (row.price * row.quantity).toFixed(2) }}
              </template>
            </el-table-column>
          </el-table>

          <el-divider />

          <div class="total-amount">
            总计: ¥{{ totalAmount.toFixed(2) }}
          </div>

          <el-button type="primary" @click="nextStep">去支付</el-button>
        </div>

        <div v-if="activeStep === 1">
          <h3>选择支付方式</h3>
          <el-radio-group v-model="paymentMethod">
            <el-radio label="alipay">支付宝</el-radio>
            <el-radio label="wechat">微信支付</el-radio>
            <el-radio label="bank">银行卡支付</el-radio>
          </el-radio-group>

          <el-button type="primary" @click="handlePayment">立即支付</el-button>
        </div>

        <div v-if="activeStep === 2">
          <el-result
            icon="success"
            title="支付成功"
            sub-title="您的订单已支付成功，我们将尽快为您发货"
          >
            <template #extra>
              <el-button type="primary" @click="goToHome">返回首页</el-button>
            </template>
          </el-result>
        </div>
      </el-card>
    </el-main>
  </el-container>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

const router = useRouter()
const store = useStore()

const activeStep = ref(0)
const paymentMethod = ref('alipay')

const cartItems = computed(() => store.getters['cart/items'])
const totalAmount = computed(() => store.getters['cart/totalAmount'])

const nextStep = () => {
  activeStep.value++
}

const handlePayment = () => {
  // 模拟支付处理
  setTimeout(() => {
    activeStep.value++
    store.dispatch('cart/clearCart')
  }, 1000)
}

const goToHome = () => {
  router.push('/')
}
</script>

<style scoped>
.total-amount {
  font-size: 1.2rem;
  font-weight: bold;
  text-align: right;
  margin: 20px 0;
}
</style>
