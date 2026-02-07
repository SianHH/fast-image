export function menuFlatToTree(data, parent) {
    let result = []
    if (!data) {
        return result
    }
    data.forEach(item => {
        if (item.parentCode === parent) {
            let children = menuFlatToTree(data, item.code)
            if (children.length === 0) {
                children = null
            }
            result.push({
                label: item.name,
                key: item.key,
                iconSvg: item.icon,
                show: item.hidden === 2,
                children: children
            })
        }
    })
    return result
}