<template>
  <el-container>
    <el-main>
      <el-row>
        <el-col :span="24">
          <h2>购物车</h2>
          <el-empty v-if="cartItems.length === 0" description="购物车为空" />
          <el-table v-else :data="cartItems" style="width: 100%">
            <el-table-column prop="name" label="商品名称" />
            <el-table-column prop="price" label="价格" width="120" />
            <el-table-column prop="quantity" label="数量" width="120">
              <template #default="{ row }">
                <el-input-number 
                  v-model="row.quantity" 
                  :min="1" 
                  :max="10" 
                  size="small"
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
              :disabled="cartItems.length === 0"
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
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

const router = useRouter()
const store = useStore()

const cartItems = computed(() => store.getters['cart/items'])
const totalAmount = computed(() => store.getters['cart/totalAmount'])

const removeItem = (item) => {
  store.dispatch('cart/removeItem', item.id)
}

const clearCart = () => {
  store.dispatch('cart/clearCart')
}

const goToPayment = () => {
  router.push('/payment')
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
