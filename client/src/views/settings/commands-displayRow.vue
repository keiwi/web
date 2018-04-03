<template>
    <div class="-nested-dsp-row-comp">
        <button class="btn btn-xs btn-link -nested-dsp-row-close-btn"
            @click="nested.$toggle(false)">
            <i class="fa fa-times fa-lg"></i>
        </button>

        <b-row style="margin: 10px 0">
            <b-col md="2" />
            <b-col md="1"><label for="input-command" style="margin-top: .5rem; margin-bottom: 0">Command:</label></b-col>
            <b-col md="6"><input id="input-command" type="input" v-model="models.command" class="form-control" style="border-radius:0.25em"></input></b-col>
        </b-row>

        <b-row style="margin: 10px 0">
            <b-col md="2" />
            <b-col md="1"><label for="input-name" style="margin-top: .5rem; margin-bottom: 0">Name:</label></b-col>
            <b-col md="6"><input id="input-name" type="input" v-model="models.name" class="form-control" style="border-radius:0.25em"></input></b-col>
        </b-row>

        <b-row style="margin: 10px 0">
            <b-col md="2" />
            <b-col md="1"><label for="input-description" style="margin-top: .5rem; margin-bottom: 0">Description:</label></b-col>
            <b-col md="6"><input id="input-description" type="input" v-model="models.description" class="form-control" style="border-radius:0.25em"></input></b-col>
        </b-row>

        <b-row style="margin: 10px 0">
            <b-col md="2" />
            <b-col md="1"><label for="input-command" style="margin-top: .5rem; margin-bottom: 0">Format:</label></b-col>
            <b-col md="6">
                <b-form-select :options="formatOptions" v-model="models.format" />
            </b-col>
        </b-row>

        <b-row style="margin: 10px 0">
            <b-col md="4" />
            <b-col md="3">
                <div class='btn-group btn-group-md' style='margin-bottom:10px; width: 100%'>
                    <button type='button' class='btn btn-primary' style='border-radius: 5px 0 0 5px; width: 100%' @click='saveCommand'>
                        <i class='fa fa-save'></i> Save
                    </button>
                    <button type='button' class='btn btn-warning' style='border-radius: 0 5px 5px 0; width: 100%' @click='removeCommand'>
                        <i class='fa fa-trash'></i> Delete
                    </button>
                </div>
            </b-col>
        </b-row>
    </div>
</template>
<script>
    import { mapActions } from 'vuex'
    import { formatOptions } from '../../_fields.js'
    import { notify } from '../../utils/_wrapper.js'
    
    export default {
        props: ['row', 'nested'],
        computed: {
            models () {
                return {
                    command: this.row.command,
                    description: this.row.description,
                    name: this.row.name,
                    format: this.row.format
                }
            }
        },
        data () {
            return {
                formatOptions
            }
        },
        methods: {
            ...mapActions([
                'editCommand',
                'deleteCommand'
            ]),
            saveCommand () {
                if (this.row.command !== this.models.command) {
                    notify(this.editCommand, { id: this.row.id, option: 'command', value: this.models.command })
                }
                if (this.row.description !== this.models.description) {
                    notify(this.editCommand, { id: this.row.id, option: 'description', value: this.models.description })
                }
                if (this.row.name !== this.models.name) {
                    notify(this.editCommand, { id: this.row.id, option: 'name', value: this.models.name })
                }
                if (this.row.format !== this.models.format) {
                    notify(this.editCommand, { id: this.row.id, option: 'format', value: this.models.format })
                }
            },
            removeCommand () {
                notify(this.deleteCommand, this.row.id)
            }
        }
    }
</script>
<style>
    .-nested-dsp-row-comp {
        position: relative;
        padding: 10px;
    }
    .-nested-dsp-row-close-btn {
        position: absolute;
        top: 5px;
        right: 5px;
    }
</style>