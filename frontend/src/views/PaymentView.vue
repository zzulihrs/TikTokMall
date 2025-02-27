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
            <el-table-column prop="Price" label="单价" />
            <el-table-column prop="Qty" label="数量" />
            <el-table-column label="小计">
              <template #default="{row}">
                {{ (row.Price * row.Qty).toFixed(2) }}
              </template>
            </el-table-column>
          </el-table>

          <el-divider />

          <div class="total-amount">
            总计: ¥{{ totalAmount?.toFixed(2) }}
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
<!--                    <el-radio label="wechat">微信支付</el-radio>-->
                  </el-radio-group>
                </el-form-item>

                <div v-if="form.paymentMethod === 'card'">
                  <el-form-item label="卡号" prop="cardNumber">
                    <el-input v-model="form.cardNumber" placeholder="卡号"></el-input>
                  </el-form-item>

                  <el-row :gutter="10">
                    <el-col :span="8">
                      <el-form-item label="有效期(月)" prop="expMonth">
                        <el-input
                          v-model.number="form.expMonth"
                          type="number"
                          placeholder="MM"
                        ></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8">
                      <el-form-item label="有效期(年)" prop="expYear">
                        <el-input
                          v-model.number="form.expYear"
                          type="number"
                          placeholder="YYYY"
                        ></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8">
                      <el-form-item label="CVV" prop="cvv">
                        <el-input
                          v-model.number="form.cvv"
                          type="number"
                          placeholder="CVV"
                        ></el-input>
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
import axios from 'axios'

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
  expMonth: null,
  expYear: null,
  cvv: null
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
    { required: true, message: '请输入有效期月份', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (!value || value < 1 || value > 12) {
          callback(new Error('月份必须在1-12之间'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  expYear: [
    { required: true, message: '请输入有效期年份', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        const currentYear = new Date().getFullYear()
        if (!value || value <= currentYear) {
          callback(new Error('年份不能小于当前年份'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  cvv: [
    { required: true, message: '请输入CVV', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (!value || String(value).length < 3 || String(value).length > 4) {
          callback(new Error('CVV必须是3-4位数字'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
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

const paymentForm = ref(null)

const handlePayment = async () => {
  // 表单验证
  const valid = await paymentForm.value.validate()
  if (!valid) {
    return
  }

  try {
    // 调用支付接口，确保转换为数字类型
    const response = await axios.post('/api/checkout/waiting', {
      firstname: form?.firstName,
      lastname: form?.lastName,
      email: form?.email,
      country: form?.country,
      zipcode: form?.zipCode,
      city: form?.city,
      province: form?.province,
      street: form?.street,
      card_num: form?.cardNumber,
      payment: form?.paymentMethod,
      expiration_year: parseInt(form?.expYear),
      expiration_month: parseInt(form?.expMonth),
      cvv: parseInt(form?.cvv)
    })
    if (paymentMethod.value === 'alipay') {
      // 跳转到订单页面
      await router.push('/orders')
      return
    }

    // 如果支付提交成功，开始轮询结果
    if (response.data.redirect) {
      const result = await pollCheckoutResult()
    } else {
      ElMessage.error('支付失败，请仔细检查您的信用卡信息')
    }
  } catch (err) {
    ElMessage.error('支付失败：' + err.message)
  }
}

// 轮询支付结果
const pollCheckoutResult = async () => {
  const maxAttempts = 10
  const interval = 2000 // 2秒间隔

  for (let i = 0; i < maxAttempts; i++) {
    try {
      const response = await axios.get('/api/checkout/result')

      if (response?.data?.user_id) {
        activeStep.value = 2
        store.dispatch('cart/clearCart')
        ElMessage.success('支付成功！')
        return response.data
      }
      await new Promise(resolve => setTimeout(resolve, interval))
    } catch (err) {
      console.error('轮询出错:', err)
    }
  }
  throw new Error('支付超时，请检查订单状态')
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
