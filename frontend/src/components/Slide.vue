<script setup>
import {onMounted, ref} from "vue";

const props = defineProps({
  visible: {
    type: Boolean,
    default: true,
  },
})

const state = ref({
  slide:true,
})

onMounted(()=>{
  setTimeout(()=>{
    state.value.slide = false
  },200)
})
</script>

<template>
  <transition name="fade">
    <div :class="{body:true,bodySlideIn:state.slide}" v-if="props.visible">
      <slot name="default"></slot>
    </div>
  </transition>
</template>

<style lang="scss" scoped>
.body {
  opacity: 1;
  transform: translateX(0px);
  transition: all 0.1s;
}

// 进入时动画
.fade-enter-active {
  transform: translateX(-10px);
  opacity: 0;
}

// 离开时动画
.fade-leave-active {
  display: none;
}

.bodySlideIn{
  transform: translateX(-10px);
  opacity: 0;
}
</style>