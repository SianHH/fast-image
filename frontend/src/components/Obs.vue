<script setup>
import {computed, onBeforeUnmount, onMounted, onUnmounted, ref, watch} from "vue";
import AppCard from "../layout/components/AppCard.vue";
import {resizeDirective} from "../utils/resize.js";
import {
  CategoryScale,
  Chart,
  Filler,
  Legend,
  LinearScale,
  LineController,
  LineElement,
  PointElement,
  Tooltip
} from 'chart.js'

// 注册需要的模块
Chart.register(LineController, LineElement, PointElement, LinearScale, CategoryScale, Filler, Tooltip, Legend)


const vResize = resizeDirective

const props = defineProps({
  loading: {
    default: false,
    type: Boolean,
  },
  dark: {
    default: false,
    type: Boolean,
  },
  data: {
    required: false,
    type: Array,
    default: [
      {
        data: '2024-01-01',
        in: 1024,
        out: 2048,
      },
      {
        data: '2024-01-02',
        in: 2048,
        out: 1024,
      },
      {
        data: '2024-01-03',
        in: 1024,
        out: 2048,
      },
      {
        data: '2024-01-04',
        in: 2048,
        out: 1024,
      },
    ]
  }
})

const wrapperSize = ref()
const wrapperResize = (size) => {
  wrapperSize.value = size
}

const chartRef = ref()
let chartInstance = null

const renderChart = () => {
  const ctx = chartRef.value.getContext('2d')

  const labels = props.data.map(d => d.date)
  const inData = props.data.map(d => (d.in / (1024 * 1024)).toFixed(2))
  const outData = props.data.map(d => (d.out / (1024 * 1024)).toFixed(2))

  // 销毁旧实例，避免重复渲染
  if (chartInstance) chartInstance.destroy()

  chartInstance = new Chart(ctx, {
    type: 'line',
    data: {
      labels,
      datasets: [
        {
          label: '上行流量',
          data: inData,
          borderColor: '#db9145',
          backgroundColor: 'rgba(219,145,69,0.3)',
          fill: true,
          tension: 0.4
        },
        {
          label: '下行流量',
          data: outData,
          borderColor: '#a45bd4',
          backgroundColor: 'rgba(164,91,212,0.3)',
          fill: true,
          tension: 0.4
        }
      ]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          labels: {
            color: props.dark ? '#ddd' : '#333'
          }
        },
        tooltip: {
          backgroundColor: props.dark ? '#333' : '#fff',
          titleColor: props.dark ? '#fff' : '#000',
          bodyColor: props.dark ? '#eee' : '#333',
        }
      },
      scales: {
        x: {
          ticks: {color: props.dark ? '#ccc' : '#333'},
          grid: {color: props.dark ? '#444' : '#eee'}
        },
        y: {
          ticks: {
            color: props.dark ? '#ccc' : '#333',
            callback: val => `${val} MB`
          },
          grid: {color: props.dark ? '#444' : '#eee'}
        }
      },
      animation: {duration: 500}
    }
  })
}

onMounted(() => renderChart())
onBeforeUnmount(() => chartInstance?.destroy())

// 数据或主题切换时自动更新
watch(() => [props.data, props.dark], () => renderChart(), {deep: true})

const chartStyleComputed = computed(() => {
  let result = {}
  if (!wrapperSize.value?.width) {
    result = {
      height: '1px',
    }
  } else {
    result = {
      height: wrapperSize.value?.height + 'px',
      width: wrapperSize.value?.width + 'px',
    }
  }
  return result
})

onUnmounted(() => {
})
</script>

<template>
  <AppCard :show-border="false" :loading="props.loading" style="margin: 12px 0 !important;" v-resize="wrapperResize">
    <canvas ref="chartRef" :style="chartStyleComputed"></canvas>
  </AppCard>
</template>

<style scoped lang="scss">

</style>