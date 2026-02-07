<script setup>
import {onBeforeMount, ref} from "vue";
import moment from "moment";

const props = defineProps({
  size: {
    default: 'medium',
    type: String,
    required: false,
    validator: (value) => {
      return ['tiny', 'small', 'medium', 'large'].includes(value);
    }
  },
  defaultValue: {
    type: Date,
    required: false,
    default: new Date(),
  },
  maxValue: {
    type: Date,
    required: false,
    default: null,
  },
  minValue: {
    type: Date,
    required: false,
    default: null,
  }
})

const emits = defineEmits(['on-before', 'on-after'])

const now = ref()
const maxValue = ref('')
const minValue = ref('')

onBeforeMount(() => {
  now.value = props.defaultValue
  if (props.maxValue) {
    maxValue.value = moment(props.maxValue).format('yyyy-MM-DD')
  }
  if (props.minValue) {
    minValue.value = moment(props.minValue).format('yyyy-MM-DD')
  }
})

const onBefore = () => {
  now.value = moment(now.value).add(-1, 'days')
  emits('on-before', now.value)
}

const onAfter = () => {
  now.value = moment(now.value).add(1, 'days')
  emits('on-after', now.value)
}
</script>

<template>
  <n-input-group size="small">
    <n-button :size="props.size" @click="onBefore" :disabled="moment(now).format('yyyy-MM-DD')<=minValue">前一天
    </n-button>
    <n-input-group-label :size="props.size">{{ moment(now).format('yyyy-MM-DD') }}</n-input-group-label>
    <n-button :size="props.size" @click="onAfter" :disabled="moment(now).format('yyyy-MM-DD')>=maxValue">后一天
    </n-button>
  </n-input-group>
</template>

<style>

</style>