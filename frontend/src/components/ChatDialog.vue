<template>
  <div>
    <div class="chat-box">
      <div class="chat-content" ref="chatContent">
        <div v-for="(chat, index) in chatHistory" :key="index" class="chat-message">
          <p><strong>User:</strong> {{ chat.message }}</p>
          <p><strong>Response:</strong> {{ chat.response }}</p>
        </div>
      </div>
    </div>
    <input v-model="message" placeholder="Type your message here" />
    <button @click="sendMessage">Send</button>
    <button @click="clearChatHistory">Clear</button>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick, computed, watch } from 'vue'
import { useStore } from 'vuex'

const store = useStore()
const message = ref('')

const chatHistory = computed(() => store.getters.chatHistory)

const sendMessage = () => {
  if (message.value.trim() === '') return

  // 先将用户的消息添加到 chatHistory 中
  store.dispatch('addMessage', {
    message: message.value,
    response: ''  // 先留空，稍后填入服务器的响应
  })

  const eventSource = new EventSource('/api/chat')
  eventSource.onmessage = (event) => {
    const data = JSON.parse(event.data)
    // 更新最后一条消息的响应
    store.dispatch('updateLastResponse', data.response)
    nextTick(() => {
      scrollToBottom()
    })
  }

  eventSource.onerror = (err) => {
    console.error(err)
    eventSource.close()
  }

  message.value = ''
  nextTick(() => {
    scrollToBottom()
  })
}

const clearChatHistory = () => {
  store.dispatch('clearChatHistory')
}

const scrollToBottom = () => {
  const chatContent = document.querySelector('.chat-content')
  chatContent.scrollTop = chatContent.scrollHeight
}

onMounted(() => {
  nextTick(() => {
    scrollToBottom()
  })
})

watch(chatHistory, () => {
  nextTick(() => {
    scrollToBottom()
  })
})
</script>

<style scoped>
.chat-box {
  max-height: 200px;
  overflow-y: auto;
  border: 1px solid #ccc;
  padding: 10px;
  margin-bottom: 10px;
}

.chat-content {
  display: flex;
  flex-direction: column;
}

.chat-message {
  margin-bottom: 10px;
}
</style>