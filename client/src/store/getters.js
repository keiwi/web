// getters are functions
export const getters = {
    getCommand: (state, getters) => (ID) => state.commands.find(cmd => cmd.id === ID),
    getGroup: (state, getters) => (ID) => state.groups.find(group => group.id === ID),
    getGroupCommand: (state, getters) => (ID) => {
        for (let group of state.groups) {
            for (let cmd of group.commands) {
                if (cmd.id === ID) return cmd
            }
        }
    },
    getGroupDisplay: (state, getters) => (ID) => {
        return state.groupDisplayRow[ID]
    }
}
