export const ActionTypes = {
    REQUEST_SIGNUP: 'REQUEST_SIGNUP',
    SUCCESSED_SIGNUP: 'SUCCESSED_SIGNUP',
    FAILED_SIGNUP: 'FAILED_SIGNUP',
}

export const Actions = {
    requestSignUp: (id, pw) => ({
        type: ActionTypes.REQUEST_SIGNUP,
        data: {id, pw}
    }),
    successedSignUp: (data) => ({
        type: ActionTypes.SUCCESSED_SIGNUP,
        data
    }),
    failedSignUp: (error) => ({
        type: ActionTypes.FAILED_SIGNUP,
        error
    }),
}