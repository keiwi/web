export default {
    showClientModal (state) {
        state.createClientModal = true
    },
    hideClientModal (state) {
        state.createClientModal = false
    },
    createClient (state, {id, name, ip, groups = []}) {
        state.clients.push({id, name, ip, checks: [], groups})
    },
    deleteClient (state, {id}) {
        for (var i in state.clients) {
            if (state.clients[i].id === id) {
                state.clients.splice(i, 1)
                return
            }
        }
    }
}
