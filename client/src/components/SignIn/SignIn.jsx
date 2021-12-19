import React from "react";
import { Actions } from "../../actions/SignIn";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

class SignIn extends React.Component {
  constructor(props) {
    super(props);
    this.state = { id: "", pw: "" };
  }

  render() {
    const changeState = (key, e) => this.setState({ [key]: e.target.value });
    return (
      <>
        {this.props.success_signin && <Redirect to="/" />}
        <div className="signin-form">
          <h1>SignIn Form</h1>
          <input type="id" name="id" placeholder="ID" onChange={(e) => { changeState("id", e); }} />
          <input type="pw" name="pw" placeholder="PW" onChange={(e) => { changeState("pw", e); }} />
          <button type="button" onClick={() => { this.props.signIn(this.state.id, this.state.pw); }} >
            SignIn
          </button>
        </div>
      </>
    );
  }
}

const mapStateToProps = (state) => ({
  success_signin: state.signIn.success_signin,
  failed_signin_reason: state.signIn.failed_signin_reason,
});

const mapDispatchToProps = (dispatch) => ({
  signIn: (id, pw) => {
    console.log("dispatch sign in");
    dispatch(Actions.requestSignIn(id, pw));
  },
});

export default connect(mapStateToProps, mapDispatchToProps)(SignIn);