import { combineReducers } from 'redux'
import CommitmentReducer from './CommitmentReducer'
import MenuReducer from './MenuReducer'

export default combineReducers({
    commitment: CommitmentReducer,
    menu: MenuReducer
})