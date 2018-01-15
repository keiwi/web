import orderBy from 'lodash/orderBy'

const fields = ['limit', 'offset', 'sort', 'order']

export default function (values, query) {
    const { limit = 10, offset = 0, sort = '', order = '' } = query

    for (let key in query) {
        if (fields.includes(key)) continue
        values = values.filter(row => row[key].toLowerCase().includes(query[key].toLowerCase()))
    }

    if (sort) {
        values = orderBy(values, sort, order)
    }

    const res = {
        rows: values.slice(offset, offset + limit),
        total: values.length
    }

    return res
}
