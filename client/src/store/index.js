import Vue from 'vue'
import Vuex from 'vuex'

import {mutations} from './mutations'
import {actions} from './actions'
import {getters} from './getters'

Vue.use(Vuex)

// root state object.
// each Vuex instance is just a single state tree.
const state = {
    createGroupModal: false,
    createCommandModal: false,
    addGroupCommandModal: false,
    createClientModal: false,
    activeGroupCommand: '',
    groupDisplayRow: {},
    commands: [],
    groups: [],
    clients: []
}

// A Vuex instance is created by combining the state, mutations, actions,
// and getters.
export default new Vuex.Store({
    state,
    getters,
    actions,
    mutations
})
