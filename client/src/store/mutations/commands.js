export default {
    showCommandModal (state) {
        state.createCommandModal = true
    },
    hideCommandModal (state) {
        state.createCommandModal = false
    },
    createCommand (state, {ID, Command, Namn, Description, Format}) {
        state.commands.push({ID, Command, Namn, Description, Format})
    },
    editCommand (state, {id, option, value}) {
        for (var i in state.commands) {
            if (state.commands[i].ID === id) {
                state.commands[i][option] = value
                return
            }
        }
    },
    deleteCommand (state, ID) {
        for (var i in state.commands) {
            if (state.commands[i].ID === ID) {
                state.commands.splice(i, 1)
                return
            }
        }
    }
}
