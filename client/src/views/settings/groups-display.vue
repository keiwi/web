<template>
    <div class="-nested-dsp-row-comp">
        <datatable v-bind="$data" />
    </div>
</template>
<script>
    import { mapActions } from 'vuex'
    import { formatOptions } from '../../_fields.js'
    import filter from '../../utils/_filter.js'
    import FilterTh from '../../components/th-Filter.vue'
    import Timepicker from '../../components/timepicker.vue'
    import ToggleTd from '../../components/td-Toggle.vue'
    import GroupsAction from './groups-action.vue'
    import CloneDeep from 'lodash/CloneDeep'
    
    export default {
        props: ['row', 'nested'],
        components: {
            FilterTh,
            Timepicker,
            ToggleTd,
            GroupsAction
        },
        computed: {
            models () {
                console.log(this.row)
                return {
                    Command: this.row.command,
                    Description: this.row.description,
                    Name: this.row.name,
                    Format: this.row.format
                }
            },
            values () {
                return this.row.commands
            }
        },
        mounted () {
            // TODO: Remove this once initial get work
            let self = this
            setTimeout(function () {
                self.handleQueryChange()
            }, 1000)
        },
        data () {
            return {
                formatOptions,
                supportNested: true,
                select: false,
                HeaderSettings: false,
                'tbl-class': ['table-bordered'],
                columns: [
                    { title: 'ID', field: 'id', thComp: 'FilterTh', tdStyle: { fontStyle: 'italic' }, sortable: true },
                    { title: 'Command ID', field: 'command_id', sortable: true },
                    { title: 'Command', field: 'command', sortable: true },
                    { title: 'Next Check', field: 'next_check', tdComp: 'Timepicker', sortable: true, colStyle: {width: '20%'} },
                    { title: 'Stop Error', field: 'stop_error', tdComp: 'ToggleTd', sortable: true, colStyle: {width: '15%'} },
                    { title: 'Action', tdComp: 'GroupsAction' }
                ],
                query: { 'limit': 10, 'offset': 0, 'sort': '', 'order': '' },
                total: 0,
                data: [],
                selection: []
            }
        },
        methods: {
            ...mapActions([
            ]),
            handleQueryChange () {
                let data = filter(this.values, this.query)
                this.data = CloneDeep(data.rows)
                this.total = data.total
            }
        },
        watch: {
            query: {
                handler () {
                    this.handleQueryChange()
                },
                deep: true
            },
            values: {
                handler () {
                    this.handleQueryChange()
                }
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