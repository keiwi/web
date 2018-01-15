const API_URL = '/api/v1/user'
const LOGIN_URL = API_URL + '/login'
const SIGNUP_URL = API_URL + '/signup'

export default {
    login (context, creds, redirect) {
        context.$http.post(LOGIN_URL, creds).then(response => {
            localStorage.setItem('id_token', response.body.id_token)
            if (redirect) {
                context.$router.replace(redirect)
            }
        }, response => {
            context.error = response.statusText
        })
    },

    signup (context, creds, redirect) {
        console.log(context)
        context.$http.post(SIGNUP_URL, creds).then(response => {
            localStorage.setItem('id_token', response.body.id_token)
            if (redirect) {
                context.$router.replace(redirect)
            }
        }, response => {
            context.error = response.statusText
        })
    },

    logout (context) {
        localStorage.removeItem('id_token')
        context.$router.replace('/home')
    },

    isAuthenticated () {
        var jwt = localStorage.getItem('id_token')
        if (jwt) {
            return true
        }
        return false
    },

    getAuthHeader () {
        return {
            'Authorization': 'Bearer ' + localStorage.getItem('id_token')
        }
    }
}
