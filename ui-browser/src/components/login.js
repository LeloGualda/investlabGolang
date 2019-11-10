import React, { Component } from "react";
import { Button, Form, Alert } from "react-bootstrap";
import { connect } from "react-redux";
import loginAction from "../api/reducer/actions/login";

class Login extends Component {
  constructor(props) {
    super(props);

    this.state = {
      username: "aurelio",
      password: "123456"
    };
    this.onChangeNode = this.onChangeNode.bind(this);
  }

  login = async e => {
    e.preventDefault();
    const { username, password } = this.state;
    this.props.loginAction(username, password);
  };
  onChangeNode({ target }) {
    this.setState({ [target.name]: target.value });
  }
  render() {
    const { username, password } = this.state;
    const { erros = {} } = this.props;
    const { auth: erro } = erros;
    // const {auth:erro} = this.props.erros
    console.log('ERRO:' + erro);
    return (
      <div>
        <Form onSubmit={this.login}>
          <Form.Group controlId="loginEmail">
            <Form.Label>username</Form.Label>
            <Form.Control
              type="text"
              name="username"
              onChange={this.onChangeNode}
              placeholder="Enter username"
              value={username}
            />
            <Form.Text className="text-muted">insira seu username</Form.Text>
          </Form.Group>

          <Form.Group controlId="loginPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control
              type="password"
              name="password"
              onChange={this.onChangeNode}
              placeholder="Password"
              value={password}
            />
          </Form.Group>
          <Button variant="primary" type="submit">
            Submit
          </Button>
          {erro ? <Alert variant={"danger"}>{erro}</Alert> : null}
        </Form>
      </div>
    );
  }
}

export default connect(
  state => ({ active: state.active, erros: state.erros }),
  { loginAction }
)(Login);
