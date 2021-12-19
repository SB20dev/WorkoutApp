import React from "react";
import { connect } from "react-redux";
import { Link } from "react-router-dom";
import styled from 'styled-components';

class RegistParts extends React.Component {
  constructor(props) {
    super(props)
    this.state = { isShow: false }
  }
  
  
  render() {
    console.log("regist parts " + this.state)

    return (
      <RegistWindow>
        <h3>hogehog</h3>
        <></>
      </RegistWindow>
    )
  }
}

const mapStateToProps = (state) => ({
})

const RegistWindow = styled.div`
position: absolute;
font-size: 200%;
left: 25%;
top: 20%;
border: 2px solid #999;
color: white;
width: 50%;
height: 50%;

&:hover {
  background: #333;
  border-color: #333;
  color: #FFF;
}
`

export default connect(mapStateToProps)(RegistParts)