<template>
    <div class='settings-commands'>
        <h6>Commands</h6>
        <div class='btn-group btn-group-md' style='margin-bottom:10px'>
            <button type='button' class='btn btn-warning' style='border-radius:5px 0 0 5px' @click='removeCommands'>
                <i class='fa fa-trash'></i> Delete
            </button>
            <button type='button' class='btn btn-primary' style='border-radius:0 5px 5px 0' @click='showCommandModal'>
                <i class='fa fa-plus'></i> Add a new Command
            </button>
        </div>
        <datatable
            v-bind='$data'
        />
    </div>
</template>

<script>
import { mapActions } from 'vuex'
import FilterTh from '../../components/th-Filter.vue'
import OptTd from '../../components/td-Opt.vue'
import DisplayRow from './commands-displayRow.vue'
import { formatOptions } from '../../_fields.js'
import filter from '../../utils/_filter.js'
import VueNotifications from 'vue-notifications'

export default {
    name: 'settings-commands',
    components: {
        FilterTh,
        OptTd,
        DisplayRow
    },
    computed: {
        values () {
            return this.$store.state.commands
        }
    },
    mounted () {
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
                { title: 'ID', field: 'id', sortable: true },
                { title: 'Command', field: 'command', thComp: 'FilterTh', tdStyle: { fontStyle: 'italic' } },
                { title: 'Name', field: 'name', sortable: true },
                { title: 'Description', field: 'description' },
                { title: 'Format', field: 'format', sortable: true },
                { title: 'Actions', tdComp: 'OptTd' }
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
            'showCommandModal',
            'deleteCommand'
        ]),
        handleQueryChange () {
            let data = filter(this.values, this.query)
            this.data = data.rows
            this.total = data.total
        },
        removeCommands () {
            if (this.selection.length <= 0) {
                VueNotifications.error({message: 'No commands has been selected.'})
                return
            }

            for (let row of this.selection) {
                this.deleteCommand(row.id).then((response) => {
                    if (!response.success) VueNotifications.error({message: response.message})
                }, (response) => VueNotifications.error({message: response}))
            }

            VueNotifications.info({message: `Tried deleting ${this.selection.length} command(s).`})
        }
    }
}
</script>