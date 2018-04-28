import Vue from 'vue'
import utils from './utils.js'
import moment from 'moment'

export default {
    // /api/v1/checks/get/all
    getAllChecks () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.GET_ALL_CHECKS_URL).then(response => {
                if (response.body.success) {
                    resolve(response.body.data)
                } else {
                    reject(response.body.message)
                }
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/checks/get/id
    getIDChecks (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.GET_ID_CHECKS_URL, payload).then(response => {
                if (response.body.success) {
                    resolve(response.body.data)
                } else {
                    reject(response.body.message)
                }
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/checks/get/client-cmd
    getClientCMDIDChecks (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.GET_CLIENT_CMD_ID_CHECKS_URL, payload).then(response => {
                if (response.body.success) {
                    resolve(response.body.data)
                } else {
                    reject(response.body.message)
                }
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/checks/get/checks-date-client
    getChecksDateClient (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.GET_CHECKS_DATE_CLIENT_URL, payload).then(response => {
                if (response.body.success) {
                    for (var i in response.body.data) {
                        response.body.data[i] = fixUTCDate(response.body.data[i])
                    }
                    resolve(response.body.data)
                } else {
                    reject(response.body.message)
                }
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/checks/delete
    deleteCheck (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.DELETE_CHECKS_URL, payload).then(response => {
                if (response.body.success) {
                    resolve(response.body)
                } else {
                    reject(response.body.message)
                }
            }, response => {
                reject(response.statusText)
            })
        })
    }
    // /api/v1/che
}

function fixUTCDate (d) {
    if (typeof d === 'undefined' || d == null) return d
    if (typeof d.created_at !== 'undefined') d.created_at = moment(d.created_at).subtract(4, 'hours').format()
    if (typeof d.updated_at !== 'undefined') d.updated_at = moment(d.created_at).subtract(4, 'hours').format()
    return d
}
