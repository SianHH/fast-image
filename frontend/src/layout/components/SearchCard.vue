<script setup>
const props = defineProps({
  showBorder: {
    type: Boolean,
    default: true,
  },
  space: {
    type: Boolean,
    default: false,
  },
  vertical: {
    type: Boolean,
    default: false
  }
})
</script>


<template>
  <n-card size="small">
    <n-space v-if="props.space" align="center" :vertical="props.vertical">
      <slot></slot>
    </n-space>
    <slot v-else></slot>
  </n-card>
</template>

<style lang="scss" scoped>
</style>