import React from "react"
import { connect } from "react-redux"
import { Link } from "react-router-dom"
import Modal from '../Common/Modal'
import RegistParts from './RegistParts'
import Table from '../Common/Table'
import styled from "styled-components"

class Parts extends React.Component {
  constructor(props) {
    super(props)
    this.state = { isShow: false }
  }


  render() {
    console.log(this.state)

    const changeState = (key, value) => this.setState({ [key]: value })
    console.log('hogehoge')
      
    const data = {
      data : [
        {
          CLASS: "arm",
          DETAIL: "upper arm r",
          STATE: "bad",        
        },
        {
          CLASS: "arm",
          DETAIL: "lower arm r",
          STATE: "good",
        },
        {
        CLASS: "arm",
        DETAIL: "upper arm l",
        STATE: "OK",
        }
      ],
      openAddFunc : 
        {
        CLASS: console.log,
        DETAIL: console.log
        },
      openFilterFunc : 
        {
        CLASS: console.log,
        DETAIL: console.log
        }
    };
    
    return (
      <>
        <button type="button" onClick={() => changeState("isShow", true)}>
          regist parts
        </button>
        {this.state.isShow && (
          <>
            <Modal changeState={() => changeState("isShow", false)} />
            <RegistParts />
          </>
        )}
        <Table data={data} />
      </>
    );
  }
}

const mapStateToProps = (state) => ({

})

export default connect(mapStateToProps)(Parts)