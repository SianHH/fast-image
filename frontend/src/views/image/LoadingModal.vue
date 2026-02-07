<script setup>

import {computed, onBeforeMount} from "vue";

const props = defineProps({
  title: {
    required: true,
    type: String,
  },
  show: {
    required: true,
    type: Boolean,
  },
  width: {
    default: '90%',
    type: String,
  },
  maxWidth: {
    default: '600px',
    type: String,
  },
  maskClosable: {
    default: false,
    type: Boolean,
  },
  autoFocus: {
    default: false,
    type: Boolean,
  },
})

const cardStyleComputed = computed(() => {
  return {
    width: props.width,
    maxWidth: props.maxWidth,
    maxHeight: 'calc(100vh - 80px)'
  }
})

onBeforeMount(() => {
})
</script>

<template>
  <n-modal :show="show" :mask-closable="props.maskClosable" :auto-focus="props.autoFocus">
    <div :style="cardStyleComputed">
      <n-card size="small">
        <template #header>
          <span style="font-weight: bold">{{ props.title }}</span>
        </template>
        <div style="display: flex;align-items: center;justify-content: center">
          <n-spin style="margin: 40px 0" :show="true" size="large"></n-spin>
        </div>
      </n-card>
    </div>
  </n-modal>
</template>

<style scoped lang="scss">

</style>