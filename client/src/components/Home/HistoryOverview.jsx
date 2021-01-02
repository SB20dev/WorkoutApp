import React from 'react'
import { Actions } from '../../actions/Commitment'
import { connect } from 'react-redux'

class HistoryOverview extends React.Component {
    constructor(props) {
        super(props)
        props.onConstruct()
    }

    render() {
        const formatDate = (rawDate) => {
            const weekday = ["日", "月", "火", "水", "木", "金", "土"]
            const year = rawDate.getYear() + 1900, month = rawDate.getMonth() + 1, date = rawDate.getDate(), day = weekday[rawDate.getDay()]
            return `${year}/${month}/${date}(${day})`
        }

        const historyOverview = this.props.history.map((commitment, idx) =>
            <li key={"history-overview-item"+idx}>
                {formatDate(commitment.date) + "\t" + commitment.score}
            </li>
        )

        return (
            <div className="history-overview">
                <h2>トレーニング履歴</h2>
                {historyOverview}
            </div>
        )
    }
}

const mapStateToProps = state => ({
    history: state.commitment.history
})

const mapDispatchToProps = dispatch => {
    const begin = 0, num = 5;
    const props = {
        onConstruct: () => { dispatch(Actions.requestFetchHistory(begin, num)) }
    }
    return props
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(HistoryOverview)