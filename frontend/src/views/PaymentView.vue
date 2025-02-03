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
          <el-form :model="form" :rules="rules" ref="paymentForm">
            <el-row :gutter="20">
              <el-col :span="12">
                <h4>联系信息</h4>
                <el-form-item label="邮箱" prop="email">
                  <el-input v-model="form.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>

                <h4>配送地址</h4>
                <el-row :gutter="10">
                  <el-col :span="12">
                    <el-form-item label="姓氏" prop="firstName">
                      <el-input v-model="form.firstName" placeholder="姓氏"></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="名字" prop="lastName">
                      <el-input v-model="form.lastName" placeholder="名字"></el-input>
                    </el-form-item>
                  </el-col>
                </el-row>

                <el-form-item label="街道地址" prop="street">
                  <el-input v-model="form.street" placeholder="街道地址"></el-input>
                </el-form-item>

                <el-form-item label="邮政编码" prop="zipCode">
                  <el-input v-model="form.zipCode" placeholder="邮政编码"></el-input>
                </el-form-item>

                <el-row :gutter="10">
                  <el-col :span="12">
                    <el-form-item label="城市" prop="city">
                      <el-input v-model="form.city" placeholder="城市"></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="省份" prop="province">
                      <el-input v-model="form.province" placeholder="省份"></el-input>
                    </el-form-item>
                  </el-col>
                </el-row>

                <el-form-item label="国家" prop="country">
                  <el-input v-model="form.country" placeholder="国家"></el-input>
                </el-form-item>
              </el-col>

              <el-col :span="12">
                <h4>支付信息</h4>
                <el-form-item label="支付方式" prop="paymentMethod">
                  <el-radio-group v-model="form.paymentMethod">
                    <el-radio label="card">银行卡</el-radio>
                    <el-radio label="alipay">支付宝</el-radio>
                    <el-radio label="wechat">微信支付</el-radio>
                  </el-radio-group>
                </el-form-item>

                <div v-if="form.paymentMethod === 'card'">
                  <el-form-item label="卡号" prop="cardNumber">
                    <el-input v-model="form.cardNumber" placeholder="卡号"></el-input>
                  </el-form-item>

                  <el-row :gutter="10">
                    <el-col :span="8">
                      <el-form-item label="有效期(月)" prop="expMonth">
                        <el-input v-model="form.expMonth" placeholder="MM"></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8">
                      <el-form-item label="有效期(年)" prop="expYear">
                        <el-input v-model="form.expYear" placeholder="YYYY"></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8">
                      <el-form-item label="CVV" prop="cvv">
                        <el-input v-model="form.cvv" placeholder="CVV"></el-input>
                      </el-form-item>
                    </el-col>
                  </el-row>
                </div>

                <el-form-item>
                  <el-button type="primary" @click="handlePayment">立即支付</el-button>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
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
import { ref, computed, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'

const form = reactive({
  email: '',
  firstName: '',
  lastName: '',
  street: '',
  zipCode: '',
  city: '',
  province: '',
  country: '',
  paymentMethod: 'card',
  cardNumber: '',
  expMonth: '',
  expYear: '',
  cvv: ''
})

const rules = reactive({
  email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
  ],
  firstName: [
    { required: true, message: '请输入姓氏', trigger: 'blur' }
  ],
  lastName: [
    { required: true, message: '请输入名字', trigger: 'blur' }
  ],
  street: [
    { required: true, message: '请输入街道地址', trigger: 'blur' }
  ],
  zipCode: [
    { required: true, message: '请输入邮政编码', trigger: 'blur' }
  ],
  city: [
    { required: true, message: '请输入城市', trigger: 'blur' }
  ],
  province: [
    { required: true, message: '请输入省份', trigger: 'blur' }
  ],
  country: [
    { required: true, message: '请输入国家', trigger: 'blur' }
  ],
  paymentMethod: [
    { required: true, message: '请选择支付方式', trigger: 'change' }
  ],
  cardNumber: [
    { required: true, message: '请输入卡号', trigger: 'blur' },
    { 
      validator: (rule, value, callback) => {
        if (form.paymentMethod === 'card' && !value) {
          callback(new Error('请输入卡号'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  expMonth: [
    { required: true, message: '请输入有效期月份', trigger: 'blur' }
  ],
  expYear: [
    { required: true, message: '请输入有效期年份', trigger: 'blur' }
  ],
  cvv: [
    { required: true, message: '请输入CVV', trigger: 'blur' }
  ]
})

const router = useRouter()
const store = useStore()

onMounted(() => {
  if (cartItems.value.length === 0) {
    ElMessage.warning('购物车为空，请先添加商品')
    router.push('/')
  }
})

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
