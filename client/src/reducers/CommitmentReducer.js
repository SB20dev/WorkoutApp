import { ActionTypes } from '../actions/Commitment';

const initialState = {
    total: 0,
    history: []
}

export default (state = initialState, action) => {
    switch(action.type) {
        // コミットメント概要
        case ActionTypes.SUCCESSED_FETCH_COMMITMENT_OVERVIEW:
            return {
                ...state,
                total: action.data
            }
        case ActionTypes.FAILED_FETCH_COMMITMENT_OVERVIEW:
            return {
                ...state,
                total: undefined
            }     
        // コミットメント履歴
        case ActionTypes.SUCCESSED_FETCH_HISTORY:
            return {
                ...state,
                history: action.data
            }
        case ActionTypes.FAILED_FETCH_HISTORY:
            return {
                ...state,
                history: []
            }
        default:
            return state
    }
}