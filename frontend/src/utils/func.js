// 函数防抖
export function debounce(func, delay) {
    let timer = null
    return function () {
        clearTimeout(timer)
        timer = setTimeout(() => {
            func.apply(this, arguments)
        }, delay)
    }
}
// 第一次执行，之后防抖
export function debounceAdvanced(func, delay, immediate) {
    let timer
    return function () {
        if (timer) clearTimeout(timer)
        if (immediate) {
            // 复杂的防抖函数
            // 判断定时器是否为空，如果为空，则会直接执行回调函数
            let firstRun = !timer
            // 不管定时器是否为空，都会重新开启一个新的定时器,不断输入，不断开启新的定时器，当不在输入的delay后，再次输入就会立即执行回调函数
            timer = setTimeout(() => {
                timer = null
            }, delay)
            if (firstRun) {
                func.apply(this, arguments)
            }
            // 简单的防抖函数
        } else {
            timer = setTimeout(() => {
                func.apply(this, arguments)
            }, delay)
        }
    }
}