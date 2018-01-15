import api from '../../api'

export default {
    showGroupModal: ({ commit }) => commit('showGroupModal'),
    hideGroupModal: ({ commit }) => commit('hideGroupModal'),
    showGroupCommandModal: ({ commit }, name) => {
        commit('setActiveGroupCommand', name)
        commit('showGroupCommandModal')
    },
    hideGroupCommandModal: ({ commit }) => commit('hideGroupCommandModal'),
    toggleGroupDisplay: ({ commit }, payload) => commit('toggleGroupDisplay', payload),
    createGroup: ({ commit }, name) => {
        return new Promise((resolve, reject) => {
            api.groupExists(name)
                .then((response) => {
                    if (!response) commit('createGroup', name)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    },
    editGroup: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            console.log(payload)
            api.editGroup(payload)
                .then((response) => {
                    commit('editGroup', payload)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    },
    addGroupCommand: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            api.addGroupCommand(payload)
                .then((response) => {
                    commit('addGroupCommand', response.data)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    },
    initGroups: ({ commit }) => {
        return new Promise((resolve, reject) => {
            api.getGroups()
                .then((response) => {
                    for (let c of response) {
                        commit('addGroupCommand', c)
                    }
                    resolve()
                }, (response) => {
                    reject(response)
                })
        })
    },
    deleteGroupCommand: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            api.deleteGroupCommand(payload)
                .then((response) => {
                    commit('deleteGroupCommand', payload)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    },
    deleteGroup: ({ commit }, payload) => {
        return new Promise((resolve, reject) => {
            api.deleteGroup(payload)
                .then((response) => {
                    commit('deleteGroup', payload)
                    resolve(response)
                }, (response) => {
                    reject(response)
                })
        })
    }
}
