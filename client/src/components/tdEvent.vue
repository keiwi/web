<template>
    <div>
        {{value}}
    </div>
</template>

<script>
export default {
    name: 'td-event',
    props: ['value', 'xprops', 'row'],
    mounted () {
        let self = this
        window.jQuery(this.$el).on('click', function (e) {
            self.xprops.eventbus.$emit('click-table', e, self.row)
        })
    }
}
</script>