import axios from 'axios'
import { put, call, delay, takeEvery, takeLatest, fork } from 'redux-saga/effects'
import { ActionTypes } from '../actions/SignUp'
import { ActionTypes as SignInActionTypes } from '../actions/SignIn'
import { baseURI } from '../common'

const url = new URL('/api/user', baseURI).href;

const API = {
    requestSignUp: (userdata) => {
        console.log(`CALL_API SIGINUP ID:${userdata.id} PW:${userdata.pw}`)
        return axios.post(url + '/signup', {
            id: userdata.id,
            password: userdata.pw
        })
            .then((result) => ({ data: result }))
            .catch((error) => ({ error }))
    },
}

function* requestSignUp(action) {
    const { data, error } = yield call(API.requestSignUp, action.data)

    console.log(data)
    if (data) {
        console.log(data)
        yield put({ type: SignInActionTypes.REQUEST_SIGNIN, data: { id: action.data.id, pw: action.data.pw } })
    } else if (error) {
        yield put({ type: ActionTypes.FAILED_SIGNUP, error: error.response.data })
    }
}

export default function* () {
    yield takeEvery(ActionTypes.REQUEST_SIGNUP, requestSignUp)
}