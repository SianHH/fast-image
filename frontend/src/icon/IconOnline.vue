<script setup>
const props = defineProps({
  size: {
    default: 12,
    type: Number,
  },
  online: {
    default: false,
    type: Boolean,
  }
})

</script>

<template>
  <span :style="{
    display: 'inline-block',
    width: props.size + 'px',
    height: props.size + 'px',
    backgroundColor: props.online ? '#0F0' : '#F00',
    borderRadius: '50%'
  }"></span>
</template>

<style scoped lang="scss">

</style>