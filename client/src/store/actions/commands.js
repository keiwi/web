import api from '../../api'

export default {
    showCommandModal: ({ commit }) => commit('showCommandModal'),
    hideCommandModal: ({ commit }) => commit('hideCommandModal'),
    createCommand: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            api.createCommand(payload)
                .then((response) => {
                    payload.id = response.data.id
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
                            id: c.id,
                            command: c.command,
                            name: c.name,
                            description: c.description,
                            format: c.format
                        })
                    }
                    resolve()
                }, (response) => {
                    reject(response)
                })
        })
    },
    deleteCommand: ({ commit }, id) => {
        return new Promise((resolve, reject) => {
            api.deleteCommand(id)
                .then((response) => {
                    commit('deleteCommand', id)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    }
}
