import {createApp} from 'vue'
import './style.css'
import App from './App.vue'
import {createPinia} from "pinia";
import piniaPluginPersistence from 'pinia-plugin-persistedstate'
import router from "./router/index.js";
import naive from 'naive-ui'
import {setupStore} from "./store/index.js";
import "flag-icons/css/flag-icons.min.css";

const pinia = createPinia()
pinia.use(piniaPluginPersistence)

const app = createApp(App)
setupStore(app)
app.use(naive)
app.use(router)
app.mount('#app')
