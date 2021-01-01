import React from 'react'
import { Actions } from '../../actions/Menu'
import { connect } from 'react-redux'

class Suggestion extends React.Component {
    constructor(props) {
        super(props)
        props.onConstruct()
    }

    render() {
        const suggestion = this.props.suggestion.map((menu, idx) =>
            <li key={"suggestion-item"+idx}>
                {menu}
            </li>
        )
        return (
            <div className="suggestion">
                <h2>本日のおすすめメニュー</h2>
                {suggestion}
            </div>
        )
    }
}

const mapStateToProps = state => ({
    suggestion: state.menu.suggestion
})

const mapDispatchToProps = dispatch => ({
    onConstruct: () => { dispatch(Actions.requestFetchSuggestion()) }
})

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(Suggestion)