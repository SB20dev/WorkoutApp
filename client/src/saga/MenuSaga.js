import axios from 'axios'
import { put, call, delay, takeEvery } from 'redux-saga/effects'
import { ActionTypes } from '../actions/Menu'

const API = {
    url : 'http://localhost:8080/api/menu/',
    fetchSuggestion : () => {
        // 本来はこんな実装
        // return axios.get(API.url + "suggestion")
        // .then((result) => ({ data: result.data }))
        // .catch((error) => ({ error }))
        return Promise.resolve({ data: ["ダンベルカール","フレンチフライ","バイシクルクランチ"] })
    }
}

function* fetchSuggestion() {
    yield delay(1000) // 擬似的にネットワーク遅延を再現
    const { data, error } = yield call(API.fetchSuggestion)
    if (data) {
        yield put({ type: ActionTypes.SUCCESSED_FETCH_SUGGESTION, data })
    } else if (error) {
        yield put({ type: ActionTypes.FAILED_FETCH_SUGGESTION, error: error.toString() })
    }
}

export default function*() {
    yield takeEvery(ActionTypes.REQUEST_FETCH_SUGGESTION, fetchSuggestion)
}