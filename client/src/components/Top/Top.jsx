import React from "react";
import { connect } from "react-redux";
import { Link } from "react-router-dom";

class Top extends React.Component {
  constructor(props) {
    super(props)
  }

  render() {
    console.log(this.props.success_signin)
    return (
      <>
        {this.props.success_signin ?
          <h1>You are login.<br /> Let's begin training.</h1>
          :
          <>
            <Link to="/signup">
              SignUp
            </Link>
            <br />
            <Link to="/signin">
              SignIn
            </Link>
          </>
        }
      </>
    )
  }
}

const mapStateToProps = (state) => ({
  success_signin: state.signIn.success_signin,
})

export default connect(mapStateToProps)(Top)