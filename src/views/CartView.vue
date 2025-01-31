<template>
  <el-container>
    <el-main>
      <el-row>
        <el-col :span="24">
          <h2>购物车</h2>
          <el-empty v-if="cartItems?.length === 0" description="购物车为空" />
          <el-table v-else :data="cartItems" style="width: 100%">
            <el-table-column prop="Id" label="id" />
            <el-table-column prop="Name" label="商品名称" />
            <el-table-column prop="Price" label="价格" width="120" />
            <el-table-column prop="Qty" label="数量" width="120">
              <template #default="{ row }">
                <el-input-number
                  v-model="row.Qty"
                  :min="1"
                  :max="99"
                  size="small"
                  @change="updateQuantity(row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="操作" width="120">
              <template #default="{ row }">
                <el-button
                  type="danger"
                  size="small"
                  @click="removeItem(row)"
                >
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>

          <div class="cart-actions">
            <el-button type="danger" @click="clearCart">清空购物车</el-button>
            <el-button
              type="primary"
              :disabled="cartItems?.length === 0"
              @click="goToPayment"
            >
              去结算
            </el-button>
          </div>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'

const router = useRouter()
const store = useStore()

const cartItems = computed(() => store.getters['cart/items'])
const totalAmount = computed(() => store.getters['cart/totalAmount'])
const loading = computed(() => store.getters['cart/loading'])
const error = computed(() => store.getters['cart/error'])

onMounted(async () => {
  try {
    await store.dispatch('cart/fetchCart')
  } catch (err) {
    ElMessage.error('获取购物车数据失败：' + err.message)
  }
})

const removeItem = async (item) => {
  try {
    await store.dispatch('cart/removeItem', item.id)
    ElMessage.success('删除成功')
  } catch (err) {
    ElMessage.error('删除失败：' + err.message)
  }
}

const clearCart = async () => {
  try {
    await store.dispatch('cart/clearCart')
    ElMessage.success('购物车已清空')
  } catch (err) {
    ElMessage.error('清空购物车失败：' + err.message)
  }
}

const goToPayment = () => {
  router.push('/payment')
}

const updateQuantity = async (item) => {
  console.warn(item);
  try {
    await store.dispatch('cart/updateQuantity', {
      Id: item.Id,
      Qty: item.Qty
    })
    ElMessage.success('数量更新成功')
  } catch (err) {
    ElMessage.error('数量更新失败：' + err.message)
  }
}
</script>

<style scoped>
.el-main {
  padding: 20px;
}

.cart-actions {
  margin-top: 20px;
  text-align: right;
}

.cart-actions .el-button {
  margin-left: 10px;
}
</style>
