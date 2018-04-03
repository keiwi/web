<template>
    <div>
        <b-btn
            v-b-tooltip.hover.top
            title="Save group command"
            @click="save">
            <i class="fa fa-save"></i>
        </b-btn>
        <b-btn
            v-b-tooltip.hover.top
            title="Delete group command"
            variant="warning"
            @click="deleteCommand">
            <i class="fa fa-trash"></i>
        </b-btn>
    </div>
</template>
<script>
    import { mapActions } from 'vuex'
    import { notify } from '../../utils/_wrapper.js'

    export default {
        props: ['row', 'nested'],
        methods: {
            ...mapActions([
                'editGroup',
                'deleteGroupCommand'
            ]),
            save () {
                let cmd = this.$store.getters.getGroupCommand(this.row.id)
                if (!cmd) return
                if (cmd.next_check !== this.row.next_check) {
                    notify(this.editGroup, {
                        id: this.row.id,
                        option: 'NextCheck',
                        value: parseInt(this.row.next_check)
                    })
                }
                if (cmd.stop_error !== this.row.stop_error) {
                    notify(this.editGroup, {
                        id: this.row.id,
                        option: 'StopError',
                        value: this.row.stop_error
                    })
                }
            },
            deleteCommand () {
                notify(this.deleteGroupCommand, {id: this.row.id})
            }
        }
    }
</script>
<style>
</style>