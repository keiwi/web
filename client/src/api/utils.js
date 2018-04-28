// Main URL
const API_URL = '/api'

// Commands URL
const COMMANDS_URL = API_URL + '/commands'
const CREATE_COMMANDS_URL = COMMANDS_URL + '/create'
const DELETE_COMMANDS_URL = COMMANDS_URL + '/delete'
const EDIT_COMMANDS_URL = COMMANDS_URL + '/edit'
const GET_COMMANDS_URL = COMMANDS_URL + '/get'

// Groups URL
const GROUPS_URL = API_URL + '/groups'
const CREATE_GROUPS_URL = GROUPS_URL + '/create'
const DELETE_ID_GROUPS_URL = GROUPS_URL + '/delete/id'
const DELETE_NAME_GROUPS_URL = GROUPS_URL + '/delete/name'
const EDIT_GROUPS_URL = GROUPS_URL + '/edit'
const GET_GROUPS_URL = GROUPS_URL + '/get'
const EXISTS_GROUP_URL = GROUPS_URL + '/exists'

// Clients URL
const CLIENTS_URL = API_URL + '/clients'
const CREATE_CLIENTS_URL = CLIENTS_URL + '/create'
const DELETE_CLIENTS_URL = CLIENTS_URL + '/delete'
const EDIT_CLIENTS_URL = CLIENTS_URL + '/edit'
const GET_ALL_CLIENTS_URL = CLIENTS_URL + '/get/all'
const GET_ID_CLIENTS_URL = CLIENTS_URL + '/get/id'

// Checks URL
const CHECKS_URL = API_URL + '/checks'
const DELETE_CHECKS_URL = CHECKS_URL + '/delete'
const GET_ALL_CHECKS_URL = CHECKS_URL + '/get/all'
const GET_ID_CHECKS_URL = CHECKS_URL + '/get/id'
const GET_CLIENT_CMD_ID_CHECKS_URL = CHECKS_URL + '/get/client-cmd'
const GET_CHECKS_DATE_CLIENT_URL = CHECKS_URL + '/get/checks-date-client'

export default {
    // Main URL
    API_URL,
    // Commands URL
    COMMANDS_URL,
    CREATE_COMMANDS_URL,
    DELETE_COMMANDS_URL,
    EDIT_COMMANDS_URL,
    GET_COMMANDS_URL,
    // Groups URL
    GROUPS_URL,
    CREATE_GROUPS_URL,
    DELETE_ID_GROUPS_URL,
    DELETE_NAME_GROUPS_URL,
    EDIT_GROUPS_URL,
    GET_GROUPS_URL,
    EXISTS_GROUP_URL,
    // Clients URL
    CLIENTS_URL,
    CREATE_CLIENTS_URL,
    DELETE_CLIENTS_URL,
    EDIT_CLIENTS_URL,
    GET_ALL_CLIENTS_URL,
    GET_ID_CLIENTS_URL,
    // Checks URL
    CHECKS_URL,
    DELETE_CHECKS_URL,
    GET_ALL_CHECKS_URL,
    GET_ID_CHECKS_URL,
    GET_CLIENT_CMD_ID_CHECKS_URL,
    GET_CHECKS_DATE_CLIENT_URL
}
