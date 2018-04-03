<template>
    <div class="user-table">
        <datatable
            v-bind='$data'
        />
    </div>
</template>

<script>
    import API from '../../api'
    import VueNotifications from 'vue-notifications'
    import OptTd from './td-Opt.vue'
    import DisplayRow from './checks-displayRow.vue'
    import filter from '../../utils/_filter.js'

    export default {
        props: ['user-data'],
        name: 'latest-checks',
        components: {
            DisplayRow,
            OptTd
        },
        mounted () {
            let self = this
            setTimeout(function () {
                self.handleQueryChange()
            }, 1000)
        },
        data () {
            return {
                HeaderSettings: false,
                'tbl-class': ['table-bordered'],
                supportNested: true,
                columns: [
                    { title: 'ID', field: 'id' },
                    { title: 'Command Name', field: 'command_name' },
                    { title: 'Command ID', field: 'command_id' },
                    { title: 'Error', field: 'error' },
                    { title: 'Finished', field: 'finished' },
                    { title: 'Actions', tdComp: 'OptTd' }
                ],
                query: { 'limit': 10, 'offset': 0, 'sort': '', 'order': '' },
                total: 0,
                data: []
            }
        },
        watch: {
            query: {
                handler () {
                    this.handleQueryChange()
                },
                deep: true
            },
            checks: {
                handler () {
                    this.handleQueryChange()
                }
            }
        },
        methods: {
            handleQueryChange () {
                let data = filter(this.checks || [], this.query)
                this.data = data.rows
                this.total = data.total
            }
        },
        asyncComputed: {
            async checks () {
                let cmds = []
                if (this.userData == null) return []

                if (typeof this.userData.groups !== 'undefined') {
                    for (let group of this.userData._group_ids) {
                        for (let g of this.$store.state.groups) {
                            if (group === g.id) {
                                for (let cmd of g.commands) {
                                    let exists = false
                                    for (let id of cmds) {
                                        if (id === cmd.command_id) {
                                            exists = true
                                            break
                                        }
                                    }
                                    if (!exists) {
                                        cmds.push(cmd.command_id)
                                    }
                                }
                            }
                        }
                    }
                }

                let data = []
                try {
                    data = await API.getClientCMDIDChecks({client_id: this.userData.id, command_id: cmds})
                } catch (e) {
                    VueNotifications.error({message: e})
                }

                for (let i in data) {
                    for (let c of this.$store.state.commands) {
                        if (data[i].command_id === c.id) {
                            data[i].command_name = c.name
                            break
                        }
                    }
                }

                return data
            }
        }
    }
</script>