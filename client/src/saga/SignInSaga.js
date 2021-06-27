import axios from 'axios'
import { put, call, delay, takeEvery, takeLatest, fork } from 'redux-saga/effects'
import { ActionTypes } from '../actions/SignIn'
import { baseURI } from '../common'

const url = new URL('/api/user', baseURI).href;

const API = {
    requestSignIn: (userdata) => {
        console.log(`CALL_API SIGNIN ID:${userdata.id} PW:${userdata.pw}`)
        return axios.post(url + '/signin', {
            id: userdata.id,
            password: userdata.pw
        })
            .then((result) => ({ data: result }))
            .catch((error) => ({ error }))
    }
}

function* requestSignIn(action) {
    const { data, error } = yield call(API.requestSignIn, action.data)

    console.log(data)
    if (data) {
        yield put({ type: ActionTypes.SUCCESSED_SIGNIN, data })
    } else if (error) {
        yield put({ type: ActionTypes.FAILED_SIGNIN, error: error.response.data })
    }
}

export default function* () {
    yield takeEvery(ActionTypes.REQUEST_SIGNIN, requestSignIn)
}