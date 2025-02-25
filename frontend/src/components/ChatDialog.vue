<template>
  <div class="chat-container">
    <!-- 聊天按钮 -->
    <div class="chat-trigger-wrapper" @mouseenter="showHint = true" @mouseleave="showHint = false">
      <el-button
        class="chat-trigger"
        type="primary"
        circle
        size="large"
        @click="toggleChat"
      >
        <el-icon v-if="!visible" class="trigger-icon"><ChatDotRound /></el-icon>
        <el-icon v-else class="trigger-icon"><Close /></el-icon>
      </el-button>
      
      <!-- 悬浮提示 -->
      <div class="chat-hint" :class="{ 'hint-show': showHint && !visible }">
        <span>联系客服</span>
        <div class="hint-arrow"></div>
      </div>
    </div>

    <!-- 聊天窗口 -->
    <div class="chat-window" :class="{ 'chat-window-show': visible }">
      <div class="chat-header">
        <span>智能客服</span>
        <el-button type="text" @click="visible = false">
          <el-icon><Close /></el-icon>
        </el-button>
      </div>

      <div class="chat-messages" ref="messageContainer">
        <div v-for="(message, index) in messages" :key="index" 
          :class="['message', message.role === 'assistant' ? 'message-left' : 'message-right']">
          <div class="message-avatar">
            <el-avatar :size="36" :src="message.role === 'assistant' ? 'static/image/logo.jpg' : 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'" />
          </div>
          <div class="message-content">
            <div v-html="marked(message.content)"></div>
          </div>
        </div>
      </div>

      <div class="chat-input">
        <el-input
          v-model="inputMessage"
          placeholder="请输入消息..."
          :disabled="loading"
          @keyup.enter="sendMessage"
        >
          <template #append>
            <el-button :loading="loading" @click="sendMessage">发送</el-button>
          </template>
        </el-input>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, nextTick } from 'vue'
import { ChatDotRound, Close } from '@element-plus/icons-vue'
import {marked} from 'marked';

const visible = ref(false)
const showHint = ref(false)
const loading = ref(false)
const inputMessage = ref('')
const messages = ref([
  {
    role: 'assistant',
    content: '您好！我是智能客服助手，请问有什么可以帮您？'
  }
])
const messageContainer = ref(null)

const toggleChat = () => {
  visible.value = !visible.value
}

// 在 script setup 中添加
const typewriterEffect = (text, callback) => {
  let index = 0
  const interval = setInterval(() => {
    if (index < text.length) {
      callback(text.slice(0, index + 1))
      index++
    } else {
      clearInterval(interval)
    }
  }, 50) // 控制打字速度，可以调整
  return interval
}

// 修改 sendMessage 函数中的数据处理部分
const sendMessage = async () => {
  if (!inputMessage.value.trim() || loading.value) return

  // 添加用户消息
  messages.value.push({
    role: 'user',
    content: inputMessage.value
  })

  loading.value = true
  const userMessage = inputMessage.value
  inputMessage.value = ''

  try {
    // 创建一个新的消息对象用于存储助手的回复
    const assistantMessage = {
      role: 'assistant',
      content: ''
    }
    messages.value.push(assistantMessage)

    // 使用 fetch 发送 POST 请求
    const response = await fetch(`agent/api/chat?message=${encodeURIComponent(userMessage)}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'text/event-stream',
      'Cache-Control': 'no-cache',
      'Connection': 'keep-alive',
      'Cookie': 'cloudwego-shop=' + document.cookie.split('cloudwego-shop=')[1]?.split(';')[0]
    },
    credentials: 'include'
  })

    // 创建 reader 来读取流数据
    const reader = response.body.getReader()
    const decoder = new TextDecoder()

    const readChunk = () => {
      reader.read().then(({ value, done }) => {
        if (done) {
          loading.value = false
          return
        }

        const chunk = decoder.decode(value)
        const lines = chunk.split('\n')
        assistantMessage.content += chunk.replaceAll('data:', '').replaceAll('\n\n', '')
        // 确保滚动到底部
        nextTick(() => {
          if (messageContainer.value) {
            messageContainer.value.scrollTop = messageContainer.value.scrollHeight
          }
        })
        // for (const line of lines) {
        //   if (line.startsWith('data:')) {
        //     const content = line.slice(5) 
        //     if (content) {
        //       // 累积内容并更新消息
        //       console.log("content: " + content)
        //       const formattedContent = content.includes("- ") 
        // ? content.replace(/- /g, '\n- ') // 在每个 "- " 前添加换行符
        // : content
              
        //     }
        //   }
        // }

        // 递归继续读取
        readChunk()
        console.log('assistantMessage:', assistantMessage)
      }).catch(error => {
        console.error('读取流错误:', error)
        loading.value = false
      })
    }

    // 开始读取流
    readChunk()
  } catch (error) {
    console.error('发送消息失败:', error)
    loading.value = false
  }
}
</script>

<style scoped>

.message-left .message-content::after {
  content: '';
  position: absolute;
  right: -2px;
  top: 50%;
  transform: translateY(-50%);
  width: 2px;
  height: 16px;
  background-color: var(--el-color-primary);
  animation: cursor-blink 0.8s steps(2) infinite;
  opacity: 0;
}

.message-left:last-child .message-content::after {
  opacity: 1;
}

@keyframes cursor-blink {
  0% { opacity: 0; }
  100% { opacity: 1; }
}

.chat-container {
  position: fixed;
  right: 20px;
  bottom: 20px;
  z-index: 2000;
}

.chat-trigger-wrapper {
  position: relative;
  display: inline-block;
}

.chat-trigger {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, var(--el-color-primary), var(--el-color-primary-light-3));
  border: none;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  transition: all 0.3s ease;
}

.chat-trigger:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.2);
  background: linear-gradient(135deg, var(--el-color-primary-light-3), var(--el-color-primary));
}

.trigger-icon {
  font-size: 24px;
  transition: all 0.3s ease;
}

.chat-trigger:hover .trigger-icon {
  transform: scale(1.1);
}

/* 悬浮提示样式 */
.chat-hint {
  position: absolute;
  right: 75px;
  top: 50%;
  transform: translateY(-50%) scale(0);
  background: white;
  padding: 8px 16px;
  border-radius: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  font-size: 14px;
  color: var(--el-text-color-primary);
  white-space: nowrap;
  opacity: 0;
  transition: all 0.3s ease;
  z-index: 1;
}

.hint-show {
  transform: translateY(-50%) scale(1);
  opacity: 1;
}

.hint-arrow {
  position: absolute;
  right: -6px;
  top: 50%;
  transform: translateY(-50%) rotate(45deg);
  width: 12px;
  height: 12px;
  background: white;
}

/* 优化对话框样式 */
:deep(.chat-dialog) {
  border-radius: 12px;
  overflow: hidden;
}

:deep(.el-dialog__header) {
  margin: 0;
  padding: 16px 20px;
  background: var(--el-color-primary);
}

:deep(.el-dialog__title) {
  color: white;
  font-size: 16px;
  font-weight: 500;
}

:deep(.el-dialog__headerbtn) {
  top: 16px;
}

:deep(.el-dialog__headerbtn .el-dialog__close) {
  color: white;
}

:deep(.el-dialog__body) {
  padding: 20px;
}

.chat-window {
  position: absolute;
  right: 0;
  bottom: 80px;
  width: 360px;
  height: 500px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  transform: scale(0);
  transform-origin: bottom right;
  transition: transform 0.3s ease;
  overflow: hidden;
}

.chat-window-show {
  transform: scale(1);
}

.chat-header {
  padding: 16px 20px;
  background: var(--el-color-primary);
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chat-header .el-button {
  color: white;
  font-size: 18px;
}

.chat-messages {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  background: #f5f7fa;
}

.message {
  display: flex;
  align-items: flex-start;
  margin-bottom: 20px;
  gap: 12px;
}

.message-left {
  flex-direction: row;
}

.message-right {
  flex-direction: row-reverse;
}

.message-content {
  max-width: 70%;
  padding: 12px 16px;
  border-radius: 12px;
  word-break: break-word;
  line-height: 1.4;
}

.message-left .message-content {
  background: white;
  border: 1px solid #e4e7ed;
  border-top-left-radius: 4px;
}

.message-right .message-content {
  background: var(--el-color-primary);
  color: white;
  border-top-right-radius: 4px;
}

.chat-input {
  padding: 16px;
  background: white;
  border-top: 1px solid #e4e7ed;
}

/* 滚动条样式 */
.chat-messages::-webkit-scrollbar {
  width: 6px;
}

.chat-messages::-webkit-scrollbar-thumb {
  background-color: #dcdfe6;
  border-radius: 3px;
}

.chat-messages::-webkit-scrollbar-track {
  background-color: transparent;
}

/* 消息输入框样式 */
:deep(.el-input-group__append) {
  padding: 0;
}

:deep(.el-input-group__append button) {
  border: none;
  margin: -1px;
  border-radius: 0 4px 4px 0;
}
</style>