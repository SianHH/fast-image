<script setup>
import {computed, onBeforeMount, ref, watch} from "vue";
import {apiImageDelete, apiImageList} from "../../api/image/index.js";
import moment from "moment";
import {localStore} from "../../store/local.js";
import {customRequest} from "./upload.js";
import {flowFormat} from "../../utils/flow.js";
import PreviewWrapperButton from "./PreviewWrapperButton.vue";
import {copyToClipboardDirect} from "../../utils/copy.js";
import LoadingModal from "./LoadingModal.vue";


const copyType = ref('md')
watch(() => copyType.value, (val) => {
  localStore().copyType = val
})

const copyContent = (row) => {
  let url = `${window.location.protocol}//${window.location.host}/api/v1/image/${row.filename}`
  switch (copyType.value) {
    case 'url':
      return url
    case 'html':
      return `<img src='${url}' alt=''>`
    case 'md':
      return `![${row.filename}](${url})`
  }
  return ''
}

const todayData = ref({
  list: [],
  search: {
    startOn: moment().format('YYYY-MM-DD'),
    endOn: moment().format('YYYY-MM-DD'),
  },
  loading: true,
  dateOnly: new Date().getTime(),
})

watch(() => todayData.value.dateOnly, (val) => {
  loadTodayImageList()
})

const loadTodayImageList = async () => {
  try {
    todayData.value.loading = true
    let dateOnly = moment(todayData.value.dateOnly).format('YYYY-MM-DD')
    todayData.value.search.startOn = dateOnly
    todayData.value.search.endOn = dateOnly
    let res = await apiImageList(todayData.value.search)
    todayData.value.list = res.data || []
    checkClear()
  } finally {
    todayData.value.loading = false
  }
}


const imageUploadCount = ref(0)
watch(() => imageUploadCount.value, (val) => {
  if (val === 0) {
    loadTodayImageList()
  }
})

const imageCustomUpload = async ({file, onFinish, onError}) => {
  try {
    await customRequest({
      file: file.file,
      onStart(file) {
        imageUploadCount.value++
      },
      onEnd(file) {
        imageUploadCount.value--
      },
      onFinish,
      onError,
      onResult: (data) => {
      }
    })
  } finally {
  }
}

const handlePaste = async (e) => {
  const clipboardItems = e.clipboardData?.items
  if (!clipboardItems) return

  for (const item of clipboardItems) {
    if (item.kind === 'file') {
      const file = item.getAsFile()
      if (file) {
        try {
          await customRequest({
            file,
            onStart(file) {
              imageUploadCount.value++
            },
            onEnd(file) {
              imageUploadCount.value--
            },
            onFinish() {
            },
            onError() {
            },
            onResult: (data) => {
            }
          })
        } finally {
        }
      }
    }
  }
}

const checkCountComputed = computed(() => {
  let count = 0
  todayData.value.list.forEach(item => {
    if (item.check === true) {
      count++
    }
  })
  return count
})

const checkClear = () => {
  todayData.value.list.forEach((item, index) => {
    todayData.value.list[index].check = false
  })
}

const checkCopy = () => {
  let content = []
  todayData.value.list.forEach(item => {
    if (item.check === true) {
      content.push(copyContent(item))
    }
  })
  copyToClipboardDirect(content.join('\n'))
}

const checkAll = () => {
  todayData.value.list.forEach((item, index) => {
    todayData.value.list[index].check = true
  })
}

const checkReverse = () => {
  todayData.value.list.forEach((item, index) => {
    todayData.value.list[index].check = !item.check
  })
}

const deleteLoading = ref(false)
const deleteFunc = async () => {
  try {
    deleteLoading.value = true
    let ids = []
    todayData.value.list.forEach((item) => {
      if (item.check === true) {
        ids.push(item.id)
      }
    })
    await apiImageDelete({ids})
    await loadTodayImageList()
    checkClear()
  } finally {
    deleteLoading.value = false
  }
}

onBeforeMount(() => {
  copyType.value = localStore().copyType
  loadTodayImageList()
})

</script>

<template>
  <n-card size="small" embedded style="min-height: calc(100vh - 60px)" @paste="handlePaste">
    <n-card size="small" style="width: 100%;max-width: 1200px;margin: 0 auto">
      <template #header>
        <span style="font-weight: bold">
          上传图片
        </span>
      </template>

      <n-upload
          multiple
          directory-dnd
          :custom-request="imageCustomUpload"
          accept=".jpg,.jpeg,.png,.gif,.webp"
          :show-file-list="false"
      >
        <n-upload-dragger>
          <div style="margin-bottom: 12px">
            <n-icon size="48" :depth="3">
              <svg t="1770347830755" class="icon" viewBox="0 0 1024 1024" version="1.1"
                   xmlns="http://www.w3.org/2000/svg"
                   p-id="11381">
                <path
                    d="M1008.400138 538.60651c0.550493-7.801856-0.526112-15.347072-4.87616-22.150598l-124.662878-200.233091c-10.555603-17.112755-34.142102-23.172025-52.050441-13.08864-17.859577 9.564973-24.362835 31.204857-14.057456 47.782518l97.0266 155.97424H746.346321c-18.159846 0-34.692595 12.559961-37.119126 30.18343-8.129072 44.270399-32.540668 83.515787-68.308584 111.702558-34.116438 27.65296-79.113128 44.021459-128.688277 44.021459-49.073417 0-94.621883-16.368499-129.288814-44.021459-34.691311-27.443798-59.078527-65.661343-67.233263-108.687038-1.076605-18.107235-17.333465-32.68567-36.843237-33.19895H114.109841l97.297356-155.980656c10.579984-16.577661 4.052346-38.217545-13.807232-47.782518-18.433168-10.076969-41.745062-4.017699-52.050441 13.08864L20.863548 516.455912a30.835296 30.835296 0 0 0-4.851779 22.150598 15.81159 15.81159 0 0 0-0.550492 4.016416v322.989134c0 19.387868 17.359129 35.73327 38.495999 35.73327h916.498638c20.860982 0 37.944223-16.345401 37.944224-35.73327V538.60651zM931.962481 830.158528H92.449426V577.822385h155.000292c14.62848 49.314658 45.522802 92.083714 86.168162 124.797614 47.94805 37.962188 110.554094 61.111116 178.612454 61.111115 68.009599 0 130.864584-23.148928 178.56241-61.111115 40.645359-32.7139 71.011004-75.482956 86.192543-124.797614H931.962481v252.336143z m-635.463727-464.903353h112.980625v217.117437c0 11.562915 10.305379 20.850716 23.037289 20.850716H591.619351c13.006515 0 22.735737-9.287801 22.735737-20.850716v-217.117437h113.280894c12.456022 0 22.787065-9.543158 22.787065-21.384528a20.777574 20.777574 0 0 0-6.528921-14.836358l-0.524829-0.25664-215.155424-199.717245c-8.705229-8.567926-23.587782-8.567926-32.015839 0L280.516498 329.039422c-8.954169 8.312569-8.954169 21.384528 0 29.952454a22.571488 22.571488 0 0 0 15.982256 6.268432z m215.73158-191.174983l160.129243 148.689515H591.619351c-12.456022 0-23.036006 9.543158-23.036006 21.10479v216.860797h-113.00629v-216.860797c0-11.561632-10.32976-21.10479-23.060387-21.10479h-80.465621l160.179287-148.689515z"
                    p-id="11382"></path>
              </svg>
            </n-icon>
          </div>
          <n-text style="font-size: 16px">
            点击或者拖动文件到该区域来上传
          </n-text>
          <n-p depth="3" style="margin: 8px 0 0 0">
            也可以使用Ctrl+V快捷键上传已复制的图片
          </n-p>
        </n-upload-dragger>
      </n-upload>
    </n-card>

    <div style="height: 10px"></div>


    <n-card size="small" style="width: 100%;max-width: 1200px;margin: 0 auto">
      <template #header>
        <span style="font-weight: bold">上传记录({{ todayData.list.length }})</span>
      </template>
      <template #header-extra>
        <n-space>
          <n-radio-group size="small" v-model:value="copyType">
            <n-radio-button size="small" value="url">URL</n-radio-button>
            <n-radio-button size="small" value="html">HTML</n-radio-button>
            <n-radio-button size="small" value="md">MD</n-radio-button>
          </n-radio-group>
        </n-space>
      </template>
      <n-space style="margin-bottom: 10px" justify="end">
        <n-button size="small" :focusable="false" @click="checkAll">全选</n-button>
        <n-button size="small" :focusable="false" @click="checkReverse">反选</n-button>
        <n-date-picker size="small" style="width: 110px" v-model:value="todayData.dateOnly" type="date"/>
      </n-space>
      <n-spin :show="todayData.loading">
        <n-empty v-if="todayData.list.length===0" description="无上传记录" style="margin: 50px 0"></n-empty>
        <n-grid v-else cols="1200:6 800:4 600:3 400:2 1" :x-gap="10" :y-gap="10">
          <n-grid-item v-for="row in todayData.list" :key="row.id">
            <n-card size="small" embedded style="box-shadow: 0 0 2px 2px rgba(0,0,0,0.04)">
              <template #header>
                <div style="opacity: 0.6;font-size: 0.7rem">
                  {{ row.dateOnly }} {{ flowFormat(row.size) }}
                </div>
              </template>
              <template #header-extra>
                <n-checkbox size="large" v-model:checked="row.check"/>
              </template>
              <div style="aspect-ratio: 1 / 0.6;border-radius: 8px;overflow: hidden">
                <PreviewWrapperButton :data="row" :mask-closable="true" @refresh="loadTodayImageList">
                  <img class="images" style="width: 100%;height: 100%;object-fit: cover"
                       :src="`/api/v1/image/${row.filename}`" alt="">
                </PreviewWrapperButton>
              </div>
              <div style="margin-top: 10px">
                <n-input
                    @click="copyToClipboardDirect(copyContent(row))"
                    v-if="copyType==='url'"
                    :default-value="copyContent(row)" size="small" readonly
                />
                <n-input
                    @click="copyToClipboardDirect(copyContent(row))"
                    v-if="copyType==='html'"
                    :default-value="copyContent(row)" size="small" readonly
                />
                <n-input
                    @click="copyToClipboardDirect(copyContent(row))"
                    v-if="copyType==='md'"
                    :default-value="copyContent(row)" size="small" readonly
                />
              </div>
            </n-card>
          </n-grid-item>
        </n-grid>
      </n-spin>
    </n-card>

    <n-popover trigger="manual" :show="true" :show-arrow="false" v-if="checkCountComputed !== 0"
               style="padding:0">
      <template #trigger>
        <div style="position:fixed;left: 50%;bottom: 20px"></div>
      </template>
      <n-card size="small" :bordered="false">
        <div style="margin-bottom: 10px">已选中{{ checkCountComputed }}项，可以进行以下操作</div>
        <n-grid :x-gap="10" :y-gap="10" cols="3">
          <n-grid-item>
            <n-button style="width:100%" size="small" type="default" @click="checkClear">清空</n-button>
          </n-grid-item>
          <n-grid-item>
            <n-button style="width:100%" size="small" type="error" @click="deleteFunc" :loading="deleteLoading">
              删除
            </n-button>
          </n-grid-item>
          <n-grid-item>
            <n-button style="width:100%" size="small" type="primary" @click="checkCopy">复制</n-button>
          </n-grid-item>
        </n-grid>
      </n-card>
    </n-popover>

    <LoadingModal :title="`上传中，剩余${imageUploadCount}项`" :show="imageUploadCount>0"/>
  </n-card>
</template>

<style scoped lang="scss">
.images {
  transition: transform ease-in-out 0.3s;

  &:hover {
    transform: scale(1.1);
  }
}
</style>