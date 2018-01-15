<template>
    <div class="-nested-dsp-row-comp">
        <button class="btn btn-xs btn-link -nested-dsp-row-close-btn"
            @click="nested.$toggle(false)">
            <i class="fa fa-times fa-lg"></i>
        </button>

        <b-row style="margin: 10px 0">
            <b-col md="2" />
            <b-col md="1"><label for="input-command" style="margin-top: .5rem; margin-bottom: 0">Command:</label></b-col>
            <b-col md="6"><input id="input-command" type="input" v-model="models.Command" class="form-control" style="border-radius:0.25em"></input></b-col>
        </b-row>

        <b-row style="margin: 10px 0">
            <b-col md="2" />
            <b-col md="1"><label for="input-name" style="margin-top: .5rem; margin-bottom: 0">Name:</label></b-col>
            <b-col md="6"><input id="input-name" type="input" v-model="models.Namn" class="form-control" style="border-radius:0.25em"></input></b-col>
        </b-row>

        <b-row style="margin: 10px 0">
            <b-col md="2" />
            <b-col md="1"><label for="input-description" style="margin-top: .5rem; margin-bottom: 0">Description:</label></b-col>
            <b-col md="6"><input id="input-description" type="input" v-model="models.Description" class="form-control" style="border-radius:0.25em"></input></b-col>
        </b-row>

        <b-row style="margin: 10px 0">
            <b-col md="2" />
            <b-col md="1"><label for="input-command" style="margin-top: .5rem; margin-bottom: 0">Command:</label></b-col>
            <b-col md="6">
                <b-form-select :options="formatOptions" v-model="models.Format" />
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
                    Command: this.row.Command,
                    Description: this.row.Description,
                    Namn: this.row.Namn,
                    Format: this.row.Format
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
                if (this.row.Command !== this.models.Command) {
                    notify(this.editCommand, { id: this.row.ID, option: 'Command', value: this.models.Command })
                }
                if (this.row.Description !== this.models.Description) {
                    notify(this.editCommand, { id: this.row.ID, option: 'Description', value: this.models.Description })
                }
                if (this.row.Namn !== this.models.Namn) {
                    notify(this.editCommand, { id: this.row.ID, option: 'Namn', value: this.models.Namn })
                }
                if (this.row.Format !== this.models.Format) {
                    notify(this.editCommand, { id: this.row.ID, option: 'Format', value: this.models.Format })
                }
            },
            removeCommand () {
                notify(this.deleteCommand, this.row.ID)
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