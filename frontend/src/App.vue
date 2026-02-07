<script setup>
import {computed, defineAsyncComponent, markRaw, onMounted, onUnmounted} from "vue";
import {darkTheme, dateZhCN, zhCN} from 'naive-ui';
import router from "./router/index.js";
import MessageApi from "./components/MessageApi.vue";
import DialogApi from "./components/DialogApi.vue";
import ModalApi from "./components/ModalApi.vue";
import {naiveThemeOverrides} from "./setting.js";
import {localStore} from "./store/local.js";
import {appStore} from "./store/app.js";
import NotifyApi from "./components/NotifyApi.vue";

const layouts = new Map()

function getLayout(name) {
  if (layouts.get(name)) return layouts.get(name)
  const layout = markRaw(defineAsyncComponent(() => import(`./layout/${name}/index.vue`)))
  layouts.set(name, layout)
  return layouts.get(name)
}

const Layout = computed(() => {
  if (!router.currentRoute.value.matched.length) return null
  return getLayout(router.currentRoute.value.meta?.layout || localStore().layout)
})

let resizeTimer = null
const resizeEventFunc = () => {
  if (resizeTimer) clearTimeout(resizeTimer)
  resizeTimer = setTimeout(() => {
    appStore().width = window.innerWidth
    appStore().height = window.innerHeight
  }, 100)
}

onMounted(async () => {
  window.addEventListener('resize', resizeEventFunc)
})
onUnmounted(() => {
  window.removeEventListener('resize', resizeEventFunc)
})
</script>

<template>
  <n-config-provider
      :locale="zhCN"
      :date-locale="dateZhCN"
      :theme="localStore().darkTheme ? darkTheme : null"
      :theme-overrides="naiveThemeOverrides"
  >
    <n-message-provider>
      <MessageApi/>
    </n-message-provider>
    <n-modal-provider>
      <ModalApi/>
    </n-modal-provider>
    <n-dialog-provider>
      <DialogApi/>
    </n-dialog-provider>
    <n-notification-provider>
      <NotifyApi/>
    </n-notification-provider>
    <router-view v-if="Layout" v-slot="{ Component, route: curRoute }">
      <component :is="Layout">
        <transition name="fade-slide" mode="out-in" appear>
          <component :is="Component" :key="curRoute.fullPath"/>
        </transition>
      </component>
    </router-view>
  </n-config-provider>
</template>

<style scoped>

</style>
