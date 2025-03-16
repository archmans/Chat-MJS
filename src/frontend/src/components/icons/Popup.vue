<template>
  <div v-if="showPopup" class="fixed inset-0 bg-black bg-opacity-60 flex items-center justify-center z-50">
    <div class="rounded-lg w-11/12 max-w-lg p-6 relative" style="background-color: #1e1e2e; box-shadow: 0 0 5px #0d9488, 0 0 20px #0d9488, 0 0 100px #0d9488;">

      <!-- Image Carousel -->
      <div class="relative w-full overflow-hidden h-72 rounded-lg items-center flex justify-center">
        <img :src="tutorialImages[currentIndex]" alt="Tutorial" class="w-full h-auto rounded-lg transition-opacity duration-700 ease-in-out" />
      </div>
      <!-- Text -->
      <div class="mt-4 flex justify-center text-white text-sm">
        <span
          v-if="currentIndex === 0"
        >
          You can add the answer by prompting the question with 'tambah pertanyaan' [the question] 'dengan jawaban' [the answer]
        </span>
        <span
          v-else-if="currentIndex === 1"
        >
          You also can delete the question by prompting the question with 'hapus pertanyaan' [the question]
        </span>
        <span
          v-else
        >
          You can ask Hari [dd/mm/yyyy] to get the event on that day.
          You can also ask equational math by prompting the equation with 'hasil' [the equation]
        </span>
      </div>


      <!-- Dots Indicator -->
      <div class="flex justify-center mt-4">
        <span
          v-for="(image, index) in tutorialImages"
          :key="index"
          class="w-3 h-3 mx-1 rounded-full transition-all duration-300"
          :class="currentIndex === index ? 'bg-gray-700 w-5' : 'bg-gray-400 w-3'"
        ></span>
      </div>

      <!-- Close Button -->
      <div class="mt-4 flex justify-center">
        <button @click="closePopup" class="text-white px-4 py-2 rounded-lg" style="background-color: rgb(72, 84, 254);">
          Got it!
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showPopup: true, // Show popup on page load
      currentIndex: 0,
      tutorialImages: [
        "/src/assets/tutorial1.png",
        "/src/assets/tutorial2.png",
        "/src/assets/tutorial3.png",
      ],
      intervalId: null,
    };
  },
  methods: {
    nextImage() {
      // Loop back to first image when reaching the last
      this.currentIndex = (this.currentIndex + 1) % this.tutorialImages.length;
    },
    prevImage() {
      if (this.currentIndex > 0) {
        this.currentIndex--;
      } else {
        this.currentIndex = this.tutorialImages.length - 1; // Go to last image if on first
      }
    },
    startAutoSlide() {
      this.intervalId = setInterval(() => {
        this.nextImage();
      }, 6000); // Move every 3 seconds
    },
    stopAutoSlide() {
      clearInterval(this.intervalId);
    },
    closePopup() {
      this.stopAutoSlide();
      this.showPopup = false;
    },
  },
  mounted() {
    this.startAutoSlide();
  },
  beforeUnmount() {
    this.stopAutoSlide();
  }
};
</script>

<style scoped>
/* Smooth transition effect */
.transition-opacity {
  transition: opacity 0.7s ease-in-out;
}

</style>
