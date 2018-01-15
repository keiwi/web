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
    createGroup (state, name) {
        state.groups.push({ Commands: [], Name: name })
    },
    editGroup (state, {id, option, value}) {
        for (var i in state.groups) {
            for (var j in state.groups[i].Commands) {
                if (state.groups[i].Commands[j].ID === id) {
                    state.groups[i].Commands[j][option] = value
                    return
                }
            }
        }
    },
    addGroupCommand (state, payload) {
        for (var i in state.groups) {
            if (state.groups[i].Name === payload.group_name) {
                state.groups[i].Commands.push({
                    CommandID: payload.command_id,
                    ID: payload.ID,
                    NextCheck: payload.next_check,
                    StopError: payload.stop_error
                })
                return
            }
        }
        state.groups.push({
            Name: payload.group_name,
            Commands: [{
                CommandID: payload.command_id,
                ID: payload.ID,
                NextCheck: payload.next_check,
                StopError: payload.stop_error
            }]
        })
    },
    deleteGroupCommand (state, payload) {
        for (var i in state.groups) {
            for (var j in state.groups[i].Commands) {
                if (state.groups[i].Commands[j].ID === payload.id) {
                    state.groups[i].Commands.splice(j, 1)
                    return
                }
            }
        }
    },
    deleteGroup (state, payload) {
        for (var i in state.groups) {
            if (state.groups[i].Name === payload.name) {
                state.groups.splice(i, 1)
                return
            }
        }
    }
}
