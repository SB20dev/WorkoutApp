import { fork } from 'redux-saga/effects'
import handleFetchCommitmentOverview from './CommitmentSaga'
import handleFetchSuggestion from './MenuSaga'

export default function* rootSaga() {
    yield fork(handleFetchCommitmentOverview)
    yield fork(handleFetchSuggestion)
}