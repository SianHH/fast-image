import request from "../../request/index.js"

const baseUrl = '/v1/image'

export const apiImageList = (data = {}) => {
    return request.request({
        url: `${baseUrl}/list`,
        method: 'POST',
        data
    })
}

export const apiImageDelete = (data = {}) => {
    return request.request({
        url: `${baseUrl}/delete`,
        method: 'POST',
        data
    })
}
