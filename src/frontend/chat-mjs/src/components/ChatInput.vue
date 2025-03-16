<template>
    <div class="flex items-center w-full">
      <!-- KMP/BM Toggle -->
      <ToggleSwitch 
        :value="method" 
        leftLabel="KMP" 
        rightLabel="BM"
        @toggle="$emit('toggle-method')" 
      />
      
      <!-- Input Field -->
      <div class="flex-1 mx-2">
        <input 
          :value="modelValue"
          @input="$emit('update:modelValue', $event.target.value)"
          @keyup.enter="handleSend"
          type="text" 
          placeholder="Type your message here..." 
          class="w-full bg-gray-700 text-white rounded-lg p-3 focus:outline-none focus:ring-2 focus:ring-blue-500 border border-blue-500"
          :disabled="isThinking"
          style="background-color: #374151;"
        />
      </div>
      
      <!-- Send Button -->
      <button 
        @click="handleSend" 
        class="bg-teal-600 hover:bg-teal-700 text-white rounded-full p-3 focus:outline-none"
        :disabled="isThinking"
        style="background-color: #0d9488;"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 transform rotate-90" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
        </svg>
      </button>
    </div>
  </template>
  
  <script>
  import ToggleSwitch from './ToggleSwitch.vue';
  
  export default {
    name: 'ChatInput',
    components: {
      ToggleSwitch
    },
    props: {
      modelValue: {
        type: String,
        default: ''
      },
      method: {
        type: String,
        default: 'kmp'
      },
      isThinking: {
        type: Boolean,
        default: false
      }
    },
    methods: {
      handleSend() {
        if (!this.modelValue.trim() || this.isThinking) return;
        this.$emit('send-message', this.modelValue.trim());
      }
    }
  }
  </script>
  
  <style scoped>
  button:disabled, input:disabled {
    opacity: 0.7;
    cursor: not-allowed;
  }
  
  .border-blue-500 {
    border-color: #3b82f6;
  }
  </style>