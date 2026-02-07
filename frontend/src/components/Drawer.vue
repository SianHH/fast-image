<script setup>

import {onBeforeMount, onBeforeUpdate, ref, watch} from "vue";

const emit = defineEmits(['close'])
const drawerContainerRef = ref()
const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  width: {
    type: [Number, String],
    default: 400,
  },
  to: {
    type: Element,
    default: document.querySelector('body')
  },
  maskClosable: {
    type: Boolean,
    default: false
  }
})

const state = ref({
  show: props.show,
})

onBeforeMount(() => {
  state.value.show = props.show
})

onBeforeUpdate(() => {
  state.value.show = props.show
})

watch(state.value, (val) => {
  state.value.show = val.show
  emit('close', val.show)
})

</script>

<template>
  <div ref="drawerContainerRef" class="drawer-container"></div>
  <n-drawer v-model:show="state.show" :width="props.width" :to="props.to" :mask-closable="props.maskClosable"
            :autoFocus="false">
    <n-drawer-content header-class="header" :body-content-style="{padding: '0 !important'}">
      <template #header>
        <slot name="header"/>
      </template>
      <n-scrollbar style="max-height: calc(100vh - 51px - 67px)">
        <slot/>
      </n-scrollbar>
      <template #footer>
        <slot name="footer"/>
      </template>
    </n-drawer-content>
  </n-drawer>
</template>

<style lang="scss" scoped>
:deep(.n-drawer-header) {
  height: 60px !important;
  box-sizing: border-box !important;
}
</style>
