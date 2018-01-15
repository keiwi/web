<template>
    <div class="timepicker">
        <div class="timepicker-col hours">
            <input
                type="timepicker"
                v-model="fixedHours"
                @click="setFocus"
                ref="hours"
                name="hours"
                @keydown.up.prevent="step('up')"
                @keydown.down.prevent="step('down')"
                @keydown.left.prevent="switchFocus('left', $event)"
                @keydown.right.prevent="switchFocus('right', $event)"
                @keydown.enter.prevent="switchFocus('', $event)"
                @keydown="changeValue"
                @keyup="correctTime" />
        </div>:
        <div class="timepicker-col minutes">
            <input
                type="timepicker"
                v-model="fixedMinutes"
                @click="setFocus"
                ref="minutes"
                name="minutes"
                @keydown.up.prevent="step('up')"
                @keydown.down.prevent="step('down')"
                @keydown.left.prevent="switchFocus('left', $event)"
                @keydown.right.prevent="switchFocus('right', $event)"
                @keydown.enter.prevent="switchFocus('', $event)"
                @keydown="changeValue"
                @keyup="correctTime"  />
        </div>:
        <div class="timepicker-col seconds">
            <input
                type="timepicker"
                v-model="fixedSeconds"
                @click="setFocus"
                ref="seconds"
                name="seconds"
                @keydown.up.prevent="step('up')"
                @keydown.down.prevent="step('down')"
                @keydown.left.prevent="switchFocus('left', $event)"
                @keydown.right.prevent="switchFocus('right', $event)"
                @keydown.enter.prevent="switchFocus('', $event)"
                @keydown="changeValue"
                @keyup="correctTime"  />
        </div>
        <div class="timepicker-col actions">
            <div class="arrow"><div class="arrow-up" @mousedown.prevent="stepArrow('up')" @mouseup.prevent="resetTimer"></div></div>
            <div class="arrow"><div class="arrow-down" @mousedown.prevent="stepArrow('down')" @mouseup.prevent="resetTimer"></div></div>
        </div>
    </div>
</template>
<script>
    function sleep (ms) {
        return new Promise(resolve => setTimeout(resolve, ms))
    }
    let arrowTimer = false

    export default {
        props: ['row', 'nested', 'field'],
        mounted () {
            let value = this.row[this.field]
            this.value = value

            this.hoursTime = Math.floor(value / 3600)
            value = value - this.hoursTime * 3600

            this.minutesTime = Math.floor(value / 60)
            this.secondsTime = value - this.minutesTime * 60
        },
        computed: {
            fixedSeconds: {
                get () {
                    return this.secondsTime.toString().padStart(2, '0')
                },
                set (newValue) {
                    this.secondsTime = parseInt(newValue)
                }
            },
            fixedMinutes: {
                get () {
                    return this.minutesTime.toString().padStart(2, '0')
                },
                set (newValue) {
                    this.minutesTime = parseInt(newValue)
                }
            },
            fixedHours: {
                get () {
                    return this.hoursTime.toString().padStart(2, '0')
                },
                set (newValue) {
                    this.hoursTime = parseInt(newValue)
                }
            }
        },
        data () {
            return {
                secondsTime: 0,
                secondsStep: 5,
                secondsMax: 59,
                minutesTime: 0,
                minutesStep: 5,
                minutesMax: 59,
                hoursTime: 0,
                hoursStep: 1,
                value: 0,
                focused: null
            }
        },
        methods: {
            outputNumber (num) {
                return (num).toString().padStart(2, '0')
            },
            step (dir, e) {
                if (this.focused === null) {
                    this.focused = this.$refs.seconds
                }
                this.focused.focus()

                let name = this.focused.getAttribute('name')
                if (dir === 'up') {
                    this[name + 'Time'] += this[name + 'Step']
                } else {
                    this[name + 'Time'] -= this[name + 'Step']
                }

                this.correctTime()
            },
            async stepArrow (dir) {
                arrowTimer = true
                let wait = 1000
                /* eslint-disable no-unmodified-loop-condition */
                while (arrowTimer) {
                    this.step(dir)
                    if (wait > 100) wait -= 100
                    await sleep(wait)
                }
                /* eslint-enable no-unmodified-loop-condition */
            },
            resetTimer () {
                arrowTimer = false
            },
            correctTime (e) {
                // When increasing, lowest (seconds) need to be first.
                // TODO: Maybe switch focus when going above max/min
                this.hoursTime = parseInt(this.hoursTime) || 0
                this.minutesTime = parseInt(this.minutesTime) || 0
                this.secondsTime = parseInt(this.secondsTime) || 0

                // Handle when seconds go above max
                while (this.secondsMax && this.secondsTime > this.secondsMax) {
                    this.secondsTime = this.secondsTime - (this.secondsMax + 1)
                    this.minutesTime++
                }

                // Handle when minutes go above max
                while (this.minutesMax && this.minutesTime > this.minutesMax) {
                    this.minutesTime = this.minutesTime - (this.minutesMax + 1)
                    this.hoursTime++
                }

                // Handle when hours go above max
                while (this.hoursMax && this.hoursTime > this.hoursMax) {
                    this.hoursTime = this.hoursMax
                }

                // Handle when seconds go below 0
                if (this.hoursTime < 0) this.hoursTime = 0
                // Handle when seconds go below 0
                if (this.minutesTime < 0) this.minutesTime = 0
                // Handle when seconds go below 0
                if (this.secondsTime < 0) this.secondsTime = 0

                this.value = (this.hoursTime * 60 * 60) + (this.minutesTime * 60) + this.secondsTime
                this.row[this.field] = this.value
            },
            setFocus (e) {
                this.focused = e.target
            },
            changeValue (e) {
                if (
                    e.key !== 'Backspace' &&
                    e.key !== 'Delete' &&
                    e.key !== 'ArrowUp' &&
                    e.key !== 'ArrowDown' &&
                    e.key !== 'ArrowLeft' &&
                    e.key !== 'ArrowRight' &&
                    isNaN(e.key)) {
                    e.preventDefault()
                }
            },
            switchFocus (dir, e) {
                let name = e.target.getAttribute('name')

                switch (dir) {
                // Left arrow clicked
                case 'left':
                    switch (name) {
                    case 'hours':
                        this.$refs.seconds.focus()
                        break
                    case 'minutes':
                        this.$refs.hours.focus()
                        break
                    case 'seconds':
                        this.$refs.minutes.focus()
                        break
                    }
                    break

                // Right arrow clicked
                case 'right':
                    switch (name) {
                    case 'hours':
                        this.$refs.minutes.focus()
                        break
                    case 'minutes':
                        this.$refs.seconds.focus()
                        break
                    case 'seconds':
                        this.$refs.hours.focus()
                        break
                    }
                    break

                // Enter clicked
                default:
                    switch (name) {
                    case 'hours':
                        this.$refs.minutes.focus()
                        break
                    case 'minutes':
                        this.$refs.seconds.focus()
                        break
                    }
                    break
                }
            }
        }
    }
</script>
<style>
    .timepicker {
        background-color: white;
        user-select: text;
        cursor: auto;
        border-width: 2px;
        border-style: inset;
        border-color: initial;
        border-image: initial;
        -webkit-appearance: listbox;
        display: flex;
        width: 100%;
        justify-content: flex-end;
    }

    .timepicker-col {
        margin: 0 3px;
    }

    .timepicker input {
        /* Reset input */
        margin: 0;
        border: 0;
        padding: 0;
        display: inline-block;
        vertical-align: middle;
        white-space: normal;
        background: none;
        line-height: 1;
	    box-sizing: content-box;
        
        /* Browsers have different default form fonts */
        font-size: inherit;
        font-family: inherit;

        /* Style input */
        width:1vw;
	    height:100%;
        text-align: right;
    }
    
    .timepicker input:hover {
        outline-color: rgb(77, 144, 254);
        outline-offset: 2px;
        outline-style: auto;
        outline-width: 5px;
    }

    .timepicker input:focus {
        outline-color: rgb(77, 144, 254);
        outline-offset: 2px;
        outline-style: auto;
        outline-width: 5px;
    }

    .hours {
        flex-grow: 1;
    }

    .hours input {
        width: 100%;
    }

    .actions {
        display: inline-flex;
        flex-direction: column;
        height: 100%;
        position: relative;
    }

    .arrow {
        padding: 2px;
    }
    .arrow-up {
        width: 0; 
        height: 0; 
        border-left: 1vh solid transparent;
        border-right: 1vh solid transparent;
        border-bottom: 1vh solid black;
    }

    .arrow-down {
        width: 0; 
        height: 0; 
        border-left: 1vh solid transparent;
        border-right: 1vh solid transparent;
        border-top: 1vh solid black;
    }

    .arrow:hover {
        background: #bdbdbd;
    }
</style>