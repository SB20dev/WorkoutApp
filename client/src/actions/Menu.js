export const ActionTypes = {
    REQUEST_FETCH_SUGGESTION: 'REQUEST_FETCH_SUGGESTION',
    SUCCESSED_FETCH_SUGGESTION: 'SUCCESSED_FETCH_SUGGESTION',
    FAILED_FETCH_SUGGESTION: 'FAILED_FETCH_SUGGESTION'
}

export const Actions = {
    requestFetchSuggestion: () => ({
        type: ActionTypes.REQUEST_FETCH_SUGGESTION
    }),
    successedFetchSuggestion: (data) => ({
        type: ActionTypes.SUCCESSED_FETCH_SUGGESTION,
        data
    }),
    failedFetchSuggestion: (error) => ({
        type: ActionTypes.FAILED_FETCH_SUGGESTION,
        error
    })
}