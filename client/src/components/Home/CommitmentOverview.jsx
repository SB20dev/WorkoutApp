import React from 'react'
import { Actions } from '../../actions/Commitment'
import { connect } from 'react-redux'

class CommitmentOverview extends React.Component {
    constructor(props) {
        super(props)
        props.onConstruct()
    }

    render() {
        return (
            <div className="commitment-overview">
                <h2>total commitment</h2>
                {this.props.total}
            </div>
        )
    }
}

const mapStateToProps = state => ({
    total: state.commitment.total
})

const mapDispatchToProps = dispatch => ({
    onConstruct: () => { dispatch(Actions.requestFetchCommitmentOverview()) }
})

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(CommitmentOverview)