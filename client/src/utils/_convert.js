/**
 * fixedDecimals - Changes a number to have 2 decimals
 *
 * @param {Number} v - The number to change
 */
export const fixedDecimals = (v) => {
    return (Math.round((v + 0.00001) * 100) / 100)
}

/**
 * fixedUptime - Converts seconds to days/hours/minutes/seconds
 * @param {Integer} v - The amount of seconds to convert
 */
export const fixedUptime = (v) => {
    let out = ''
    var delta = v

    // calculate (and subtract) whole days
    var days = Math.floor(delta / 86400)
    delta -= days * 86400
    if (days === 1) out += days + ' day '
    if (days > 1) out += days + ' days '

    // calculate (and subtract) whole hours
    var hours = Math.floor(delta / 3600) % 24
    delta -= hours * 3600
    if (hours === 1) out += hours + ' hour '
    if (hours > 1) out += hours + ' hours '

    // calculate (and subtract) whole minutes
    var minutes = Math.floor(delta / 60) % 60
    delta -= minutes * 60
    if (minutes === 1) out += minutes + ' minute '
    if (minutes > 1) out += minutes + ' minutes '

    // what's left is seconds
    var seconds = Math.floor(delta % 60)
    if (seconds === 1) out += seconds + ' second'
    if (seconds > 1) out += seconds + ' seconds'

    return out
}

/**
 * Const for Byte/Kilo Byte/Mega Bytes/Giga Byte/Tera Byte
 */
export const B = 1 // 1
export const KB = B * 1024 // 1024^1
export const MB = KB * 1024 // 1024^2
export const GB = MB * 1024 // 1024^3
export const TB = GB * 1024 // 1024^4

/**
 * convertBytesBits - converts bytes/bits to a higher value like KB/MB/GB/TB
 * @param {*} v - The amount of bytes/bits to convert
 * @param {*} f - The prefix to convert to
 */
export const convertBytesBits = (v, f) => {
    switch (f) {
    case 'B':
    case 'b':
        return fixedDecimals(v / B)
    case 'KB':
    case 'Kb':
        return fixedDecimals(v / KB)
    case 'MB':
    case 'Mb':
        return fixedDecimals(v / MB)
    case 'GB':
    case 'Gb':
        return fixedDecimals(v / GB)
    case 'TB':
    case 'Tb':
        return fixedDecimals(v / TB)
    }
}

/**
 * getNormalPrefix - Will check which prefix is the most suitable to convert to, for bytes/bits
 * @param {*} v - The raw amount to check from
 */
export const getNormalPrefix = (v) => {
    if (v >= TB) return 'T'
    if (v >= GB) return 'G'
    if (v >= MB) return 'M'
    if (v >= KB) return 'K'
    return ''
}
