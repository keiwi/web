import Vue from 'vue'
import utils from './utils.js'

export default {
    // /api/v1/clients/create
    createClient (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.CREATE_CLIENTS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/clients/get/all
    getAllClients () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.GET_ALL_CLIENTS_URL).then(response => {
                resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/clients/get/id
    getIDClients (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.GET_ID_CLIENTS_URL, payload).then(response => {
                resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/clients/edit
    editClient (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.EDIT_CLIENTS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/clients/delete
    deleteClient (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.DELETE_CLIENTS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    }
}
