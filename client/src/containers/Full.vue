<template>
    <div class="app">
        <AppHeader/>
        <div class="app-body">
            <main class="main">
                <div class="container-fluid">
                    <div class="row">
                        <div :class="left"></div>
                        <div :class="middle">
                            <router-view></router-view>
                        </div>
                        <div :class="right"></div>
                    </div>
                </div>
            </main>
            <AppAside/>
        </div>
        <AppFooter/>
        <AppModal />
    </div>
</template>

<script>
import pageSettings from '../_pageSettings.js'
import { Header as AppHeader, Aside as AppAside, Footer as AppFooter, Modal as AppModal } from '../components/'

export default {
    name: 'full',
    components: {
        AppHeader,
        AppAside,
        AppFooter,
        AppModal
    },
    data () {
        return {
        }
    },
    computed: {
        left () {
            return 'col-md-' + ((12 - this.size) / 2)
        },
        middle () {
            return 'col-md-' + this.size
        },
        right () {
            return 'col-md-' + ((12 - this.size) / 2)
        },
        size () {
            let size = (typeof pageSettings[this.name] === 'undefined')
                ? pageSettings.Default.Size
                : pageSettings[this.name].Size
            return (size % 2 === 0) ? size : size + 1
        },
        name () {
            return this.$route.name
        },
        list () {
            return this.$route.matched
        }
    }
}
</script>
