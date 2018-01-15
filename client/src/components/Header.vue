<template>
    <header class="app-header navbar">
        <b-link class="navbar-brand" to="home"></b-link>
        <b-nav is-nav-bar class="d-md-down-none">
            <b-nav-item class="px-3"><router-link :to="{name: 'Home'}">Home</router-link></b-nav-item> 
            <b-nav-item class="px-3"><router-link :to="{name: 'Users'}">Users</router-link></b-nav-item>
            <b-nav-item class="px-3"><router-link :to="{name: 'Settings'}">Settings</router-link></b-nav-item>
            <b-nav-item class="px-3"><router-link :to="{name: 'Upload'}">Upload</router-link></b-nav-item>
        </b-nav>
        <b-nav is-nav-bar class="ml-auto">
            <b-nav-item class="d-md-down-none">
                <i class="icon-bell"></i><span class="badge badge-pill badge-danger">5</span>
            </b-nav-item>
            <b-nav-item-dropdown right>
                <template slot="button-content">
                    <img src="/static/img/avatars/6.jpg" class="img-avatar" alt="admin@bootstrapmaster.com">
                    <span class="d-md-down-none">admin</span>
                </template>
                <b-dropdown-header tag="div" class="text-center"><strong>Profile</strong></b-dropdown-header>
                <b-dropdown-item><i class="fa fa-user"></i> <router-link to="/profile/settings">Settings</router-link></b-dropdown-item>
                <b-dropdown-item><i class="fa fa-lock"></i> <router-link to="/logout">Logout</router-link></b-dropdown-item>
            </b-nav-item-dropdown>
        </b-nav>
        <button class="navbar-toggler aside-menu-toggler d-md-down-none" type="button" @click="asideToggle">&#9776;</button>
    </header>
</template>
<script>
export default {
    name: 'header',
    methods: {
        mobileSidebarToggle (e) {
            e.preventDefault()
            document.body.classList.toggle('sidebar-mobile-show')
        },
        asideToggle (e) {
            e.preventDefault()
            document.body.classList.toggle('aside-menu-hidden')
        }
    }
}
</script>
