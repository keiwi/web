import GROUPS from './actions/groups'
import COMMANDS from './actions/commands'
import CLIENTS from './actions/clients'

export const actions = {
    ...GROUPS,
    ...COMMANDS,
    ...CLIENTS
}
