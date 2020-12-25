import React from 'react'
import { Actions } from '../../actions/SamplePing'
import { connect } from 'react-redux'

const SamplePing = props => (
    <div className="sample-ping">
        <h2>sample ping</h2>
        <input type='button' value='ping' onClick={() => {props.onPingBtnClicked()}} />
        {props.pong}
    </div>
)

const mapStateToProps = state => ({
    pong: state.ping.pong
})

const mapDispatchToProps = dispatch => ({
    onPingBtnClicked: () => { dispatch(Actions.requestPing()) }
})

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(SamplePing)