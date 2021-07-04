import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'

import { createStore, applyMiddleware } from 'redux'
import CombinedReducer from './reducers/CombinedReducer'

import Root from './components/Root'
import createSagaMiddleware from 'redux-saga';
import rootSaga from './saga/RootSaga';

import './common.css'

const sagaMiddleWare = createSagaMiddleware()

const store = createStore(CombinedReducer, applyMiddleware(sagaMiddleWare))
sagaMiddleWare.run(rootSaga)

ReactDOM.render(
    <Provider store={store}>
        <Root history={history} />
    </Provider>,
    document.getElementById('root')
)