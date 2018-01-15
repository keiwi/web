<template>
    <div class="animated fadeIn">
        <b-card>
            <h4 class="card-title mb-0">User</h4>
            <div slot="footer">
                <div class="row">
                    <div class="col-md-5">
                        <b-card>
                            <h4 class="card-title mb-0">User information</h4>
                            <div slot="footer">
                                <user-information :user-data="userData" :columns="userInfoColumns" />
                            </div>
                        </b-card>
                    </div>
                    <div class="col-md-7">
                        <b-card>
                            <h4 class="card-title mb-0">Latest checks</h4>
                            <div slot="footer">
                                <latest-checks :user-data="userData" />
                            </div>
                        </b-card>
                    </div>
                </div>
                <b-card>
                    <h4 class="card-title mb-0">Manual check</h4>
                    <div slot="footer">
                        <manual-check />
                    </div>
                </b-card>
                <div v-for="(g, i) in graphs" :key="i" class="row">
                    <div class="col-md-12">
                        <b-card>
                            <h4 class="card-title mb-0">{{g.Title}}</h4>
                            <div slot="footer">
                                <usage-graphs :chartData="g.chartData" :format="g.Format" />
                            </div>
                        </b-card>
                    </div>
                </div>
            </div>
        </b-card>
    </div>
</template>

<script>
    import UserInformation from './user/UserInformation'
    import LatestChecks from './user/LatestChecks'
    import UsageGraphs from './user/UsageGraphs'
    import ManualCheck from './user/ManualCheck'
    import VueNotifications from 'vue-notifications'
    import moment from 'moment'
    import API from '../api'
    import chroma from 'chroma-js'
    import * as d3 from 'd3'

    // const brandSuccess = '#4dbd74'
    const brandInfo = '#63c2de'
    const brandDanger = '#f86c6b'
    const colour = chroma.scale([brandInfo, brandDanger])

    export default {
        components: {
            UserInformation,
            LatestChecks,
            UsageGraphs,
            ManualCheck
        },
        mounted () {
            for (let i in this.graphs) {
                this.getChecks(i)
            }
        },
        computed: {
            userInfoColumns () {
                let groups = ''
                let ip = ''
                let name = ''
                if (this.userData != null) {
                    groups = this.userData.Groups.split(',')
                    ip = this.userData.IP
                    name = this.userData.Name
                }

                return [
                    { title: 'ID', 'field': 'ID' },
                    { title: 'Namn', 'field': 'Name', edit: 'input', 'value': name },
                    { title: 'IP', 'field': 'IP', edit: 'input', 'value': ip },
                    { title: 'Latest Check', 'field': 'Latest' },
                    { title: 'Groups', 'field': 'Groups', edit: 'dropdown', 'values': this.$store.state.groups.map(group => ({ Name: group.Name, Checked: groups.includes(group.Name) })) },
                    { title: 'Uptime', 'field': 'Uptime' },
                    { title: 'Hostname', 'field': 'Hostname' },
                    { title: 'OS', 'field': 'OS' },
                    { title: 'Platform', 'field': 'Platform' },
                    { title: 'Client Version', 'field': 'Version' }
                ]
            }
        },
        asyncComputed: {
            async userData () {
                let data = {}
                try {
                    data = await API.getIDClients({id: parseInt(this.$route.params.id)})
                } catch (e) {
                    VueNotifications.error({message: e})
                }

                return {
                    ID: data.id,
                    Name: data.namn,
                    IP: data.ip,
                    Groups: data.group_names,
                    Latest: '',
                    Uptime: '',
                    Hostname: '',
                    OS: '',
                    Platform: '',
                    Version: ''
                }
            }
        },
        name: 'User',
        methods: {
            getChecks (i) {
                API.getChecksDateClient({
                    client_id: parseInt(this.$route.params.id),
                    command_id: this.graphs[i].CommandID,
                    from: moment().subtract(7, 'days').format('YYYY-MM-DD hh:mm:ss'),
                    to: moment().format('YYYY-MM-DD hh:mm:ss'),
                    max: 50
                }).then((resp) => {
                    let self = this
                    var dataSets = []
                    for (var j = resp.length - 1; j >= 0; j--) {
                        let d = new Date(resp[j].created_at)
                        if (Array.isArray(this.graphs[i].Key)) {
                            for (var k = 0; k < this.graphs[i].Key.length; k++) {
                                if (dataSets.length < this.graphs[i].Key.length) {
                                    dataSets[k] = {
                                        label: this.graphs[i].Key[k],
                                        styleLine: (function (index) {
                                            return function (style) {
                                                style.stroke = colour(index / (self.graphs[i].Key.length - 1))
                                                return style
                                            }
                                        })(k),
                                        styleDot: (function (index) {
                                            return function (style) {
                                                style.fill = colour(index / (self.graphs[i].Key.length - 1))
                                                return style
                                            }
                                        })(k),
                                        data: []
                                    }
                                }

                                if (resp[j].error) {
                                    dataSets[k].data.push({y: 0, x: d})
                                    continue
                                }
                                dataSets[k].data.push({
                                    y: JSON.parse(resp[j].response)[this.graphs[i].Key[k]],
                                    x: d,
                                    item: resp[j] // Custom data provided
                                })
                            }
                        } else {
                            if (dataSets.length <= 0) {
                                dataSets[0] = {
                                    label: this.graphs[i].Key,
                                    styleLine: function (style) {
                                        style.stroke = colour(0)
                                        return style
                                    },
                                    styleDot: function (style) {
                                        style.fill = colour(0)
                                        return style
                                    },
                                    data: []
                                }
                            }
                            if (resp[j].error) {
                                dataSets[0].data.push(0)
                                continue
                            }
                            dataSets[0].data.push({
                                y: JSON.parse(resp[j].response)[this.graphs[i].Key],
                                x: d,
                                item: resp[j] // Custom data provided
                            })
                        }
                    }

                    this.$set(this.graphs[i], 'chartData', {dataSets, options: {xFormat: d3.timeFormat('%y-%m-%d %H:%M:%S')}})
                }, (e) => {
                    VueNotifications.error({message: e})
                })
            }
        },
        data () {
            return {
                graphs: [
                    { Title: 'CPU Usage', CommandID: 18, Key: 'procent', Labels: 'CPU Usage', Format: 'procent', chartData: {} },
                    { Title: 'RAM Usage', CommandID: 12, Key: 'size', Labels: 'RAM Usage', Format: 'B', chartData: {} },
                    { Title: 'Network Usage', CommandID: 23, Key: ['sent', 'recv'], Labels: ['Network Sent', 'Network Recieved'], Format: 'b', chartData: {} }
                ]
            }
        }
    }
</script>