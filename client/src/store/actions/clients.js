import api from '../../api'

export default {
    showClientModal: ({ commit }) => commit('showClientModal'),
    hideClientModal: ({ commit }) => commit('hideClientModal'),
    createClient: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            api.createClient(payload)
                .then((response) => {
                    payload.ID = response.data.id
                    commit('createClient', payload)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    },
    initClients: ({ commit }) => {
        return new Promise((resolve, reject) => {
            api.getAllClients()
                .then((response) => {
                    for (let c of response) {
                        commit('createClient', {
                            ID: c.id,
                            namn: c.namn,
                            ip: c.ip,
                            Groups: c.group_names.split(',')
                        })
                    }
                    resolve()
                }, (response) => {
                    reject(response)
                })
        })
    },
    deleteClient: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            api.deleteClient(payload)
                .then((response) => {
                    commit('deleteClient', payload)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    }
}
