<script setup>

import {computed} from "vue";

const emit = defineEmits(['onCancel', 'onConfirm'])
const props = defineProps({
  title: {
    type: String,
    default: '标题',
  },
  confirmText: {
    type: String,
    default: '确认',
  },
  confirmLoading: {
    type: Boolean,
    default: false,
  },
  cancelText: {
    type: String,
    default: '取消',
  },
  autoFocus: {
    type: Boolean,
    default: false,
  },
  show: {
    type: Boolean,
    default: false,
  },
  width: {
    type: String,
    default: '600px',
  },
  maskClose: {
    type: Boolean,
    default: false,
  }
})

const cardStyleComputed = computed(() => {
  return {
    width: props.width,
    maxWidth: '90%',
  }
})

const maskClose = (ok) => {
  emit('onCancel')
}

</script>

<template>
  <n-modal :show="props.show" :mask-closable="props.maskClose" :on-update-show="maskClose"
           :auto-focus="props.autoFocus">
    <n-card :style="cardStyleComputed" :bordered="false">
      <template #header>
        <span style="font-weight: bold">{{ props.title }}</span>
      </template>
      <n-scrollbar style="max-height: calc(100vh - 142px - 10vw)">
        <slot></slot>
      </n-scrollbar>
      <template #footer v-if="!(props.cancelText === '' && props.confirmText === '')">
        <n-space justify="end" style="width: 100%">
          <n-button
              v-if="props.cancelText !== ''"
              type="default"
              :focusable="false"
              @click="()=>{emit('onCancel')}"
          >
            {{ props.cancelText }}
          </n-button>
          <n-button
              v-if="props.confirmText !== ''"
              type="success"
              :focusable="false"
              @click="()=>{emit('onConfirm')}"
              :loading="props.confirmLoading"
          >
            {{ props.confirmText }}
          </n-button>
        </n-space>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped lang="scss">
</style>