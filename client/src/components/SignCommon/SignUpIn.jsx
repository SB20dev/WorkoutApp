import React from "react"
import { Actions as SignUpActions } from "../../actions/SignUp"
import { Actions as SignInActions } from "../../actions/SignIn"

import { connect } from "react-redux"
import ModalWithoutClose from "../Common/ModalWithoutClose"
import styled from 'styled-components'
import { Redirect } from "react-router"

class SignUpIn extends React.Component {
  constructor(props) {
    super(props)
    this.state = { sign_in_id: '', sign_in_pw: '', sign_up_id: '', sign_up_pw: '' }
  }

  render() {
    const changeState = (key, e) => this.setState({ [key]: e.target.value });
    return (

      <>
        <ModalWithoutClose />
        <SignInDiv>
          <h3>Sign in</h3>
          <input type='text' placeholder='ID' onChange={(e) => { changeState('sign_in_id', e) }}></input>
          <input type='text' placeholder='PW' onChange={(e) => { changeState('sign_in_pw', e) }}></input>
          <button onClick={() => { this.props.signIn(this.state.sign_in_id, this.state.sign_in_pw) }}> sign in </button>
        </SignInDiv>
        <SignUpDiv>
          <h3>Sign up</h3>
          <input type='text' placeholder='ID' onChange={(e) => { changeState('sign_up_id', e) }}></input>
          <input type='text' placeholder='PW' onChange={(e) => { changeState('sign_up_pw', e) }}></input>
          <button onClick={() => { this.props.signUp(this.state.sign_up_id, this.state.sign_up_pw) }}> sign up </button>
        </SignUpDiv>
      </>
    )
  }
}

const SignInDiv = styled.div`
        position: absolute;
        font-size: 200%;
        left: 25%;
        top: 10%;
        border: 2px solid #999;
        width: 50%;
        height:20%;
        background-color: white;
        `
const SignUpDiv = styled.div`
        position: absolute;
        font-size: 200%;
        left: 25%;
        top: 40%;
        border: 2px solid #999;
        width: 50%;
        height: 20%;
        background-color: white;
        `
const mapStateToProps = (state) => ({
  success_signup: state.signUp.success_signup,
  failed_signup_reason: state.signUp.failed_signup_reason,

  success_signin: state.signIn.success_signin,
  failed_signin_reason: state.signIn.failed_signin_reason,
})

const mapDispatchToProps = (dispatch) => ({
  signUp: (id, pw) => {
    dispatch(SignUpActions.requestSignUp(id, pw));
  },
  signIn: (id, pw) => {
    dispatch(SignInActions.requestSignIn(id, pw));
  }
})

export default connect(mapStateToProps, mapDispatchToProps)(SignUpIn);

