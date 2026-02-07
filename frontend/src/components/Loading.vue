<script setup>
import {reactive} from "vue";

const props = defineProps({
  show: {
    type: Boolean,
    default: true
  }
})
</script>


<template>
  <div v-show="props.show" class="loading-container" style="background-color: #303133">
    <div class="loading-body">
      <div class="loading-animation">
        <div class="loading-animation-item" style="background-color:white"></div>
        <div class="loading-animation-item" style="background-color:white"></div>
        <div class="loading-animation-item" style="background-color:white"></div>
        <div class="loading-animation-item" style="background-color:white"></div>
        <div class="loading-animation-item" style="background-color:white"></div>
        <div class="loading-animation-item" style="background-color:white"></div>
        <div class="loading-animation-item" style="background-color:white"></div>
        <div class="loading-animation-item" style="background-color:white"></div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@keyframes loading {
  0% {
    height: 0;
  }

  50% {
    height: 50px;
  }

  100% {
    height: 0;
  }
}

.loading-container {
  position: fixed;
  left: 0;
  top: 0;
  width: 100vw;
  height: 100vh;
  background-color: white;
  z-index: 9999;

  .loading-body {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);

    .loading-animation {
      height: 40px;
      display: flex;
      align-items: center;
      justify-content: center;

      .loading-animation-item {
        height: 0;
        width: 5px;
        margin: 0 3px;
        border-radius: 10px;
        animation: loading 1s infinite;
      }

      .loading-animation-item:nth-child(2) {
        animation-delay: 0.1s;
      }

      .loading-animation-item:nth-child(3) {
        animation-delay: 0.2s;
      }

      .loading-animation-item:nth-child(4) {
        animation-delay: 0.3s;
      }

      .loading-animation-item:nth-child(5) {
        animation-delay: 0.4s;
      }

      .loading-animation-item:nth-child(6) {
        animation-delay: 0.5s;
      }

      .loading-animation-item:nth-child(7) {
        animation-delay: 0.6s;
      }

      .loading-animation-item:nth-child(8) {
        animation-delay: 0.7s;
      }
    }
  }
}
</style>