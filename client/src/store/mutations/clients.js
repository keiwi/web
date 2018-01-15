export default {
    showClientModal (state) {
        state.createClientModal = true
    },
    hideClientModal (state) {
        state.createClientModal = false
    },
    createClient (state, {ID, namn, ip, Groups = []}) {
        state.clients.push({ID, Name: namn, IP: ip, Checks: [], Groups})
    },
    deleteClient (state, {id}) {
        for (var i in state.clients) {
            if (state.clients[i].ID === id) {
                state.clients.splice(i, 1)
                return
            }
        }
    }
}
