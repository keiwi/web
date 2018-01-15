<template>
    <div class="-nested-dsp-row-comp">
        <div class="btn-group pull-right">
            <button type="button" class="btn btn-outline-secondary btn-sm" :class="{active:raw}" @click="raw = true">Raw</button>
            <button type="button" class="btn btn-outline-secondary btn-sm" :class="{active:!raw}" @click="raw = false">Formatted</button>
        </div>
        <pre v-show="raw" v-highlightjs="pretty" style="width:100%"><code class="json"></code></pre>

        <div v-show="!raw">
            <div v-for="(r, key) in JSON.parse(this.row.response)" :key="key">
                <b>{{key}}: </b>
                <template v-if="r == ''">""</template>
                <template v-else-if="Array.isArray(r)">
                    <div v-for="(v, index) in r" :key="index">
                        <div v-if="typeof v == 'object'" v-for="(o, i) in v" :key="i">
                            <b>&nbsp;&nbsp;&nbsp;&nbsp;{{i}}:</b>
                            <template v-if="typeof o == 'number' 
                                && row.command_id in format
                                && (
                                    (typeof format[row.command_id].Key == 'string' && format[row.command_id].Key == i) ||
                                    (format[row.command_id].Key.includes(i))
                                    )">
                                <byte-converter v-if="format[row.command_id]" :value='o' :format="format[row.command_id].Value" />
                            </template>
                            <template v-else>
                                {{o}}
                            </template>
                        </div>
                        <br>
                    </div>
                </template>
                <template v-else-if="typeof r == 'number' 
                    && row.command_id in format
                    && (
                        (typeof format[row.command_id].Key == 'string' && format[row.command_id].Key == key) ||
                        (format[row.command_id].Key.includes(key))
                        )">
                        <template v-if="format[row.command_id].Key == 'procent'">{{fixedDecimals(r) + '%'}}</template>
                        <template v-else-if="format[row.command_id].Key == 'uptime'">{{fixedUptime(r)}}</template>
                        <template v-else>
                            <byte-converter v-if="format[row.command_id]" :value='r' :format="format[row.command_id].Value" />
                        </template>
                </template>
                <template v-else>{{r}}</template>
            </div>
        </div>
    </div>
</template>
<script>
    import ByteConverter from '../../components/ByteConverter.vue'
    import { fixedDecimals, fixedUptime } from '../../utils/_convert'

    export default {
        props: ['row', 'nested'],
        components: {
            ByteConverter
        },
        methods: {
            fixedDecimals,
            fixedUptime
        },
        data () {
            return {
                raw: true
            }
        },
        computed: {
            pretty () {
                return JSON.stringify(JSON.parse(this.row.response), null, 4)
            },
            format () {
                return {
                    12: { Key: 'size', Value: 'B' },
                    16: { Key: 'size', Value: 'B' },
                    18: { Key: 'procent', Value: 'procent' },
                    20: { Key: 'uptime', Value: 'uptime' },
                    23: { Key: ['sent', 'recv'], Value: 'b' }
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