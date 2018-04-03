export default {
    showCommandModal (state) {
        state.createCommandModal = true
    },
    hideCommandModal (state) {
        state.createCommandModal = false
    },
    createCommand (state, {id, command, name, description, format}) {
        state.commands.push({id, command, name, description, format})
    },
    editCommand (state, {id, option, value}) {
        for (var i in state.commands) {
            if (state.commands[i].id === id) {
                state.commands[i][option] = value
                return
            }
        }
    },
    deleteCommand (state, ID) {
        for (var i in state.commands) {
            if (state.commands[i].id === ID) {
                state.commands.splice(i, 1)
                return
            }
        }
    }
}
