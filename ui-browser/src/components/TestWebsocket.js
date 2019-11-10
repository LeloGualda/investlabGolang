import React, { Component } from 'react'
import { w3cwebsocket as W3CWebSocket } from "websocket";
import Output from './outputstream'

class TestWebsocket extends Component {
 
  constructor(props) {
    super(props)
    this.send = this.send.bind(this)
    this.client = new W3CWebSocket(`ws://${document.location.host}/logger`);
    this.reciveMsg = this.reciveMsg.bind(this)
    this.state = {
      msg:[]
    }
  }
  reciveMsg(message){
    this.setState(state => ({msg:[...state.msg,message.data]}) )
  }
  componentDidMount() {
    this.client.onopen = () => {
      console.log('WebSocket Client Connected');
    };
    this.client.onmessage = this.reciveMsg
  }
  send(text) {
    this.client.send(text)
  }
  render() {
    const {msg} = this.state

    return <div>
      <div>
        <textarea placeholder="write your event" ref={el => this.textarea = el} style={{ height: '100px', width: '200px' }} />
        <div>
          <button onClick={() => this.send(this.textarea.value)}> Send Event </button>
        </div>
      </div>
      <Output lines={msg}/>
    </div>
  }
}


export default TestWebsocket