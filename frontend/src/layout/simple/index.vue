<script setup>
import Header from "../components/Header.vue";
import {computed} from "vue";
import {appStore} from "../../store/app.js";


// 是否为侧滑菜单
const layoutDirectionIsSideComputed = computed(() => {
  return appStore().width >= 768
})

</script>

<template>
  <n-layout class="layout-container">
    <n-layout-header>
      <Header/>
    </n-layout-header>

    <n-layout-content has-sider class="layout-body">
      <n-scrollbar style="max-height: calc(100vh - 60px)">
        <slot></slot>
      </n-scrollbar>
    </n-layout-content>
  </n-layout>
</template>

<style lang="scss" scoped>
.layout-container {
  height: 100vh;
  background-color: black;

  .layout-body {
    height: calc(100vh - 60px);
    margin: 0 auto;
  }
}
</style>