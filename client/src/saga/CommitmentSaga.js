import axios from 'axios'
import { put, call, delay, takeEvery } from 'redux-saga/effects'
import { ActionTypes } from '../actions/Commitment'

const API = {
    url : 'http://localhost:8080/api/commitment/',
    fetchOverview : () => {
        // 本来はこんな実装
        // return axios.get(API.url + "overview")
        // .then((result) => ({ data: result.data }))
        // .catch((error) => ({ error }))
        return Promise.resolve({ data: 100})
    },
    fetchHistory : () => {
        const changeDate = (date, diff) => {
            date.setDate(date.getDate() + diff)
            return date
        }
        return Promise.resolve({
            data: [
                { date: new Date(), menus: [{}, {}, {}], score: 3 }, // 今のところこんなオブジェクトを想定
                { date: changeDate(new Date(), -1), menus: [], score: 2 },
                { date: changeDate(new Date(), -2), menus: [], score: 1 },
                { date: changeDate(new Date(), -3), menus: [], score: -1 },
                { date: changeDate(new Date(), -4), menus: [], score: 2 },
            ]
        })
    }
}

function* fetchCommitmentOverview() {
    yield delay(1000)
    const { data, error } = yield call(API.fetchOverview)
    if (data) {
        yield put({ type: ActionTypes.SUCCESSED_FETCH_COMMITMENT_OVERVIEW, data })
    } else if (error) {
        yield put({ type: ActionTypes.FAILED_FETCH_COMMITMENT_OVERVIEW, error: error.toString() })
    }
}

function* fetchHistory() {
    yield delay(1500)
    const { data, error } = yield call(API.fetchHistory)
    if (data) {
        yield put({ type: ActionTypes.SUCCESSED_FETCH_HISTORY, data })
    } else if (error) {
        yield put({ type: ActionTypes.FAILED_FETCH_HISTORY, error: error.toString() })
    }
}

export default function*() {
    yield takeEvery(ActionTypes.REQUEST_FETCH_COMMITMENT_OVERVIEW, fetchCommitmentOverview)
    yield takeEvery(ActionTypes.REQUEST_FETCH_HISTORY, fetchHistory)
}