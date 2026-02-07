import axios from 'axios'
import {localStore} from "../store/local.js";

const request = axios.create({
    baseURL: "/api",
    timeout: 60000,
    headers: {}
})

request.interceptors.request.use(config => {
    config.headers['Token'] = localStore().auth.token
    return config
}, (err) => {
    return Promise.reject(err)
})

request.interceptors.response.use((res) => {
    switch (res.data.code) {
        case 0: // 成功
            return res.data
        case 1: // 错误
            $message.create(res.data.msg, {type: "error", closable: true, duration: 3000,})
            return Promise.reject(new Error(res.data.msg))
        case 2: // 登录失效
            $message.create(res.data.msg, {type: "error", closable: true, duration: 3000,})
            localStore().auth.token = ""
            setTimeout(() => {
                window.location.reload()
            }, 1500)
            return Promise.reject(new Error(res.data.msg))
        case 3: // 未登录
            $message.create(res.data.msg, {type: "error", closable: true, duration: 3000,})
            localStore().auth.token = ""
            setTimeout(() => {
                window.location.reload()
            }, 1500)
            return Promise.reject(new Error(res.data.msg))
        case 4: // 未授权
            $message.create(res.data?.msg || '未知错误', {type: "error", closable: true, duration: 3000,})
            return Promise.reject(new Error(res.data?.msg || '未知错误'))
        default:
            $message.create(res.data?.msg || '未知错误', {type: "error", closable: true, duration: 3000,})
            return Promise.reject(new Error(res.data?.msg || '未知错误'))
    }
}, (err) => {
    $message.create( '网络错误', {type: "error", closable: true, duration: 3000,})
    return Promise.reject(err)
})

export default request