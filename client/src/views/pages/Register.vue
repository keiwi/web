<template>
    <div class="app flex-row align-items-center">
        <div class="container">
            <div class="row justify-content-center">
                <div class="col-md-6">
                    <div class="card mx-4">
                        <div class="card-body p-4">
                            <h1>Register</h1>
                            <p class="text-muted">Create your account</p>
                            <div class="input-group mb-3">
                                <span class="input-group-addon"><i class="icon-user"></i></span>
                                <input
                                    type="text"
                                    class="form-control"
                                    placeholder="Username"
                                    v-model="credentials.username"
                                    v-on:keyup.enter="submit"
                                >
                            </div>

                            <div class="input-group mb-3">
                                <span class="input-group-addon">@</span>
                                <input
                                    type="text"
                                    class="form-control"
                                    placeholder="Email"
                                    v-model="credentials.email"
                                    v-on:keyup.enter="submit"
                                >
                            </div>

                            <div class="input-group mb-3">
                                <span class="input-group-addon"><i class="icon-lock"></i></span>
                                <input
                                    type="password"
                                    class="form-control"
                                    placeholder="Password"
                                    v-model="credentials.password"
                                    v-on:keyup.enter="submit"
                                >
                            </div>

                            <div class="input-group mb-4">
                                <span class="input-group-addon"><i class="icon-lock"></i></span>
                                <input
                                    type="password"
                                    class="form-control"
                                    placeholder="Repeat password"
                                    v-model="credentials.cpassword"
                                    v-on:keyup.enter="submit"
                                >
                            </div>

                            <button type="button" class="btn btn-block btn-success" @click="submit">Create Account</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import auth from '../../auth'
    export default {
        name: 'Register',
        data () {
            return {
                credentials: {
                    username: '',
                    email: '',
                    password: '',
                    cpassword: ''
                },
                error: ''
            }
        },
        methods: {
            submit () {
                if (this.credentials.password !== this.credentials.cpassword) {
                    this.error = 'Your password is not correct.' // TODO: Better error message
                    return
                }

                var credentials = {
                    username: this.credentials.username,
                    email: this.credentials.email,
                    password: this.credentials.password
                }
                auth.signup(this, credentials, 'home')
            }
        }
    }
</script>
