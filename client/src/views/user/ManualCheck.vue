<template>
    <div class="row justify-content-center">
        <div class="col-md-8">
            <b-form>
                <b-form-checkbox id="checkbox1"
                    v-model="manual">
                    Manual command
                </b-form-checkbox>
                <b-form-checkbox id="checkbox2"
                    v-model="mysql">
                    Save to mysql
                </b-form-checkbox>
                <br>
                <b-form-select
                  v-if="!manual"
                  v-model="commandSelected"
                  :options="commandList"
                  style="display:block">
                </b-form-select>
                <b-form-input
                    v-if="manual"
                    v-model="command"
                    type="text"
                    placeholder="Write the command to send here">
                </b-form-input>
                <br>
                <b-button
                    type="submit"
                    variant="primary"
                    @click.prevent="sendCheck">
                    Send manual check
                </b-button>
            </b-form>
            <pre
                v-if="result"
                v-highlightjs="result"
                style="width:100%"
            ><code class="json"></code></pre>
        </div>
    </div>
</template>
<script>
    export default {
        name: 'ManualCheck',
        data () {
            return {
                manual: true,
                mysql: true,
                command: '',
                commandSelected: 0,
                result: ''
            }
        },
        methods: {
            sendCheck () {
                console.log(this)
            }
        },
        computed: {
            commandList () {
                let cmds = this.$store.state.commands.map((v) => {
                    return {
                        value: v.id,
                        text: v.name + ' (' + v.command + ')'
                    }
                })
                this.commandSelected = cmds[0].value
                return cmds
            }
        }
    }
</script>
