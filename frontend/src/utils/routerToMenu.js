export function routerToMenu(children) {
    let result = []
    children.forEach(item => {
        if (item.meta?.hidden === 1) {
            return
        }
        let data = {
            label: item.meta?.title,
            key: item.name,
            iconSvg: item.meta?.icon,
            show: item.meta?.hidden === 2,
            link: item.meta?.link,
        }
        if (item.children) {
            children = routerToMenu(item.children)
            if (children.length !== 0) {
                data.children = children
            } else {
                data.show = false
            }
        }
        result.push(data)
    })
    return result
}