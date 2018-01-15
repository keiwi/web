<template>
    <div class='chart' style="height: 700px" />
</template>

<script>
    import SimpleGraph from '../../utils/_graphs'
    import moment from 'moment'

    export default {
        props: ['values', 'format', 'chartData'],
        mounted () {
            this.graphs = new SimpleGraph(this.$el, {
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
            this.graphs.update(this.chartData)
        },
        watch: {
            chartData: {
                handler (newValue) {
                    if (this.format === 'procent') {
                        this.graphs.options({yPrefix: '.2f'})
                    }
                    this.graphs.update(newValue)
                },
                nested: true
            }
        },
        data () {
            return {
                graphs: {}
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
