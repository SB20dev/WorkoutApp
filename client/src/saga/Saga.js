import axios from 'axios'
import { fork, put, call, takeEvery } from 'redux-saga/effects'
import { ActionTypes } from '../actions/SamplePing'

const API = {
    url : 'http://localhost:8080/',
    ping : () => {
        return axios.get(API.url + "ping")
        .then((result) => ({ data: result.data }))
        .catch((error) => ({ error }))
    }
}

function* ping() {
    const { data, error } = yield call(API.ping)
    if (data) {
        yield put({ type: ActionTypes.SUCCESSED_PING, data})
    } else if (error) {
        console.log(error)
        yield put({ type: ActionTypes.FAILED_PING, error: error.toString()})
    }
}

function* handlePing() {
    yield takeEvery(ActionTypes.REQUEST_PING, ping)
}

export default function* rootSaga() {
    yield fork(handlePing);
}