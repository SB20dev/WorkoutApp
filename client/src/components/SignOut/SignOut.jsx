import React from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

class SignOut extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <>
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

const mapDispatchToProps = (dispatch) => ({
  signOut: () => {
    dispatch(Actions.requestSignOut());
  },
});

export default connect(mapDispatchToProps)(SignOut);