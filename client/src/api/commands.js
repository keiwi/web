import Vue from 'vue'
import utils from './utils.js'

export default {
    // /api/v1/commands/create
    createCommand (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.CREATE_COMMANDS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/commands/get
    getCommand () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.GET_COMMANDS_URL).then(response => {
                resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/commands/edit
    editCommand (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.EDIT_COMMANDS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/commands/delete
    deleteCommand (id) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.DELETE_COMMANDS_URL, {id}).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    }
}
