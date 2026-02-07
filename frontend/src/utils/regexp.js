export function regexpLocalIp(value) {
    return true
    // 验证10.0.0.0 - 10.255.255.255
    if (/^10\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})$/.test(value)) {
        return true
    }
    // 验证172.16.0.0 - 172.31.255.255
    if (/^172\.(1[6-9]|2\d|3[01])\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})$/.test(value)) {
        return true
    }
    // 验证192.168.0.0 - 192.168.255.255
    if (/^192\.168\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})$/.test(value)) {
        return true
    }
    // 验证127.0.0.0 - 127.255.255.255
    return /^127\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})\.(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})$/.test(value)
}

export function regexpPort(value) {
    return /^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$/.test(value)
}

export function regexpDomain(value) {
    return /^(?:[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?\.)+[A-Za-z]{2,63}$/.test(value)
}

export function regexpDomainPrefix(value) {
    return /^(?!-)[a-z0-9-]{1,63}(?<!-)$/.test(value);
}
