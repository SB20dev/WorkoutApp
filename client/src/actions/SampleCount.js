export const ActionTypes = {
    COUNT_DOWN: "COUNT_DOWN",
    COUNT_UP: "COUNT_UP"
}

export const Actions = {
    countDown: (delta) => ({
        type: ActionTypes.COUNT_DOWN,
        payload: {
            delta
        }
    }),
    countUp: (delta) => ({
        type: ActionTypes.COUNT_UP,
        payload: {
            delta
        }
    })
}