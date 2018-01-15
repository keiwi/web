<template>
    <div class="animated fadeIn">
        <b-card>
            <h4 class="card-title mb-0">Traffic</h4>
            <div slot="footer">
                <div class="row">
                    <settings-commands></settings-commands>
                </div>
                <hr />
                <div class="row">
                    <settings-groups></settings-groups>
                </div>
            </div>
        </b-card>
    </div>
</template>

<script>
import SettingsGroups from './settings/SettingsGroups'
import SettingsCommands from './settings/SettingsCommands'

export default {
    name: 'dashboard',
    components: {
        SettingsGroups,
        SettingsCommands
    }
}
</script>