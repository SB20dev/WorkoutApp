import React from "react";
import { Actions } from "../../actions/SignUp";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

class SignUp extends React.Component {
  constructor(props) {
    super(props)
    this.state = { id: "", pw: "" }
  }

  render() {
    const changeState = (key, e) => this.setState({ [key]: e.target.value });

    return (
      <>
        {this.props.success_signin && <Redirect to="/" />}
        <div className="signup-form">
          <h1>SignUp Form</h1>
          <input type="id" name="id" placeholder="ID" onChange={(e) => { changeState("id", e); }} />
          <input type="pw" name="pw" placeholder="PW" onChange={(e) => { changeState("pw", e); }} />
          <button type="button" onClick={() => { this.props.signUp(this.state.id, this.state.pw); }} >
            SignUp
          </button>
        </div>
      </>
    )
  }
}

const mapStateToProps = (state) => ({
  success_signup: state.signUp.success_signup,
  failed_signup_reason: state.signUp.failed_signup_reason,

  success_signin: state.signIn.success_signin,
  failed_signin_reason: state.signIn.failed_signin_reason,
})

const mapDispatchToProps = (dispatch) => ({
  signUp: (id, pw) => {
    dispatch(Actions.requestSignUp(id, pw));
  }
})

export default connect(mapStateToProps, mapDispatchToProps)(SignUp);

