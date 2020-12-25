import { ActionTypes } from '../actions/SamplePing';

const initialState = {
    pong: '...'
}

export default (state = initialState, action) => {
    switch(action.type) {
        case ActionTypes.SUCCESSED_PING:
            return {
                ...state,
                pong: action.data
            }
        case ActionTypes.FAILED_PING:
            return {
                ...state,
                pong: action.error
            }
        default:
            return state
    }
}