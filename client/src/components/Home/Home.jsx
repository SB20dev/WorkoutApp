import React from 'react'
import CommitmentOverview from './CommitmentOverview'
import Suggestion from './Suggestion'
import HistoryOverview from './HistoryOverview'
import './home.css'

export default () => (
    <div>
        <div className="commit-and-suggest">
            <CommitmentOverview />
            <Suggestion />
        </div>
        <div className="history">
            <HistoryOverview />
        </div>
    </div>
)