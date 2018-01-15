import Commands from './commands.js'
import Groups from './groups.js'
import Clients from './clients.js'
import Checks from './checks.js'

export default {
    ...Commands,
    ...Groups,
    ...Clients,
    ...Checks
}
