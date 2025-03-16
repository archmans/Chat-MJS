<template>
  <div class="flex" :class="message.isUser ? 'justify-start' : 'justify-end'">
    <div 
      class="inline-block p-3 rounded-lg mb-3 text-white"
      :class="message.isUser ? 'rounded-tr-2xl rounded-bl-2xl rounded-br-2xl' : 'rounded-tl-2xl rounded-bl-2xl rounded-br-2xl'"
      :style="{
        backgroundColor: '#0f766e',
        maxWidth: '80%',
        wordBreak: 'break-word'
      }"
    >
      {{ message.text }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'MessageBubble',
  props: {
    message: {
      type: Object,
      required: true
    }
  }
}
</script>

<style scoped>
/* Force the teal color for message bubbles */
</style>