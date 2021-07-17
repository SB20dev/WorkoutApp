import { ActionTypes } from '../actions/SignUp';

const initialState = {
    success_signup: undefined,
    failed_signup_reason: undefined,
}

export default (state = initialState, action) => {

    switch (action.type) {
        case ActionTypes.SUCCESSED_SIGNUP:
            return {
                ...state,
                success_signup: true,
                failed_signup_reason: undefined,
            }
        case ActionTypes.FAILED_SIGNUP:
            return {
                ...state,
                success_signup: false,
                failed_signup_reason: action.error,
            }
        default:
            return state
    }
}