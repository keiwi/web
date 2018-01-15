<template>
    <div>
        <span v-if="!edit" @click="focus">{{ value }}</span>
        <input
            type="input"
            v-model="modelValue"
            class="form-control"
            @keyup.enter="edit = false"
            @blur="done"
            v-if="edit"
            ref="input"
            v-focus>
    </div>
</template>

<script>
import VueNotifications from 'vue-notifications'

export default {
    name: 'textbox-table',
    props: ['value', 'action', 'id', 'option'],
    data: function () {
        return {
            edit: false,
            modelValue: this.value
        }
    },
    directives: { focus },
    methods: {
        focus: function () {
            this.edit = true
        },
        done: function (e) {
            this.action({id: this.id, option: this.option, value: this.modelValue})
                .then((response) => {
                    if (response.success) VueNotifications.info({message: response.message})
                    else VueNotifications.error({message: response.message})
                }, (response) => VueNotifications.error({message: response}))
            this.edit = false
        }
    }
}
</script>