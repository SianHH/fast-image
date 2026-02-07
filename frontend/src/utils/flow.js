export function flowFormat(value) {
    if (value < 0) {
        return '--'
    }
    if (value < 1024) {
        return value + 'B'
    }
    if (value < 1024 * 1024) {
        return (value / 1024).toFixed(2) + 'KB'
    }
    if (value < 1024 * 1024 * 1024) {
        return (value / 1024 / 1024).toFixed(2) + 'MB'
    }
    if (value < 1024 * 1024 * 1024 * 1024) {
        return (value / 1024 / 1024 / 1024).toFixed(2) + 'GB'
    }
    if (value < 1024 * 1024 * 1024 * 1024 * 1024) {
        return (value / 1024 / 1024 / 1024 / 1024).toFixed(2) + 'TB'
    }
    return '--'
}