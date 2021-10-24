import * as React from "react";
import * as ReactDOM from "react-dom";
import styled from "styled-components";

export default (props) => {
  return (
    <>
      <Modal>
        <button onClick={() => props.changeState('isShow',false)} >
          closeModal
        </button>
      </Modal>
    </>
  )
}

const Modal = styled.div`
  position : absolute;
  left : 0;
  top : 0;
  width : 100%;
  height : 100%;
  background : rgba(100, 100, 100, .8);
  zIndex : 100;
`
