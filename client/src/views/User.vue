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
    import { mapGetters } from 'vuex'
    import UserInformation from './user/UserInformation'
    import LatestChecks from './user/LatestChecks'
    import UsageGraphs from './user/UsageGraphs'
    import ManualCheck from './user/ManualCheck'
    import VueNotifications from 'vue-notifications'
    import moment from 'moment'
    import API from '../api'

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
            ...mapGetters([
                'getGroup'
            ]),
            userInfoColumns () {
                let groups = ''
                let ip = ''
                let name = ''
                if (this.userData != null) {
                    groups = this.userData.groups
                    ip = this.userData.ip
                    name = this.userData.name
                }

                return [
                    { title: 'ID', 'field': 'id' },
                    { title: 'Namn', 'field': 'name', edit: 'input', 'value': name },
                    { title: 'IP', 'field': 'ip', edit: 'input', 'value': ip },
                    { title: 'Latest Check', 'field': 'latest' },
                    { title: 'Groups', 'field': 'groups', edit: 'dropdown', 'values': this.$store.state.groups.map(group => ({ name: group.name, id: group.id, checked: groups.includes(group.name) })) },
                    { title: 'Uptime', 'field': 'uptime' },
                    { title: 'Hostname', 'field': 'hostname' },
                    { title: 'OS', 'field': 'os' },
                    { title: 'Platform', 'field': 'platform' },
                    { title: 'Client Version', 'field': 'version' }
                ]
            }
        },
        asyncComputed: {
            userData: {
                async get () {
                    let data = {}
                    try {
                        data = await API.getIDClients({id: this.$route.params.id})
                    } catch (e) {
                        VueNotifications.error({message: e})
                        return
                    }

                    var groups = []
                    for (var g of data.group_ids) {
                        var group = this.getGroup(g)
                        if (typeof group === 'undefined') continue
                        groups.push(group.name)
                    }

                    return {
                        id: data.id,
                        name: data.name,
                        ip: data.ip,
                        groups: groups.join(','),
                        latest: '',
                        uptime: '',
                        hostname: '',
                        os: '',
                        platform: '',
                        version: '',
                        _group_ids: data.group_ids
                    }
                },
                watch () {
                    return this.$store.state.groups
                }
            }
        },
        name: 'User',
        methods: {
            getChecks (i) {
                API.getChecksDateClient({
                    client_id: this.$route.params.id,
                    command_id: this.graphs[i].command_id,
                    from: moment().subtract(7, 'days').format('YYYY-MM-DD HH:mm:ss'),
                    to: moment().format('YYYY-MM-DD HH:mm:ss'),
                    max: 50
                }).then((resp) => {
                    var dataSets = []
                    if (resp != null) {
                        for (var j = resp.length - 1; j >= 0; j--) {
                            let d = new Date(resp[j].created_at)
                            if (Array.isArray(this.graphs[i].Key)) {
                                for (var k = 0; k < this.graphs[i].Key.length; k++) {
                                    if (dataSets.length < this.graphs[i].Key.length) {
                                        dataSets[k] = newDataset(this.graphs[i])
                                    }

                                    if (resp[j].error || resp[j].response === '') {
                                        dataSets[k].x.push(d)
                                        dataSets[k].y.push(0)
                                        continue
                                    }
                                    dataSets[k].x.push(d)
                                    dataSets[k].y.push(JSON.parse(resp[j].response).message[this.graphs[i].Key[k]])
                                    dataSets[k].text.push(formatCheck(resp[j]))
                                }
                            } else {
                                if (dataSets.length <= 0) {
                                    dataSets[0] = newDataset(this.graphs[i])
                                }
                                if (resp[j].error || resp[j].response === '') {
                                    dataSets[0].x.push(d)
                                    dataSets[0].y.push(0)
                                    continue
                                }
                                dataSets[0].x.push(d)
                                dataSets[0].y.push(JSON.parse(resp[j].response).message[this.graphs[i].Key])
                                dataSets[0].text.push(formatCheck(resp[j]))
                            }
                        }
                    }

                    var yAxis = '.4s'
                    var ySuffix = ''
                    if (this.graphs[i].Format === 'procent') {
                        yAxis = '.4s'
                        ySuffix = '%'
                    } else if (this.graphs[i].Format.toLowerCase() === 'b') {
                        yAxis = '.3s'
                        ySuffix = this.graphs[i].Format
                    }
                    this.$set(this.graphs[i], 'chartData', {
                        dataSets,
                        layout: {
                            yaxis: {
                                tickformat: yAxis,
                                ticksuffix: ySuffix
                            },
                            margin: { t: 10 },
                            xaxis: {
                                autorange: true,
                                rangeselector: {
                                    buttons: [
                                        {
                                            count: 5,
                                            label: '5m',
                                            step: 'minute',
                                            stepmode: 'backward'
                                        },
                                        {
                                            count: 1,
                                            label: '1h',
                                            step: 'hour',
                                            stepmode: 'backward'
                                        },
                                        {
                                            count: 1,
                                            label: '1d',
                                            step: 'day',
                                            stepmode: 'backward'
                                        },
                                        {
                                            count: 7,
                                            label: '7d',
                                            step: 'day',
                                            stepmode: 'backward'
                                        },
                                        {
                                            count: 1,
                                            label: '1m',
                                            step: 'month',
                                            stepmode: 'backward'
                                        },
                                        {
                                            count: 3,
                                            label: '3m',
                                            step: 'month',
                                            stepmode: 'backward'
                                        },
                                        {
                                            count: 1,
                                            label: '1y',
                                            step: 'year',
                                            stepmode: 'backward'
                                        },
                                        {step: 'all'}
                                    ]
                                },
                                rangeslider: {range: ['2015-02-17', '2017-02-16']},
                                type: 'date'
                            }
                        }
                    })
                }, (e) => {
                    VueNotifications.error({message: e})
                })
            }
        },
        data () {
            return {
                graphs: [
                    { Title: 'CPU Usage', command_id: '5abc2b302664138444e1d679', Key: 'procent', Labels: 'CPU Usage', Format: 'procent', chartData: {} },
                    { Title: 'RAM Usage', command_id: '5abc2b302664138444e1d673', Key: 'size', Labels: 'RAM Usage', Format: 'B', chartData: {} },
                    { Title: 'Network Usage', command_id: '5abc2b302664138444e1d67e', Key: ['sent', 'recv'], Labels: ['Network Sent', 'Network Recieved'], Format: 'b', chartData: {} }
                ]
            }
        }
    }

    function formatCheck (check) {
        return moment(check.created_at).format('YYYY-MM-DD HH:mm:ss')
    }

    function newDataset (graph) {
        return {
            x: [],
            y: [],
            name: graph.Labels,
            mode: 'lines+markers',
            text: [],
            textposition: 'none',
            type: 'scatter',
            hoverlabel: {
                namelength: -1
            },
            hoverinfo: 'y+text'
        }
    }
</script>