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
            <el-avatar :size="36" :src="message.role === 'assistant' ? '/bot-avatar.png' : '/user-avatar.png'" />
          </div>
          <div class="message-content">
            {{ message.content }}
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

// 发送消息
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
    // 保持原有的发送逻辑
    const response = await fetch('/api/chat', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ message: userMessage })
    })
    const data = await response.json()

    // 添加助手回复
    messages.value.push({
      role: 'assistant',
      content: data.message
    })
  } catch (error) {
    console.error('发送消息失败:', error)
  } finally {
    loading.value = false
  }
}

// 监听消息变化，自动滚动到底部
watch(messages, () => {
  nextTick(() => {
    if (messageContainer.value) {
      messageContainer.value.scrollTop = messageContainer.value.scrollHeight
    }
  })
})
</script>

<style scoped>
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