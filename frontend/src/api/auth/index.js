import request from "../../request/index.js"

const baseUrl = '/v1/auth'

// 登录
export const apiAuthLogin = (data) => {
    return request.request({
        url: `${baseUrl}/login`,
        method: 'POST',
        data
    })
}

// 获取验证码
export const apiAuthCaptcha = () => {
    return request.request({
        url: `${baseUrl}/captcha?timestamp=${new Date().getTime()}`,
        method: 'POST'
    })
}
