<template>
    <div class="btn-group">
        {{ title }}
        <i class="fa fa-filter" :class="{ 'text-muted': !keyword}" style="cursor: pointer" @click="toggle = !toggle"></i>
        <ul class="dropdown-menu" :style="{'display': show, 'padding': '3px'}">
            <div class="input-group input-group-sm">
                <input type="search" class="form-control" v-focus
                    v-model="keyword" @keydown.enter="search" :placeholder="`Search ${field}...`">
                <span class="input-group-btn" style="min-width: 0">
                    <button class="btn btn-default fa fa-search" @click="search"></button>
                </span>
            </div>
        </ul>
    </div>
</template>
<script>
const focus = {
    inserted (el) {
        el.focus()
    }
}

export default {
    name: 'th-filter',
    props: ['field', 'title', 'query'],
    computed: {
        show () {
            return this.toggle ? 'block' : 'none'
        }
    },
    data: () => ({
        keyword: '',
        toggle: false
    }),
    watch: {
        keyword (kw) {
            // reset immediately if empty
            if (kw === '') this.search()
        }
    },
    directives: { focus },
    methods: {
        search () {
            const { query } = this
            // `$props.query` would be initialized to `{ limit: 10, offset: 0, sort: '', order: '' }` by default
            // custom query conditions must be set to observable by using `Vue.set / $vm.$set`
            this.$set(query, this.field, this.keyword)
            query.offset = 0 // reset pagination
            this.toggle = false
        }
    }
}
</script>
<style>
input[type=search]::-webkit-search-cancel-button {
    -webkit-appearance: searchfield-cancel-button;
    cursor: pointer;
}
</style>