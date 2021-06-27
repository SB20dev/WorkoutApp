import React from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { Actions } from "../../actions/SignIn";

class SignOut extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <>
        {!this.props.success_signin && <Redirect to="/" />}
        <div className="signout-form">
          <h1>SignOut Form</h1>
          <button type="button" onClick={() => { this.props.signOut(); }} >
            SignOut
          </button>
        </div>
      </>
    );
  }
}

const mapStateToProps = (state) => ({
  success_signin: state.signIn.success_signin,
});

const mapDispatchToProps = (dispatch) => ({
  signOut: () => {
    dispatch(Actions.requestSignOut());
  },
});

export default connect(mapStateToProps, mapDispatchToProps)(SignOut);