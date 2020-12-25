import { combineReducers } from 'redux'
import SampleCountReducer from './SampleCountReducer'
import PingReducer from './SamplePingReducer'

export default combineReducers({
    count: SampleCountReducer,
    ping: PingReducer
})