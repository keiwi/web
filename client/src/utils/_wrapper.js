import VueNotifications from 'vue-notifications'

export function notify (func, payload) {
    func(payload).then((response) => {
        if (response.success) VueNotifications.info({message: response.message})
        else VueNotifications.error({message: response.message})
    }, (response) => VueNotifications.error({message: response}))
}
