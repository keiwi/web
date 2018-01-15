<template>
    <div>
        <b-btn
            v-b-tooltip.hover.top
            title="Display response"
            :class="{ '-nested-comp-open-btn': isDisplayRowVisible }"
            @click="toggleNestedComp('DisplayRow')">
            <i class="fa fa-edit"></i>
        </b-btn>
    </div>
</template>
<script>
    export default {
        props: ['row', 'nested'],
        computed: {
            isDisplayRowVisible () {
                if (this.nested.comp !== 'DisplayRow') return
                return this.nested.visible
            }
        },
        methods: {
            toggleNestedComp (comp) {
                const { nested } = this
                if (nested.comp === comp) return nested.$toggle()
                nested.$toggle(comp, true)
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