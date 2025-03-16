<template>
  <div class="h-screen bg-[#1e1e2e] flex p-4 overflow-hidden">
    <!-- Include the Popup -->
    <PopupTutorial />

    <!-- Left Panel - Logo and Toggle -->
    <div class="w-[200px] mr-4 h-full">
      <div class="border-2 border-blue-500 rounded-xl bg-[#1e1e2e] h-full flex flex-col">
        <div class="flex-grow flex flex-col items-center justify-center px-16">
          <img src="/src/components/icons/MJS-LOGO.svg" alt="MJS Logo" class="w-48 h-48" />
        </div>

        <div class="p-4">
          <div class="bg-[#1e1e2e] rounded-full p-4 flex items-center border border-blue-500 justify-center">
            <span class="text-white text-sm mr-2">Knuth-Morris</span>
            <div
              class="w-12 h-6 bg-gray-700 rounded-full p-1 cursor-pointer"
              @click="toggleMethod"
            >
              <div
                class="w-4 h-4 rounded-full transform duration-300 ease-in-out"
                :class="method === 'kmp' ? 'translate-x-0' : 'translate-x-6'"
                style="background-color: #0d9488"
              ></div>
            </div>
            <span class="text-white text-sm ml-2">Boyer-Moore</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Right Panel - Chat Section -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <div class="flex-1 overflow-y-auto py-4 px-48" ref="messagesContainer">
        <div
          v-for="(message, index) in messages"
          :key="index"
          class="flex"
          :class="message.isUser ? 'justify-start' : 'justify-end'"
        >
          <div
            class="inline-block p-3 rounded-lg mb-3 text-white"
            :class="message.isUser ? 'rounded-tr-2xl rounded-bl-2xl rounded-br-2xl' : 'rounded-tl-2xl rounded-bl-2xl rounded-br-2xl'"
            :style="{ backgroundColor: message.isUser ? '#0f766e' : '#4854FE', maxWidth: '80%', wordBreak: 'break-word' }"
          >
            {{ message.text }}
          </div>
        </div>

        <div v-if="isThinking" class="flex justify-end">
          <div
            class="inline-block p-3 rounded-lg mb-3 text-white rounded-tl-2xl rounded-bl-2xl rounded-br-2xl"
            style="background-color: #4854FE"
          >
            <div class="flex items-center">
              <div class="thinking-dots">
                <span></span>
                <span></span>
                <span></span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Fixed Input -->
      <div class="bg-[#252538] px-48 py-4">
        <div class="flex">
          <input
            v-model="userInput"
            @keyup.enter="sendMessage"
            type="text"
            placeholder="Type your message here..."
            class="flex-1 bg-gray-700 text-white rounded-3xl p-3 focus:outline-none border border-blue-500"
          />

          <button
            @click="sendMessage"
            class="ml-2 bg-teal-600 hover:bg-teal-700 text-white rounded-full p-3"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6 transform rotate-90"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Popup from './components/icons/Popup.vue'

export default {
  name: 'App',
  components: {
    PopupTutorial: Popup,
  },
  data() {
    return {
      userInput: '',
      messages: [],
      method: 'kmp',
      isThinking: false,
      apiUrl: 'http://localhost:8000/api/gpt',
    }
  },
  methods: {
    toggleMethod() {
      this.method = this.method === 'kmp' ? 'bm' : 'kmp'
    },
    async sendMessage() {
      if (!this.userInput.trim() || this.isThinking) return

      const userMessage = this.userInput.trim()
      this.messages.push({ text: userMessage, isUser: true })
      this.userInput = ''
      this.$nextTick(this.scrollToBottom)

      this.isThinking = true
      try {
        const response = await fetch(this.apiUrl, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ pertanyaan: userMessage, method: this.method }),
        })
        if (!response.ok) throw new Error('Network error')
        const data = await response.json()
        setTimeout(() => {
          this.isThinking = false
          const botMessageIndex = this.messages.length
          this.messages.push({ text: '', isUser: false })
          this.typeMessage(data.response, botMessageIndex)
        }, 1000)
      } catch (error) {
        this.isThinking = false
        this.messages.push({ text: 'Error processing request.', isUser: false })
        this.$nextTick(this.scrollToBottom)
      }
    },
    typeMessage(fullText, messageIndex) {
      let currentIndex = 0
      const typingSpeed = 30
      const typingInterval = setInterval(() => {
        if (currentIndex < fullText.length) {
          this.messages[messageIndex].text = fullText.substring(0, currentIndex + 1)
          currentIndex++
          this.$nextTick(this.scrollToBottom)
        } else {
          clearInterval(typingInterval)
        }
      }, typingSpeed)
    },
    scrollToBottom() {
      if (this.$refs.messagesContainer) {
        this.$refs.messagesContainer.scrollTop = this.$refs.messagesContainer.scrollHeight
      }
    },
  },
  mounted() {
    this.scrollToBottom()
  },
}
</script>

<style>
html,
body,
#app {
  height: 100%;
  margin: 0;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  background-color: #1e1e2e;
}

.flex-1 {
  flex: 1;
  min-height: 0;
}

.thinking-dots {
  display: flex;
  align-items: center;
}

.thinking-dots span {
  height: 8px;
  width: 8px;
  margin: 0 2px;
  background-color: white;
  border-radius: 50%;
  display: inline-block;
  opacity: 0.7;
  animation: thinking 1.4s infinite ease-in-out both;
}

.thinking-dots span:nth-child(1) {
  animation-delay: -0.32s;
}

.thinking-dots span:nth-child(2) {
  animation-delay: -0.16s;
}

@keyframes thinking {
  0%,
  80%,
  100% {
    transform: scale(0);
  }
  40% {
    transform: scale(1);
  }
}

@media (max-width: 768px) {
  .h-screen {
    height: auto;
  }
}
</style>
