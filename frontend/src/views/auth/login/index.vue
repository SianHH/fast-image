<script setup>
import {onBeforeMount, ref} from "vue";
import {apiAuthCaptcha, apiAuthLogin} from "../../../api/auth/index.js";
import {localStore} from "../../../store/local.js";
import {requiredRule} from "../../../utils/formDataRule.js";

import {Lock, User} from '@vicons/fa'

const state = ref({
  captcha: {
    src: '',
    loading: false,
    security: true,
  },
  login: {
    loading: false,
    data: {
      account: '',
      password: '',
      captchaKey: '',
      captchaValue: ''
    },
    dataRules: {
      account: requiredRule('请输入账号'),
      password: requiredRule('请输入密码'),
      captchaValue: requiredRule('请输入验证码'),
    }
  },
})

const refreshCaptchaFunc = async () => {
  try {
    state.value.captcha.loading = true
    let res = await apiAuthCaptcha()
    state.value.login.data.captchaKey = res.data.key
    state.value.captcha.src = res.data?.bs64
    state.value.captcha.security = res.data?.security
  } finally {
    state.value.captcha.loading = false
  }
}

const loginRef = ref()
const loginFunc = () => {
  loginRef.value.validate(async (errors) => {
    if (!errors) {
      try {
        state.value.login.loading = true
        let res = await apiAuthLogin(state.value.login.data)
        localStore().auth.token = res.data
        window.location.reload()
      } catch (err) {
        await refreshCaptchaFunc()
      } finally {
        state.value.login.loading = false
      }
    }
  });
}

onBeforeMount(() => {
  refreshCaptchaFunc()
})

</script>

<template>
  <div class="wrapper">
    <n-card class="bg">
      <div class="title" style="user-select: none;cursor: pointer;font-size: 1.5em">FAST IMAGE</div>
      <n-divider/>
      <n-form
          ref="loginRef"
          :model="state.login.data"
          :rules="state.login.dataRules"
          :show-label="false"
          size="large"
      >
        <n-form-item path="account" label="账号">
          <n-input v-model:value="state.login.data.account" placeholder="账号">
            <template #prefix>
              <n-icon>
                <User/>
              </n-icon>
            </template>
          </n-input>
        </n-form-item>
        <n-form-item path="password" label="密码">
          <n-input
              v-model:value="state.login.data.password"
              type="password"
              placeholder="密码"
              show-password-on="click"
          >
            <template #prefix>
              <n-icon>
                <Lock/>
              </n-icon>
            </template>
          </n-input>
        </n-form-item>
        <n-form-item path="captchaValue" label="验证码" v-if="!state.captcha.security">
          <n-input
              v-model:value="state.login.data.captchaValue"
              placeholder="验证码"
          />
          <n-tooltip trigger="hover">
            <template #trigger>
              <n-spin :show="state.captcha.loading" size="small" :delay="200" style="height: 40px">
                <n-image
                    height="40"
                    @click="refreshCaptchaFunc"
                    style="border-radius: 3px;margin-left: 8px"
                    :src="'data:image/jpeg;base64,' + state.captcha.src"
                    preview-disabled
                />
              </n-spin>
            </template>
            点击刷新
          </n-tooltip>
        </n-form-item>
        <n-button style="width: 100%" type="primary" @click="loginFunc" :loading="state.login.loading">
          登录
        </n-button>
      </n-form>
    </n-card>
  </div>
</template>

<style lang="scss" scoped>
:deep(.n-card.n-card--bordered) {
  border: 0 !important;
}

:deep(.n-card__content) {
  padding: 0 !important;
}

:deep(.n-divider:not(.n-divider--vertical)) {
  margin-top: 10px;
  margin-bottom: 10px;
}

.wrapper {
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #303133;


  .bg {
    border-radius: 8px;
    overflow: hidden;
    padding: 16px;
  }
}

@media screen and (min-width: 700px) {
  .bg {
    width: 320px;
  }
}

@media screen and (max-width: 700px) {
  .bg {
    width: 320px;
  }
}

@media screen and (max-width: 500px) {
  .bg {
    width: 80%;
  }
}
</style>