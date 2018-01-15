<template>
    <span>
        {{formatted_value}} {{form}}
        <div class="btn-group">
            <button
                class="btn btn-primary btn-sm"
                @click="expand = !expand"
                :class="{active: expand}"
                :style="{padding: '0px 6px'}">
                    <i class="fa fa-ellipsis-v"></i>
            </button>
            <button
                v-show="expand"
                v-for="(f, key) in formats"
                :key="key"
                :style="{padding: '0px 6px'}"
                :class="{active: form == f}"
                @click="form = f"
                class="btn btn-primary btn-sm">
                    {{f}}
            </button>
        </div>
    </span>
</template>

<script>
    import { convertBytesBits } from '../utils/_convert'
    export default {
        props: ['value', 'format'],
        name: 'latest-checks',
        mounted () {
            this.form = this.format
            this.formatted_value = this.value

            if (this.format === this.format.toUpperCase()) {
                this.type = 'byte'
                this.formats = this.formatsBytes
            } else {
                this.type = 'bit'
                this.formats = this.formatsBits
            }

            this.raw_value = convertBytesBits(this.value, this.format)
        },
        data () {
            return {
                form: 'B',
                formats: [],
                formatsBytes: ['B', 'KB', 'MB', 'GB', 'TB'],
                formatsBits: ['b', 'Kb', 'Mb', 'Gb', 'Tb'],
                formatted_value: '',
                type: 'byte',
                raw_value: 0,
                expand: false
            }
        },
        watch: {
            form (newValue) {
                this.formatted_value = convertBytesBits(this.raw_value, newValue)
            }
        }
    }
</script>