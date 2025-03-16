<template>
    <div class="border-2 border-blue-500 rounded-xl bg-[#252538] shadow-lg relative overflow-hidden h-[500px]">
      <!-- Neon Effect -->
      <div class="absolute inset-0 border-2 border-blue-500 rounded-xl blur-sm z-0"></div>
      
      <!-- Chat Messages - Only this area is scrollable -->
      <div class="relative z-10 h-full overflow-y-auto p-4" ref="messagesContainer">
        <!-- Message Bubbles -->
        <MessageBubble 
          v-for="(message, index) in messages" 
          :key="index"
          :message="message"
        />
        
        <!-- Thinking Indicator -->
        <ThinkingIndicator v-if="isThinking" />
      </div>
    </div>
  </template>
  
  <script>
  import MessageBubble from './MessageBubble.vue';
  import ThinkingIndicator from './ThinkingIndicator.vue';
  
  export default {
    name: 'ChatContainer',
    components: {
      MessageBubble,
      ThinkingIndicator
    },
    props: {
      messages: {
        type: Array,
        required: true
      },
      isThinking: {
        type: Boolean,
        default: false
      }
    },
    methods: {
      scrollToBottom() {
        if (this.$refs.messagesContainer) {
          this.$refs.messagesContainer.scrollTop = this.$refs.messagesContainer.scrollHeight;
        }
      }
    },
    mounted() {
      this.scrollToBottom();
    },
    updated() {
      this.scrollToBottom();
    }
  }
  </script>
  
  <style scoped>
  .border-blue-500 {
    border-color: #3b82f6;
  }
  
  .bg-\[\#252538\] {
    background-color: #252538;
  }
  </style>