export const ActionTypes = {
    REQUEST_PING: 'REQUEST_PING',
    SUCCESSED_PING: 'SUCCESSED_PING',
    FAILED_PING: 'FAILED_PING'
}

export const Actions = {
    requestPing: () => ({
        type: ActionTypes.REQUEST_PING
    }),
    successedPing: (data) => ({
        type: ActionTypes.SUCCESSED_PING,
        data
    }),
    failedPing: (error) => ({
        type: ActionTypes.FAILED_PING,
        error
    })
}