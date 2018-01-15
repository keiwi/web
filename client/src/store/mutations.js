import GROUPS from './mutations/groups'
import COMMANDS from './mutations/commands'
import CLIENTS from './mutations/clients'

// mutations are operations that modify the state in some way.
// each mutation handler gets the entire state tree as the
// first argument, followed by additional payload arguments.
// mutations must be synchronous and can be recorded by plugins
// for debugging purposes.
export const mutations = {
    ...GROUPS,
    ...COMMANDS,
    ...CLIENTS
}
