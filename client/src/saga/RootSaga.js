import { fork } from 'redux-saga/effects'
import handleFetchCommitmentOverview from './CommitmentSaga'
import handleFetchSuggestion from './MenuSaga'
import handleRequestSignIn from './SignInSaga'
import handleRequestSignUp from './SignUpSaga'

export default function* rootSaga() {
    yield fork(handleFetchCommitmentOverview)
    yield fork(handleFetchSuggestion)
    yield fork(handleRequestSignIn)
    yield fork(handleRequestSignUp)
}