export function copyToClipboard(value) {
    if (!navigator) {
        return Promise.reject(new Error('当前环境不支持复制到剪切板'))
    }
    return navigator.clipboard && navigator.clipboard.writeText(value)
}

export async function copyToClipboardDirect(value) {
    if (!navigator) {
        $message.create('当前环境不支持复制到剪切板', {
            type: "warning",
            closable: true,
            duration: 1500
        })
        return
    }
    if (!navigator.clipboard) {
        $message.create('当前环境不支持复制到剪切板', {
            type: "warning",
            closable: true,
            duration: 1500
        })
        return
    }
    try {
        await navigator.clipboard.writeText(value)
        $message.create('复制成功', {
            type: "success",
            closable: true,
            duration: 1500
        })
    } catch (err) {
        $message.create('复制失败' + err, {
            type: "warning",
            closable: true,
            duration: 1500
        })
    }
}