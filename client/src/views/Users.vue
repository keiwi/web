<template>
    <div class="animated fadeIn">
        <b-card>
            <h4 class="card-title mb-0">Users</h4>
            <div slot="footer">
                <div class='users-table'>
                    <div class='btn-group btn-group-md' style='margin-bottom:10px'>
                        <button type='button' class='btn btn-warning' style='border-radius:5px 0 0 5px' @click='deleteClientModal'>
                            <i class='fa fa-trash'></i> Delete
                        </button>
                        <button type='button' class='btn btn-primary' style='border-radius:0 5px 5px 0' @click='showClientModal'>
                            <i class='fa fa-plus'></i> Create a user
                        </button>
                    </div>
                    <datatable
                        v-bind='$data'
                    />
                </div>
            </div>
        </b-card>
    </div>
</template>

<script>
import { mapActions } from 'vuex'
import filter from '../utils/_filter.js'
import CloneDeep from 'lodash/CloneDeep'
import VueNotifications from 'vue-notifications'
import tdEvent from '../components/tdEvent.vue'
import Vue from 'vue'

let eventbus = new Vue()

export default {
    name: 'dashboard',
    components: {tdEvent},
    mounted () {
        let self = this
        setTimeout(function () {
            self.handleQueryChange()
        }, 1000)
        eventbus.$on('click-table', (e, row) => {
            this.$router.push({ name: 'User', params: { id: row.id } })
        })
    },
    computed: {
        values () {
            let values = this.$store.state.clients
            values = CloneDeep(values)
            for (let i in values) {
                values[i].groups = values[i].groups.join(', ')
            }
            return values
        }
    },
    data () {
        return {
            supportNested: true,
            select: false,
            HeaderSettings: false,
            'tbl-class': ['table-bordered'],
            columns: [
                { title: 'ID', field: 'id', sortable: true, tdComp: 'tdEvent' },
                { title: 'Name', field: 'name', sortable: true, tdComp: 'tdEvent' },
                { title: 'IP', field: 'ip', sortable: true, tdComp: 'tdEvent' },
                { title: 'Groups', field: 'groups', sortable: true, tdComp: 'tdEvent' },
                { title: 'Latest Check', field: 'latest', sortable: true, tdComp: 'tdEvent' },
                { title: 'Active', field: 'active', sortable: true, tdComp: 'tdEvent' }
            ],
            query: { 'limit': 10, 'offset': 0, 'sort': '', 'order': '' },
            total: 0,
            data: [],
            selection: [],
            xprops: {eventbus}
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
            'showClientModal',
            'deleteClient'
        ]),
        handleQueryChange () {
            let data = filter(this.values, this.query)
            this.data = data.rows
            this.total = data.total
        },
        deleteClientModal () {
            if (this.selection.length <= 0) {
                VueNotifications.error({message: 'No clients has been selected.'})
                return
            }

            for (let row of this.selection) {
                this.deleteClient({id: row.id}).then((response) => {
                    if (!response.success) VueNotifications.error({message: response.message})
                }, (response) => VueNotifications.error({message: response}))
            }

            VueNotifications.info({message: `Tried deleting ${this.selection.length} client(s).`})
        }
    }
}
</script>