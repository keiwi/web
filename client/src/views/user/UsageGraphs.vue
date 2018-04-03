<template>
    <!-- <div class='chart' style="height: 700px" /> -->
    <div class="chart" style="width: 90%; height: 700px"></div>
</template>

<script>
    // import SimpleGraph from '../../utils/_graphs'
    // import moment from 'moment'
    import Plotly from 'plotly.js/dist/plotly'

    export default {
        props: ['values', 'format', 'chartData'],
        mounted () {
            console.log(this.chartData)
            this.graph = Plotly.plot(this.$el, this.chartData.dataSets, this.chartData.layout)
            /* this.graphs = new SimpleGraph(this.$el, {
                onDataHover (value, index) {
                    return `
                        Client ID: ${value.item.client_id}<br>
                        Command ID: ${value.item.command_id}<br>
                        ID: ${value.item.id}<br>
                        Error: ${value.item.error}<br>
                        Checked: ${value.item.checked}<br>
                        Finished: ${value.item.finished}<br>
                        Created at: ${moment(value.item.created_at).format('YYYY-MM-DD HH:mm:ss')}<br>
                        Response: ${value.item.response}<br>
                    `
                }
            })
            this.graphs.update(this.chartData) */
        },
        watch: {
            chartData: {
                handler (newValue) {
                    console.log(newValue)
                    this.graph = Plotly.react(this.$el, newValue.dataSets, newValue.layout)
                },
                nested: true
            }
        },
        data () {
            return {
                graph: {}
            }
        }
    }
</script>

<style>
    /* 13. Basic Styling with CSS */
    div.tooltip {	
        position: absolute;
        text-align: center;
        padding: 8px;
        font: 12px sans-serif;
        background: #555;
        color: white;
        border: 0px;
        border-radius: 6px 6px 6px 0;
        pointer-events: none;
    }

    .close-tooltip {
        cursor: pointer;
        pointer-events: all;
    }
    .tooltip-body {
        pointer-events: all;
    }
</style>
