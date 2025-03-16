<template>
  <div class="bg-[#1e1e2e] rounded-full p-2 flex items-center border border-blue-500">
    <span class="text-white text-sm mr-2">{{ leftLabel }}</span>
    <div class="w-12 h-6 bg-gray-700 rounded-full p-1 cursor-pointer" @click="$emit('toggle')" style="background-color: #374151;">
      <div class="bg-green-400 w-4 h-4 rounded-full transform duration-300 ease-in-out"
           :class="value === leftValue ? 'translate-x-0' : 'translate-x-6'"
           style="background-color: #4ade80;">
      </div>
    </div>
    <span class="text-white text-sm ml-2">{{ rightLabel }}</span>
  </div>
</template>

<script>
export default {
  name: 'ToggleSwitch',
  props: {
    value: {
      type: String,
      required: true
    },
    leftValue: {
      type: String,
      default: 'kmp'
    },
    rightValue: {
      type: String,
      default: 'bm'
    },
    leftLabel: {
      type: String,
      default: 'Left'
    },
    rightLabel: {
      type: String,
      default: 'Right'
    }
  }
}
</script>

<style scoped>
.bg-\[\#1e1e2e\] {
  background-color: #1e1e2e;
}

.border-blue-500 {
  border-color: #3b82f6;
}
</style>