export const ActionTypes = {
    // コミットメント概要
    REQUEST_FETCH_COMMITMENT_OVERVIEW: 'REQUEST_FETCH_COMMITMENT_OVERVIEW',
    SUCCESSED_FETCH_COMMITMENT_OVERVIEW: 'SUCCESSED_FETCH_COMMITMENT_OVERVIEW',
    FAILED_FETCH_COMMITMENT_OVERVIEW: 'FAILED_FETCH_COMMITMENT_OVERVIEW',
    // コミットメント履歴
    REQUEST_FETCH_HISTORY: 'REQUEST_FETCH_HISTORY',
    SUCCESSED_FETCH_HISTORY: 'SUCCESSED_FETCH_HISTORY',
    FAILED_FETCH_HISTORY: 'FAILED_FETCH_HISTORY'
}

export const Actions = {
    // コミットメント概要
    requestFetchCommitmentOverview: () => ({
        type: ActionTypes.REQUEST_FETCH_COMMITMENT_OVERVIEW
    }),
    successedFetchCommitmentOverview: (data) => ({
        type: ActionTypes.SUCCESSED_FETCH_COMMITMENT_OVERVIEW,
        data
    }),
    failedFetchCommitmentOverview: (error) => ({
        type: ActionTypes.FAILED_FETCH_COMMITMENT_OVERVIEW,
        error
    }),
    // コミットメント履歴
    requestFetchHistory: (begin, num) => ({
        type: ActionTypes.REQUEST_FETCH_HISTORY,
        data: { begin, num }
    }),
    successedFetchHistory: (data) => ({
        type: ActionTypes.SUCCESSED_FETCH_HISTORY,
        data
    }),
    failedFetchHistory: (error) => ({
        type: ActionTypes.FAILED_FETCH_HISTORY,
        error
    })
}