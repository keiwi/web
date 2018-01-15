<template>
    <div>
        <b-btn
            v-b-tooltip.hover.top
            title="Show group commands"
            :class="{ '-nested-comp-open-btn': show }"
            @click="toggleNestedComp('DisplayRow')">
            <i class="fa fa-edit"></i>
        </b-btn>
        <b-btn
            v-b-tooltip.hover.top
            title="Delete group"
            variant="warning"
            @click="delGroup">
            <i class="fa fa-trash"></i>
        </b-btn>
        <b-btn
            v-b-tooltip.hover.top
            title="Add a command to the group"
            variant="success"
            @click='showGroupCommandModal(row.Name)'>
            <i class="fa fa-plus-square"></i>
        </b-btn>
    </div>
</template>
<script>
    import { mapActions } from 'vuex'
    import { notify } from '../../utils/_wrapper.js'

    export default {
        props: ['row', 'nested'],
        computed: {
            isDisplayRowVisible () {
                if (this.nested.comp !== 'DisplayRow') return
                return this.nested.visible
            }
        },
        data () {
            return {
                show: false
            }
        },
        watch: {
            show (newVal) {
                this.nested.comp = 'DisplayRow'
                this.nested.visible = newVal
            },
            nested: {
                handler (val) {
                    if (val.comp === '') {
                        this.nested.comp = 'DisplayRow'
                        this.nested.visible = this.show
                    }
                },
                deep: true
            }
        },
        methods: {
            ...mapActions([
                'toggleGroupDisplay',
                'showGroupCommandModal',
                'deleteGroup'
            ]),
            toggleNestedComp (comp) {
                this.toggleGroupDisplay(this.row.Name)
                this.show = this.$store.getters.getGroupDisplay(this.row.Name)
            },
            delGroup () {
                notify(this.deleteGroup, {name: this.row.Name})
            }
        }
}
</script>
<style>
    .-nested-comp-open-btn {
        color: #fff !important;
        background-color: #337ab7 !important;
    }
</style>