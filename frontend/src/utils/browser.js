export function goToUrl(url, target = '_blank') {
    if (url.startsWith('http')) {
        window.open(url, target)
    } else if (url.startsWith('https')) {
        window.open('https://' + url, target)
    } else {
        window.open(url, target)
    }
}