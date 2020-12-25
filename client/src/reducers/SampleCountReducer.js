import { ActionTypes } from '../actions/SampleCount';

// これはここに書いておくべきかどうか
const initialState = {
    number: 10
}

export default (state = initialState, action) => {
    switch(action.type) {
        case ActionTypes.COUNT_DOWN:
            return {
                ...state,
                number: state.number - action.payload.delta
            }
        case ActionTypes.COUNT_UP:
            return {
                ...state,
                number: state.number + action.payload.delta
            }
        default:
            return state
    }
}