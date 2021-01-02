import { ActionTypes } from '../actions/Menu';

const initialState = {
    suggestion: []
}

export default (state = initialState, action) => {
    switch(action.type) {
        case ActionTypes.SUCCESSED_FETCH_SUGGESTION:
            return {
                ...state,
                suggestion: action.data
            }
        case ActionTypes.FAILED_FETCH_SUGGESTION:
            return {
                ...state,
                suggestion: []
            }
        default:
            return state
    }
}