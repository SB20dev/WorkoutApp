export const ActionTypes = {
    REQUEST_SIGNIN: 'REQUEST_SIGNIN',
    SUCCESSED_SIGNIN: 'SUCCESSED_SIGNIN',
    FAILED_SIGNIN: 'FAILED_SIGNIN',

    REQUEST_SIGNOUT: 'REQUEST_SIGNOUT',
}

export const Actions = {
    requestSignIn: (id, pw) => ({
        type: ActionTypes.REQUEST_SIGNIN,
        data: {id, pw}
    }),
    successedSignIn: (data) => ({
        type: ActionTypes.SUCCESSED_SIGNIN,
        data
    }),
    failedSignIn: (error) => ({
        type: ActionTypes.FAILED_SIGNIN,
        error
    }),

    requestSignOut: () => ({
        type: ActionTypes.REQUEST_SIGNOUT
    })
}