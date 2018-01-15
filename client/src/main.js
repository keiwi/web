// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import App from './App'
import router from './router/index'
import VueResource from 'vue-resource'
import store from './store'
import VueNotifications from 'vue-notifications'
import Datatable from 'vue2-datatable-component'
import AsyncComputed from 'vue-async-computed'
import VueHighlightJS from 'vue-highlightjs'
import { options as VueNotificationsOptions } from './notifications'
import { mapActions } from 'vuex'

Vue.use(VueResource)
Vue.use(BootstrapVue)
Vue.use(VueNotifications, VueNotificationsOptions)
Vue.use(Datatable)
Vue.use(AsyncComputed)
Vue.use(VueHighlightJS)

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    store,
    template: '<App/>',
    components: {
        App
    },
    methods: mapActions(['initCommands', 'initGroups', 'initClients']),
    async created () {
        try {
            await this.initCommands()
        } catch (e) {
            VueNotifications.error({message: e})
        }

        try {
            await this.initGroups()
        } catch (e) {
            VueNotifications.error({message: e})
        }

        try {
            await this.initClients()
        } catch (e) {
            VueNotifications.error({message: e})
        }
    }
})
