<template>
  <div
      v-html="svgContent"
      :style="{
      width: size + 'px',
      height: size + 'px',
      color: color,
      display: 'inline-block'
    }"
      :title="title"
      class="simple-icon"
  />
</template>

<script>
import * as icons from 'simple-icons/icons'

export default {
  name: 'SimpleIcon',
  props: {
    name: {
      type: String,
      required: true,
      validator: value => {
        // 转换为 simple-icons 的命名格式
        const iconName = 'si' + value.charAt(0).toUpperCase() + value.slice(1)
        return iconName in icons
      }
    },
    size: {
      type: Number,
      default: 24
    },
    color: {
      type: String,
      default: null
    },
    title: {
      type: String,
      default: null
    }
  },
  computed: {
    svgContent() {
      const iconName = 'si' + this.name.charAt(0).toUpperCase() + this.name.slice(1)
      const icon = icons[iconName]

      if (!icon) {
        console.warn(`Icon "${this.name}" not found in simple-icons`)
        return ''
      }

      let svg = icon.svg

      // 如果提供了自定义颜色，替换 fill 属性
      if (this.color) {
        svg = svg.replace(/fill="[^"]*"/, `fill="${this.color}"`)
      }

      return svg
    }
  }
}
</script>