import React from "react";
import { connect } from "react-redux";
import { Link } from "react-router-dom";
import Modal from '../Common/Modal'

class Parts extends React.Component {
  constructor(props) {
    super(props)
    this.state = {isShow: false}
  }

  
  render() {
    console.log(this.state)

    const changeState = (key, value) => this.setState({ [key]: value })
    console.log('hogehoge')

    return (
      <>
        <button type='button' onClick={() => changeState('isShow', true)}>regist parts</button> 
        {this.state.isShow && <Modal changeState={()=> changeState('isShow',false) } />}
      </>
    )
  }
}

const mapStateToProps = (state) => ({

//  success_signin: state.signIn.success_signin,
})

export default connect(mapStateToProps)(Parts)