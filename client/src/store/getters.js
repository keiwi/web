// getters are functions
export const getters = {
    getCommand: (state, getters) => (ID) => state.commands.find(cmd => cmd.ID === ID),
    getGroupCommand: (state, getters) => (ID) => {
        for (let group of state.groups) {
            for (let cmd of group.Commands) {
                if (cmd.ID === ID) return cmd
            }
        }
    },
    getGroupDisplay: (state, getters) => (ID) => {
        return state.groupDisplayRow[ID]
    }
}
