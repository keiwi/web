import api from '../../api'

export default {
    createCommand: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            api.createCommand(payload)
                .then((response) => {
                    payload.ID = response.data.id
                    commit('createCommand', payload)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    },
    editCommand: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            api.editCommand(payload)
                .then((response) => {
                    commit('editCommand', payload)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    },
    initCommands: ({ commit }) => {
        return new Promise((resolve, reject) => {
            api.getCommand()
                .then((response) => {
                    for (let c of response) {
                        commit('createCommand', {
                            ID: c.id,
                            Command: c.command,
                            Namn: c.namn,
                            Description: c.description,
                            Format: c.format
                        })
                    }
                    resolve()
                }, (response) => {
                    reject(response)
                })
        })
    },
    deleteCommand: ({ commit }, ID) => {
        return new Promise((resolve, reject) => {
            api.deleteCommand(ID)
                .then((response) => {
                    commit('deleteCommand', ID)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    }
}
