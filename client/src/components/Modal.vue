<template>
    <div class="modals">
        <b-modal
            title="Create a new client"
            size="lg"
            v-model="$store.state.createClientModal"
            @ok.prevent="createClientModal"
            ok-title="Create"
            style="margin-top:100px"><b-form>
                <b-form-group 
                    id="clientNameGroup"
                    label="Client Name:"
                    label-for="clientNameInput">
                    <b-form-input
                        id="clientNameInput"
                        type="text"
                        v-model="clientForm.name"
                        required
                        placeholder="Enter the client name"
                    ></b-form-input>
                </b-form-group>
                <b-form-group 
                    id="clientIPGroup"
                    label="Client IP:"
                    label-for="clientIPInput">
                    <b-form-input
                        id="clientIPInput"
                        type="text"
                        v-model="clientForm.ip"
                        required
                        placeholder="Enter the client IP"
                    ></b-form-input>
                </b-form-group>
            </b-form>
        </b-modal>
        <b-modal
            :title="`Add a new command to ${$store.state.activeGroupCommand}`"
            size="lg"
            v-model="$store.state.addGroupCommandModal"
            @ok.prevent="createGroupCommand"
            ok-title="Add"
            style="margin-top:100px"><b-form>
                <b-form-group 
                    id="groupCommandGroup"
                    label="Group Command:"
                    label-for="groupCommandSelect">
                    <b-form-select
                        id="groupCommandSelect"
                        v-model="groupCommandForm.command"
                        :options="commandOptions" />
                </b-form-group>
            </b-form>
        </b-modal>
        <b-modal
            title="Create a new group"
            size="lg"
            v-model="$store.state.createGroupModal"
            @ok.prevent="createGroupModal"
            ok-title="Create"
            style="margin-top:100px"><b-form>
                <b-form-group 
                    id="groupNameGroup"
                    label="Group Name:"
                    label-for="groupNameInput">
                    <b-form-input
                        id="groupNameInput"
                        type="text"
                        v-model="groupForm.name"
                        required
                        placeholder="Enter the group name"
                    ></b-form-input>
                </b-form-group>
            </b-form>
        </b-modal>
        <b-modal
            title="Create a new command"
            size="lg"
            v-model="$store.state.createCommandModal"
            @ok.prevent="sendCommandModal"
            ok-title="Create"
            style="margin-top:100px">
            <b-form>
                <b-form-group 
                    id="commandInputGroup"
                    label="Command:"
                    label-for="commandInput">
                    <b-form-input
                        id="commandInput"
                        type="text"
                        v-model="commandForm.command"
                        required
                        placeholder="Enter the command"
                    ></b-form-input>
                </b-form-group>
                <b-form-group 
                    id="nameInputGroup"
                    label="Name:"
                    label-for="nameInput">
                    <b-form-input
                        id="nameInput"
                        type="text"
                        v-model="commandForm.name"
                        required
                        placeholder="Enter the name of the command"
                    ></b-form-input>
                </b-form-group>
                <b-form-group 
                    id="descriptionInputGroup"
                    label="Description:"
                    label-for="descriptionInput">
                    <b-form-input
                        id="descriptionInput"
                        type="text"
                        v-model="commandForm.description"
                        required
                        placeholder="Enter the description of the command"
                    ></b-form-input>
                </b-form-group>
                <b-form-group 
                    id="formatInputGroup"
                    label="Format:"
                    label-for="formatInput">
                    <b-form-select
                        id="commandInput"
                        :options="formatOptions"
                        v-model="commandForm.format"
                        required
                    ></b-form-select>
                </b-form-group>
            </b-form>
        </b-modal>
    </div>
</template>

<script>
import VueNotifications from 'vue-notifications'
import { mapActions } from 'vuex'
import { formatOptions } from '../_fields.js'

let IPRegex = /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/

export default {
    name: 'modal',
    computed: {
        commandOptions () {
            return this.$store.state.commands.map((o) => {
                return { value: o.id, text: o.name }
            })
        }
    },
    data: function () {
        return {
            commandForm: {
                command: '',
                name: '',
                description: '',
                format: ''
            },
            groupForm: {
                name: ''
            },
            groupCommandForm: {
                command: ''
            },
            clientForm: {
                name: '',
                ip: ''
            },
            formatOptions
        }
    },
    methods: {
        ...mapActions([
            'hideClientModal',
            'hideGroupModal',
            'hideCommandModal',
            'hideGroupCommandModal',
            'createCommand',
            'createGroup',
            'createClient',
            'addGroupCommand'
        ]),
        async createClientModal (e) {
            // Validate the form
            if (this.clientForm.name === '') {
                VueNotifications.error({
                    message: 'Please specify a client name',
                    timeout: 5000
                })
                return
            }
            if (this.clientForm.ip === '') {
                VueNotifications.error({
                    message: 'Please specify a IP',
                    timeout: 5000
                })
                return
            }
            // Validate that the IP is a correct IP
            if (!IPRegex.test(this.clientForm.ip)) {
                VueNotifications.error({
                    message: 'Please specify a valid IP',
                    timeout: 5000
                })
                return
            }

            try {
                await this.createClient({name: this.clientForm.name, ip: this.clientForm.ip})
            } catch (e) {
                VueNotifications.error({
                    message: e,
                    timeout: 5000
                })
                return
            }
            this.clientForm = { name: '', ip: '' }
            this.hideClientModal()
        },
        async createGroupCommand (e) {
            if (this.groupCommandForm.command === '') {
                VueNotifications.error({
                    message: 'Please select a valid command',
                    timeout: 5000
                })
                return
            }

            try {
                await this.addGroupCommand({group_name: this.$store.state.activeGroupCommand, command_id: this.groupCommandForm.command})
            } catch (e) {
                VueNotifications.error({
                    message: e,
                    timeout: 5000
                })
                return
            }
            this.groupCommandForm = { command: '' }
            this.hideGroupCommandModal()
        },
        async createGroupModal (e) {
            if (this.groupForm.name === '') {
                VueNotifications.error({
                    message: 'Please type a valid name',
                    timeout: 5000
                })
                return
            }

            if (await this.createGroup(this.groupForm.name)) {
                VueNotifications.error({
                    message: 'This group already exists',
                    timeout: 5000
                })
                return
            }

            this.groupForm = { name: '' }
            this.hideGroupModal()
        },
        sendCommandModal (e) {
            // Evaluate all values
            if (this.commandForm.command === '') {
                VueNotifications.error({
                    message: 'Please type a valid command',
                    timeout: 5000
                })
                return
            }
            if (this.commandForm.name === '') {
                VueNotifications.error({
                    message: 'Please type a valid name',
                    timeout: 5000
                })
                return
            }
            if (this.commandForm.description === '') {
                VueNotifications.error({
                    message: 'Please type a valid description',
                    timeout: 5000
                })
                return
            }
            this.createCommand(this.commandForm)
                .then((response) => {
                    let message = {message: response.message, timeout: 5000}
                    if (response.success) VueNotifications.info(message)
                    else VueNotifications.error(message)
                }, (response) => {
                    VueNotifications.error({message: response.message, timeout: 5000})
                })

            this.commandForm = {
                command: '',
                name: '',
                description: '',
                format: ''
            }
            this.hideCommandModal()
        }
    }
}
</script>
