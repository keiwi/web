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
                let cmd = this.$store.getters.getGroupCommand(this.row.ID)
                if (!cmd) return
                if (cmd.NextCheck !== this.row.NextCheck) {
                    this.editGroup({
                        id: this.row.ID,
                        option: 'NextCheck',
                        value: parseInt(this.row.NextCheck)
                    })
                }
                if (cmd.StopError !== this.row.StopError) {
                    this.editGroup({
                        id: this.row.ID,
                        option: 'StopError',
                        value: this.row.StopError
                    })
                }
            },
            deleteCommand () {
                notify(this.deleteGroupCommand, {id: this.row.ID})
            }
        }
    }
</script>
<style>
</style>