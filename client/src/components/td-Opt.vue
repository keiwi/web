<template>
    <div>
        <b-btn
            v-b-tooltip.hover.top
            title="Display command"
            :class="{ '-nested-comp-open-btn': isDisplayRowVisible }"
            @click="toggleNestedComp('DisplayRow')">
            <i class="fa fa-edit"></i>
        </b-btn>
        <b-btn
            v-b-tooltip.hover.top
            title="Delete command"
            @click="removeCommand"
            variant="warning">
            <i class="fa fa-trash"></i>
        </b-btn>
    </div>
</template>
<script>
    import { mapActions } from 'vuex'
    import { notify } from '../utils/_wrapper.js'

    export default {
        props: ['row', 'nested'],
        computed: {
            isDisplayRowVisible () {
                if (this.nested.comp !== 'DisplayRow') return
                return this.nested.visible
            }
        },
        methods: {
            ...mapActions([
                'deleteCommand'
            ]),
            toggleNestedComp (comp) {
                const { nested } = this
                if (nested.comp === comp) return nested.$toggle()
                nested.$toggle(comp, true)
            },
            removeCommand () {
                notify(this.deleteCommand, this.row.ID)
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