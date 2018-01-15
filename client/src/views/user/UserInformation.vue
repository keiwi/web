<template>
    <div class="user-table">
        <div class="btn-group">
            <button v-if="edit" class="btn btn-primary" @click="save">Save</button>
            <button v-else class="btn btn-success" @click="edit = true">Edit</button>
        </div>
        <table :class="tblClass" :style="tblStyle">
            <tr v-for="(column, index) in columns" :key="index">
                <th>{{column.title}}:</th>
                <td :style="tdStyle">
                    <template v-if="!edit || (edit && (column.edit == '' || column.edit == null))">{{userData != null ? userData[column.field] : ''}}</template>
                    <template v-if="edit && column.edit == 'dropdown'" v-for="(value, valueIndex) in column.values" >
                        <b-form-checkbox :key="valueIndex" :plain="true" v-model="value.Checked" style="margin-bottom: 0"> {{value.Name}}</b-form-checkbox>
                        <br :key="valueIndex">
                    </template>
                    <input v-if="edit && column.edit == 'input'" type="text" v-model="column.value">
                </td>
            </tr>
        </table>
    </div>
</template>

<script>
    import API from '../../api'
    import VueNotifications from 'vue-notifications'

    export default {
        props: ['user-data', 'columns'],
        name: 'user-information',
        methods: {
            async save () {
                this.edit = false
                let groups = ''
                let ip = ''
                let name = ''
                for (let c of this.columns) {
                    if (c.title === 'Groups') {
                        for (let g of c.values) {
                            if (g.Checked) {
                                groups += g.Name + ','
                            }
                        }
                    }
                    if (c.title === 'IP') {
                        ip = c.value
                    }
                    if (c.title === 'Namn') {
                        name = c.value
                    }
                }
                groups = groups.replace(/,\s*$/, '')

                if (groups !== this.userData.Groups) {
                    try {
                        await API.editClient({id: this.userData.ID, Option: 'group_names', Value: groups})
                        this.userData.Groups = groups
                    } catch (e) {
                        VueNotifications.error({message: e})
                        return
                    }
                }

                if (ip !== this.userData.IP) {
                    try {
                        await API.editClient({id: this.userData.ID, Option: 'ip', Value: ip})
                        this.userData.IP = ip
                    } catch (e) {
                        VueNotifications.error({message: e})
                        return
                    }
                }

                if (name !== this.userData.Name) {
                    try {
                        await API.editClient({id: this.userData.ID, Option: 'name', Value: name})
                        this.userData.Name = name
                    } catch (e) {
                        VueNotifications.error({message: e})
                        return
                    }
                }

                VueNotifications.info({message: 'Successfully saved the client'})
            }
        },
        data () {
            return {
                edit: false,
                tblClass: ['table-bordered', 'table'],
                tblStyle: {
                    whiteSpace: 'nowrap'
                },
                tdStyle: {
                    width: '100%'
                }
            }
        }
    }
</script>