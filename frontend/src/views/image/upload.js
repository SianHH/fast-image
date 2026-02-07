import {localStore} from "../../store/local.js";
import axios from "axios";

export const customRequest = async ({file, onStart, onEnd, onFinish, onError, onResult}) => {
    try {
        let split = file.name.split('.')
        if (!['jpg', 'jpeg', 'png', 'gif', 'webp'].includes(split[split.length - 1].toLowerCase())) {
            $message.create('只能上传jpg,jpeg,png,gif,webp格式图片', {
                type: "warning",
                duration: 1500,
                closable: true
            })
            onError()
            return
        }

        if (file.size > 100 * 1024 * 1024) {
            $message.create('文件大小超过100M', {
                type: "warning",
                duration: 1500,
                closable: true
            })
            onError()
            return
        }
        onStart(file)

        const fData = new FormData()
        fData.append('file', file)

        const {data} = await axios.postForm('/api/v1/image/upload', fData, {
            headers: {
                token: localStore().auth.token
            }
        })

        if (data.code === 0) {
            $message.create('文件上传成功', {
                type: "success", closable: true, duration: 1500
            })
            onFinish()
            onResult(data.data)
        } else {
            $message.create(data.msg || '文件上传失败', {
                type: "warning", closable: true, duration: 1500
            })
            onError()
        }
    } catch (err) {
        console.log(err)
        $message.create('网络错误', {
            type: "warning", closable: true, duration: 1500
        })
        onError()
    } finally {
        onEnd(file)
    }
}
