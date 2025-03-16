<template>
  <div class="flex justify-start">
    <div 
      class="inline-block p-3 rounded-lg mb-3 text-white rounded-tr-2xl rounded-bl-2xl rounded-br-2xl"
      style="background-color: #0f766e;"
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
</template>

<script>
export default {
  name: 'ThinkingIndicator'
}
</script>

<style scoped>
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
  0%, 80%, 100% { 
    transform: scale(0);
  } 40% { 
    transform: scale(1);
  }
}
</style>