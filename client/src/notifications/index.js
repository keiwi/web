import miniToastr from 'mini-toastr'

// miniToastr config https://github.com/se-panfilov/mini-toastr#global-config
export const toastConfig = {
    types: {
        success: 'success',
        error: 'error',
        info: 'info',
        warn: 'warn'
    }
}

miniToastr.init(toastConfig)

// Here we setup messages output to `mini-toastr`
function toast ({title, message, type, timeout, cb}) {
    return miniToastr[type](message, title, timeout, cb)
}

// Binding for methods .success(), .error() and etc. You can specify and map your own methods here.
// Required to pipe our output to UI library (mini-toastr in example here)
// All not-specified events (types) would be piped to output in console.
export const options = {
    success: toast,
    error: toast,
    info: toast,
    warn: toast
}
