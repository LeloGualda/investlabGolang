import React, { Component } from "react";
import { connect } from "react-redux";
import Card from "react-bootstrap/Card";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from 'react-bootstrap/Col'
import Plot from "react-plotly.js"
import TestWeb from './TestWebsocket'

class App extends Component {
  constructor(props) {
    super(props)
    this.client = this.props.newclient()
    this.reciveMsg = this.reciveMsg.bind(this)
    this.state = {
      msg:[]
    }
  }
  componentWillMount(){
    
  }
  reciveMsg(message){
    this.setState(state => ({msg:[...state.msg,message.data]}) )
  }
  render() {
    const {msg} = this.state

    console.log(msg)
    return (
      <Container>
        <TestWeb/>
        <Row className="justify-content-md-center">
          <Col className="chart">
            <Card bg="light" style={{ width: "750px", padding: "10px 2px" }}>
              <Card.Header>Header</Card.Header>
              <Card.Body>
                <Card.Title>Primary Card Title</Card.Title>
                <Card.Text>
                  Tendencia:
                  <Plot
                    data={[
                      {
                        x: [1, 2, 3],
                        y: [2, 6, 3],
                        type: "scatter",
                        mode: "lines+points",
                        marker: { color: "red" }
                      },
                      { type: "bar", x: [1, 2, 3], y: [2, 5, 3] }
                    ]}
                    layout={{ width: "500px", height: "100%", title: null }}
                  />
                </Card.Text>
              </Card.Body>
            </Card>
          </Col>
        </Row>
      </Container>
    );
  }
}

export default connect(
  state => ({ active: state.active, user: state.user,newclient:state.newclient }),
  {}
)(App);
