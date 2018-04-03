import Vue from 'vue'
import utils from './utils.js'

export default {
    // /api/v1/groups/exists
    groupExists (name) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.EXISTS_GROUP_URL, {name}).then(response => {
                if (!response.body.success === false) reject(response.body.message)
                resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/groups/create
    addGroupCommand (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.CREATE_GROUPS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/groups/get
    getGroups () {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.GET_GROUPS_URL).then(response => {
                if (response.body == null) resolve([])
                else if (!response.body.success === false) reject(response.body.message)
                else resolve(response.body)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/groups/edit
    editGroup (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.EDIT_GROUPS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/groups/delete/id
    deleteGroupCommand (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.DELETE_ID_GROUPS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    },
    // /api/v1/groups/delete/name
    deleteGroup (payload) {
        return new Promise((resolve, reject) => {
            Vue.http.post(utils.DELETE_NAME_GROUPS_URL, payload).then(response => {
                if (response.body.success) resolve(response.body)
                else reject(response.body.message)
            }, response => {
                reject(response.statusText)
            })
        })
    }
}
