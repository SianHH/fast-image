<script setup>

import {computed, onBeforeMount, ref} from "vue";
import {copyToClipboardDirect} from "../../utils/copy.js";
import {apiImageDelete} from "../../api/image/index.js";

const emits = defineEmits(['refresh'])

const props = defineProps({
  data: {
    required: true,
    type: Object,
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

const show = ref(false)
const openModal = () => {
  show.value = true
}
const closeModal = () => {
  show.value = false
}

const copyContent = (row, tp) => {
  let url = `${window.location.protocol}//${window.location.host}/api/v1/image/${row.filename}`
  switch (tp) {
    case 'url':
      return url
    case 'html':
      return `<img src='${url}' alt=''>`
    case 'md':
      return `![${row.filename}](${url})`
  }
  return ''
}

const deleteLoading = ref(false)
const deleteFunc = async () => {
  try {
    deleteLoading.value = true
    await apiImageDelete({id: props.data.id})
    emits('refresh')
    closeModal()
  } finally {
    deleteLoading.value = false
  }
}

onBeforeMount(() => {
})
</script>

<template>
  <div>
    <div @click="openModal">
      <slot></slot>
    </div>
    <n-modal :show="show" :mask-closable="props.maskClosable" :auto-focus="props.autoFocus">
      <div :style="cardStyleComputed">
        <n-card size="small">
          <template #header>
            <span style="font-weight: bold">预览图片</span>
          </template>
          <template #header-extra>
            <n-space>
              <n-button size="small" type="error" :focusable="false" @click="deleteFunc" :loading="deleteLoading">
                删除
              </n-button>
              <n-button size="small" type="default" :focusable="false" @click="closeModal">关闭</n-button>
            </n-space>
          </template>
          <n-scrollbar style="max-height: calc(100vh - 140px)">
            <div style="aspect-ratio: 1 / 0.6;border-radius: 8px;overflow: hidden">
              <img class="images" style="width: 100%;height: 100%;object-fit: cover"
                   :src="`/api/v1/image/${props.data.filename}`" alt="">
            </div>
            <n-space size="small" vertical style="margin-top: 10px">
              <n-input-group>
                <n-input-group-label style="width: 70px">URL</n-input-group-label>
                <n-input
                    @click="copyToClipboardDirect(copyContent(props.data,'url'))"
                    :default-value="copyContent(props.data,'url')" readonly
                />
              </n-input-group>
              <n-input-group>
                <n-input-group-label style="width: 70px">HTML</n-input-group-label>
                <n-input
                    @click="copyToClipboardDirect(copyContent(props.data,'html'))"
                    :default-value="copyContent(props.data,'html')" readonly
                />
              </n-input-group>
              <n-input-group>
                <n-input-group-label style="width: 70px">MD</n-input-group-label>
                <n-input
                    @click="copyToClipboardDirect(copyContent(props.data,'md'))"
                    :default-value="copyContent(props.data,'md')" readonly
                />
              </n-input-group>
            </n-space>
          </n-scrollbar>
        </n-card>
      </div>
    </n-modal>
  </div>
</template>

<style scoped lang="scss">

</style>