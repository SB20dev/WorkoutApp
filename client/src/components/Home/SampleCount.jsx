import React from 'react'
import { connect } from 'react-redux'
import { Actions } from '../../actions/SampleCount'

const SampleCount = props => (
    <div className="sample-count">
        <h2>sample count</h2>
        <a href="" onClick={() => {props.countDown()}}>-</a>
        {props.count}
        <a href="" onClick={() => {props.countUp()}}>+</a>
    </div>
)

const mapStateToProps = state => ({
    count: state.count.number
})

const mapDispatchToProps = dispatch => ({
    countDown: () => { dispatch(Actions.countDown(1))},
    countUp: () => { dispatch(Actions.countUp(1))}
})

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(SampleCount)