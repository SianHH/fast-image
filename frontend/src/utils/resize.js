export const resizeDirective = {
    mounted(el, binding) {
        const callback = binding.value
        const delay = binding.arg || 100 // 可通过指令参数设置防抖延迟，默认100ms

        // 防抖函数
        let timer = null
        const debouncedCallback = (rect) => {
            if (timer) clearTimeout(timer)
            timer = setTimeout(() => {
                callback(rect)
            }, delay)
        }

        const observer = new ResizeObserver(entries => {
            entries.forEach(entry => {
                debouncedCallback(entry.contentRect)
            })
        })

        observer.observe(el)
        el._resizeObserver = observer
    },
    unmounted(el) {
        if (el._resizeObserver) {
            el._resizeObserver.disconnect()
        }
    }
}
