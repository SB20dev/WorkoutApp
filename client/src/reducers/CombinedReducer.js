import { combineReducers } from 'redux'
import CommitmentReducer from './CommitmentReducer'
import MenuReducer from './MenuReducer'
import SignInReducer from './SignInReducer'
import SignUpReducer from './SignUpReducer'

export default combineReducers({
    commitment: CommitmentReducer,
    menu: MenuReducer,
    signUp: SignUpReducer,
    signIn: SignInReducer,
})