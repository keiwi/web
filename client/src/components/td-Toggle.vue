<template>
    <label class="switch-btn">
        <input class="checked-switch" type="checkbox" v-model="value" />
        <span class="text-switch" data-yes="yes" data-no="no"></span>
        <span class="toggle-btn" :style="{left: checkedWidth}"></span>
    </label>
</template>
<script>
    export default {
        props: ['row', 'nested', 'field'],
        computed: {
            value: {
                get () { return this.val || this.row[this.field] },
                set (newValue) {
                    this.val = newValue
                    this.row[this.field] = newValue
                }
            },
            checkedWidth () {
                if (this.value) {
                    return this.width
                }
                return ''
            }
        },
        mounted () {
            this.width = (this.$el.offsetWidth - 29) + 'px'

            let self = this
            window.addEventListener('resize', function () {
                self.width = (self.$el.offsetWidth - 29) + 'px'
            })
        },
        data () {
            return {
                val: '',
                width: ''
            }
        },
        methods: {
        }
}
</script>
<style>
    .switch-btn {
        position: relative; 
        display: block; 
        vertical-align: top; 
        width: 100%; 
        height: 30px; 
        border-radius: 18px; 
        cursor: pointer;
    }

    .checked-switch {
        position: absolute; 
        top: 0; 
        left: 0; 
        opacity: 0;
    }

    .text-switch {
        background-color: #ed5b49; 
        border: 1px solid #d2402e; 
        border-radius: inherit; 
        color: #fff; 
        display: block; 
        font-size: 15px; 
        height: inherit; 
        position: relative; 
        text-transform: uppercase;
    }

    .text-switch:before, 
    .text-switch:after {
        position: absolute; 
        top: 50%; 
        margin-top: -.5em; 
        line-height: 1; 
        -webkit-transition: inherit; 
        -moz-transition: inherit; 
        -o-transition: inherit; 
        transition: inherit;
    }

    .text-switch:before {
        content: attr(data-no); 
        right: 11px;
    }

    .text-switch:after {
        content: attr(data-yes); 
        left: 11px; 
        color: #FFFFFF; 
        opacity: 0;
    }

    .checked-switch:checked ~ .text-switch {
        background-color: #00af2c; 
        border: 1px solid #068506;
    }

    .checked-switch:checked ~ .text-switch:before {
        opacity: 0;
    }

    .checked-switch:checked ~ .text-switch:after {
        opacity: 1;
    }

    .toggle-btn {
        background: linear-gradient(#eee, #fafafa); 
        border-radius: 100%; 
        height: 28px; 
        left: 1px; 
        position: absolute; 
        top: 1px; 
        width: 28px;
        font: normal normal normal 20px/1 FontAwesome;
        text-align: center;
    }

    .toggle-btn::before {
        color: #aaaaaa;
        content: "\f00d"; 
        display: inline-block;
        padding: 4px 0; 
        vertical-align: middle;
    }

    .checked-switch:checked ~ .toggle-btn::before {
        content: "\f00c";
    }

    .text-switch:hover ~ .toggle-btn::before,
    .toggle-btn:hover::before {
        color: #ed5b49;
    }

    .checked-switch:checked ~ .text-switch:hover ~ .toggle-btn::before,
    .checked-switch:checked ~ .toggle-btn:hover::before {
        color: #00af2c;
    }

    .text-switch, .toggle-btn {
        transition: All 0.3s ease; 
        -webkit-transition: All 0.3s ease; 
        -moz-transition: All 0.3s ease;
        -o-transition: All 0.3s ease;
    }
</style>