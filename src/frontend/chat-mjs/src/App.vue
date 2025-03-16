<template>
  <div class="min-h-screen bg-[#1e1e2e] flex items-center justify-center p-4">
    <div class="w-full max-w-5xl flex">
      <!-- Logo Section -->
      <div class="mr-4">
        <ChatHeader />
      </div>
      
      <!-- Main Chat Section -->
      <div class="flex-1 flex flex-col">
        <!-- Chat Container -->
        <ChatContainer 
          :messages="messages" 
          :isThinking="isThinking" 
          ref="chatContainerRef" 
          class="flex-grow mb-4"
        />
        
        <!-- Input Area -->
        <div class="flex items-center">
          <ChatInput 
            v-model="userInput"
            :method="method"
            :isThinking="isThinking"
            @toggle-method="toggleMethod"
            @send-message="sendMessage"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ChatHeader from './components/ChatHeader.vue';
import ChatContainer from './components/ChatContainer.vue';
import ChatInput from './components/ChatInput.vue';

export default {
  name: 'App',
  components: {
    ChatHeader,
    ChatContainer,
    ChatInput
  },
  data() {
    return {
      userInput: '',
      messages: [
        { text: 'Saya sedang duduk di singgasana legacy', isUser: true },
        { text: 'Saya sedang duduk di singgasana legacy', isUser: false }
      ],
      method: 'kmp',
      isThinking: false,
      typingSpeed: 30,
      apiUrl: 'http://localhost:8000/api/gpt'
    }
  },
  methods: {
    toggleMethod() {
      this.method = this.method === 'kmp' ? 'bm' : 'kmp';
    },
    async sendMessage(message) {
      if (!message || this.isThinking) return;
      
      // Add user message to chat
      this.messages.push({ text: message, isUser: true });
      this.userInput = '';
      
      // Scroll to bottom
      this.$nextTick(() => {
        this.scrollToBottom();
      });
      
      // Show thinking animation
      this.isThinking = true;
      
      try {
        // Send request to API
        const response = await fetch(this.apiUrl, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            pertanyaan: message,
            method: this.method
          })
        });
        
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        
        const data = await response.json();
        
        // Hide thinking animation after a minimum time (for better UX)
        setTimeout(() => {
          this.isThinking = false;
          
          // Add empty bot message that will be filled character by character
          const botMessageIndex = this.messages.length;
          this.messages.push({ text: '', isUser: false });
          
          // Simulate typing effect
          this.typeMessage(data.response, botMessageIndex);
        }, Math.max(500, Math.min(data.response.length * 20, 2000)));
        
      } catch (error) {
        console.error('Error:', error);
        
        // Hide thinking animation
        this.isThinking = false;
        
        // Add error message
        this.messages.push({ 
          text: 'Sorry, there was an error processing your request.', 
          isUser: false 
        });
        
        // Scroll to bottom
        this.$nextTick(() => {
          this.scrollToBottom();
        });
      }
    },
    typeMessage(fullText, messageIndex) {
      let currentIndex = 0;
      const typingInterval = setInterval(() => {
        if (currentIndex < fullText.length) {
          // Add one character at a time
          this.messages[messageIndex].text = fullText.substring(0, currentIndex + 1);
          currentIndex++;
          
          // Scroll to bottom as text is being typed
          this.scrollToBottom();
        } else {
          // Typing complete
          clearInterval(typingInterval);
        }
      }, this.typingSpeed);
    },
    scrollToBottom() {
      if (this.$refs.chatContainerRef) {
        this.$refs.chatContainerRef.scrollToBottom();
      }
    }
  }
}
</script>

<style>
@import 'https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css';

/* Global styles */
body {
  margin: 0;
  padding: 0;
  background-color: #1e1e2e;
  color: white;
  font-family: Arial, sans-serif;
}

::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #1e1e2e;
}

::-webkit-scrollbar-thumb {
  background: #3b3b4f;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #4b4b63;
}
</style>