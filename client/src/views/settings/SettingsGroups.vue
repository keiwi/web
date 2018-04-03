<template>
    <div class='settings-commands'>
        <h6>Groups</h6>
        <div class='btn-group btn-group-md' style='margin-bottom:10px'>
            <button type='button' class='btn btn-warning' style='border-radius:5px 0 0 5px' @click='removeGroups'>
                <i class='fa fa-trash'></i> Delete
            </button>
            <button type='button' class='btn btn-primary' style='border-radius:0 5px 5px 0' @click='showGroupModal'>
                <i class='fa fa-plus'></i> Add a new Group
            </button>
        </div>
        <datatable
            v-bind='$data'
        />
    </div>
</template>

<script>
    import { mapActions, mapGetters } from 'vuex'
    import FilterTh from '../../components/th-Filter.vue'
    import GroupsTd from './groups-td.vue'
    import DisplayRow from './groups-display.vue'
    import { formatOptions } from '../../_fields.js'
    import filter from '../../utils/_filter.js'
    import CloneDeep from 'lodash/CloneDeep'
    import VueNotifications from 'vue-notifications'

    export default {
        name: 'settings-groups',
        components: {
            FilterTh,
            GroupsTd,
            DisplayRow
        },
        mounted () {
            let self = this
            setTimeout(function () {
                self.handleQueryChange()
            }, 1000)
        },
        computed: {
            ...mapGetters([
                'getCommand'
            ]),
            values () {
                let values = this.$store.state.groups
                values = CloneDeep(values)
                for (let i in values) {
                    values[i].amount = values[i].commands.length
                    for (let j in values[i].commands) {
                        let cmd = this.getCommand(values[i].commands[j].command_id)
                        if (typeof cmd === 'undefined') continue
                        values[i].commands[j].command = cmd.command
                    }
                }
                return values
            }
        },
        data () {
            return {
                formatOptions,
                supportNested: true,
                select: false,
                HeaderSettings: false,
                'tbl-class': ['table-bordered'],
                columns: [
                    { title: 'Name', field: 'name', thComp: 'FilterTh', tdStyle: { fontStyle: 'italic' }, sortable: true },
                    { title: 'Commands', field: 'amount', sortable: true },
                    { title: 'Actions', tdComp: 'GroupsTd' }
                ],
                query: { 'limit': 10, 'offset': 0, 'sort': '', 'order': '' },
                total: 0,
                data: [],
                selection: []
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
        },
        methods: {
            ...mapActions([
                'showGroupModal',
                'initGroups',
                'deleteGroup'
            ]),
            handleQueryChange () {
                let data = filter(this.values, this.query)
                this.data = data.rows
                this.total = data.total
            },
            removeGroups () {
                if (this.selection.length <= 0) {
                    VueNotifications.error({message: 'No group has been selected.'})
                    return
                }

                for (let row of this.selection) {
                    this.deleteGroup({name: row.name}).then((response) => {
                        if (!response.success) VueNotifications.error({message: response.message})
                    }, (response) => VueNotifications.error({message: response}))
                }

                VueNotifications.info({message: `Tried deleting ${this.selection.length} group(s).`})
            }
        }
    }
</script>