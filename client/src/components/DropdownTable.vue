<template>
    <div>
        <b-form-select
            :options="dropdownOptions"
            v-model="sel"
            required
            @input="done"
        ></b-form-select>
    </div>
</template>

<script>
import VueNotifications from 'vue-notifications'

export default {
    name: 'dropdown-table',
    props: ['options', 'action', 'id', 'option', 'selected'],
    data: function () {
        return {
            dropdownOptions: this.options,
            sel: this.selected
        }
    },
    methods: {
        done: function () {
            this.action({id: this.id, option: this.option, value: this.sel})
                .then((response) => {
                    if (response.success) VueNotifications.info({message: response.message})
                    else VueNotifications.error({message: response.message})
                }, (response) => VueNotifications.error({message: response}))
        }
    }
}
</script>