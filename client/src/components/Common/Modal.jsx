import * as React from 'react'
import styled from 'styled-components'
import { OverlayOut } from './OverLayout'

export default (props) => {
  return (
    <>
      <OverlayOut>
        <CloseButton onClick={() => props.changeState()} >
          ×
        </CloseButton>
      </OverlayOut>
    </>
  )
}

const CloseButton = styled.div`
  position: absolute;
  font-size: 200%;
  left: 76%;
  top: 15%;
  font-weight: bold;
  border: 2px solid #999;
  color: white;
  display: flex;
  justify-content: center;
  width: 1.3em;
  line-height: 1.3em;
  cursor: pointer;

  &:hover {
    background: #333;
    border-color: #333;
    color: #FFF;
  }
`