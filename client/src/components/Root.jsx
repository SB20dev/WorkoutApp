import React from 'react'
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom'
import Top from './Top/Top'
import Home from './Home/Home'
import Commitment from './Commitment/Commitment'
import Menu from './Menu/Menu'
import Commit from './Commit/Commit'
import NotFound from './NotFound/NotFound'
import SignIn from './SignIn/SignIn'
import SignUp from './SignUp/SignUp'

export default props => (
    <Router>
        <Switch>
            <Route path="/sign*" component={SignRoute} />
            <Route component={ContentRoute} />
        </Switch>
    </Router>
)

const SignRoute = props => (
    <Switch>
        <Route path="/signup" component={SignUp} />
        <Route path="/signin" component={SignIn} />
    </Switch>
)

const ContentRoute = props => (
    <div>
        <nav>
            <div className="nav-left">
                <div className="logo">
                    <Link to="/">ロゴ</Link>
                </div>
                <div className="nav-list-container">
                    <ul className="nav-list">
                        <li>
                            <Link to="/home">Home</Link>
                        </li>
                        <li>
                            <Link to="/commitment">Commitment</Link>
                        </li>
                        <li>
                            <Link to="/menu">Menu</Link>
                        </li>
                        <li>State</li>
                        <li>Profile</li>
                    </ul>
                </div>
            </div>
            <div className="nav-right">
                <div className="fixed-button-container">
                    <div className="fixed-button">
                        <Link to="logoff">logoff</Link>
                    </div>
                    <div className="fixed-button">
                        <Link to="/commit">Commit</Link>
                    </div>
                </div>
            </div>
        </nav>

        <div className="main-container">
            <div className="main">
                <Switch>
                    <Route exact path="/" component={Top} />
                    <Route path="/home" component={Home} />
                    <Route path="/commitment" component={Commitment} />
                    <Route path="/menu" component={Menu} />
                    <Route exact path="/commit" component={Commit} />
                    <Route component={NotFound} />
                </Switch>
            </div>
        </div>
    </div>
)