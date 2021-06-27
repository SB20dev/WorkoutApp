import { ActionTypes } from '../actions/SignIn';

const initialState = {
    success_signin: undefined,
    failed_signin_reason: undefined,
    success_signin: undefined,
    user_token: undefined
}

export default (state = initialState, action) => {
    console.log(action);
    switch (action.type) {
        case ActionTypes.SUCCESSED_SIGNIN:
            document.cookie = 'token=' + action.data.data.token
            return {
                ...state,
                success_signin: true,
                failed_signin_reason: undefined,
            }
        case ActionTypes.FAILED_SIGNIN:
            return {
                ...state,
                success_signin: false,
                failed_signin_reason: action.error,
            }
        case ActionTypes.REQUEST_SIGNOUT:
            document.cookie = ''
            return {
                success_signin: false,
            }
        default:
            return state
    }
}