<script setup>

import {computed, onBeforeMount, onBeforeUpdate, ref, watch} from "vue";

const emits = defineEmits(['onChange'])

const props = defineProps({
  type: {
    default: 'input',
    required: true,
    type: String,
  },
  label: {
    default: '搜索项',
    required: false,
    type: String
  },
  labelWidth: {
    default: 100,
    required: false,
    type: Number
  },
  default: {
    default: null,
    required: false,
    type: [String, Number, Array],
  },
  size: {
    default: 'medium',
    required: false,
    type: String,
  },
  customStyle: {
    default: {},
    type: Object,
  },
  width: {
    default: 200,
    required: false,
    type: Number
  },
  placeholder: {
    default: null,
    required: false,
    type: String
  },
  clearable: {
    default: false,
    required: false,
    type: Boolean
  },
  options: {
    default: [],
    required: false,
    type: Array
  },
  optionsKeyValue: {
    default: {
      key: 'key',
      value: 'value'
    },
    required: false,
    type: Object,
  },
  inputNumberMax: {
    default: null,
    required: false,
    type: Number
  },
  inputNumberMin: {
    default: null,
    required: false,
    type: Number
  }
})

const result = ref(null)
const state = ref({
  size: 'medium',
})

onBeforeMount(() => {
  state.value.size = props.size
  result.value = props.default

})

onBeforeUpdate(() => {
  state.value.size = props.size
  result.value = props.default
})

const styleComputed = computed(() => {
  return Object.assign({
    width: props.width + 'px',
  }, props.customStyle)
})

watch(result, () => {
  emits('onChange', result.value)
})

</script>

<template>
  <div class="search-item">
    <div v-if="props.type==='custom'" class="body">
      <slot></slot>
    </div>
    <div v-else class="label" :style="{width:props.labelWidth+'px',textAlign:'left'}">
      {{ props.label + '：' }}
    </div>
    <div v-if="props.type!=='custom'" class="body">
      <n-scrollbar x-scrollable style="overflow-y: hidden">
        <!--  输入框  -->
        <n-input
            :style="styleComputed"
            v-if="props.type==='input'"
            v-model:value="result"
            :default-value="props.default"
            :placeholder="props.placeholder"
            :size="state.size"
            :clearable="props.clearable"
        ></n-input>

        <!--  数字输入框  -->
        <n-input-number
            :style="styleComputed"
            v-if="props.type==='input-number'"
            v-model:value="result"
            :default-value="props.default"
            :placeholder="props.placeholder"
            :size="state.size"
            :clearable="props.clearable"
            :max="props.inputNumberMax"
            :min="props.inputNumberMin"
        ></n-input-number>

        <!--  单选框  -->
        <n-radio-group
            v-else-if="props.type==='radio'"
            v-model:value="result"
            :default-value="props.default"
            :size="state.size"
            style="white-space:nowrap"
        >
          <n-radio
              :style="styleComputed"
              v-for="item in props.options"
              :value="item[props.optionsKeyValue.key]"
              :label="item[props.optionsKeyValue.value]"
          ></n-radio>
        </n-radio-group>

        <!--  单选按钮  -->
        <n-radio-group
            v-else-if="props.type==='radio-btn'"
            v-model:value="result"
            :default-value="props.default"
            :size="state.size"
        >
          <n-radio-button
              :style="styleComputed"
              v-for="item in props.options"
              :value="item[props.optionsKeyValue.key]"
              :label="item[props.optionsKeyValue.value]"
          ></n-radio-button>
        </n-radio-group>

        <!--  选择器  -->
        <n-select
            :style="styleComputed"
            v-else-if="props.type==='select'"
            v-model:value="result"
            :default-value="props.default"
            :size="state.size"
            :options="props.options"
            :label-field="props.optionsKeyValue.value"
            :value-field="props.optionsKeyValue.key"
            :clearable="props.clearable"
        >
        </n-select>
      </n-scrollbar>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.search-item {
  display: flex;
  align-items: center;
  margin: 4px 0;
  overflow: hidden;

  .label {
    font-size: 1.1em;
    font-weight: bold;
  }

  .body {
    flex: 1;
    overflow: hidden;

    :deep(.n-scrollbar-container) {
      overflow-y: hidden;
    }

    :deep(.n-scrollbar-rail--horizontal) {
      bottom: 0;
    }
  }
}
</style>