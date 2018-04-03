export default {
    showGroupModal (state) {
        state.createGroupModal = true
    },
    hideGroupModal (state) {
        state.createGroupModal = false
    },
    showGroupCommandModal (state) {
        state.addGroupCommandModal = true
    },
    hideGroupCommandModal (state) {
        state.addGroupCommandModal = false
    },
    setActiveGroupCommand (state, name) {
        state.activeGroupCommand = name
    },
    toggleGroupDisplay (state, ID) {
        state.groupDisplayRow[ID] = !state.groupDisplayRow[ID]
    },
    createGroup (state, group) {
        state.groups.push(group)
    },
    editGroup (state, {id, option, value}) {
        for (var i in state.groups) {
            for (var j in state.groups[i].commands) {
                if (state.groups[i].commands[j].id === id) {
                    state.groups[i].commands[j][option] = value
                    return
                }
            }
        }
    },
    addGroupCommand (state, payload) {
        const count = ids =>
            ids.reduce((a, b) =>
                Object.assign(a, {[b]: (a[b] || 0) + 1}), {})

        const nonduplicates = dict =>
            Object.keys(dict).filter((a) => dict[a] <= 1)

        for (var i in state.groups) {
            if (state.groups[i].name === payload.name) {
                var ids = payload.commands.concat(state.groups[i].commands)
                    .map((command) => command.id)

                var nondupIds = nonduplicates(count(ids))

                for (var j in payload.commands) {
                    for (var id of nondupIds) {
                        if (payload.commands[j].id === id) {
                            state.groups[i].commands.push(payload.commands[j])
                        }
                    }
                }
                return
            }
        }
        state.groups.push(payload)
    },
    deleteGroupCommand (state, payload) {
        for (var i in state.groups) {
            for (var j in state.groups[i].commands) {
                if (state.groups[i].commands[j].id === payload.id) {
                    state.groups[i].commands.splice(j, 1)
                    return
                }
            }
        }
    },
    deleteGroup (state, payload) {
        for (var i in state.groups) {
            if (state.groups[i].name === payload.name) {
                state.groups.splice(i, 1)
                return
            }
        }
    }
}
