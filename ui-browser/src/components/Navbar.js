import React, { Component } from 'react'
import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'
import { connect } from 'react-redux'

class App extends Component {
    render() {
        const { active, user } = this.props

        return <div className="App">
            <Navbar bg="dark" variant="dark">{
                active ? <Navbar.Collapse>
                    <Nav className="mr-auto">
                        <Navbar.Brand  href="/" style={{ color: 'white', textTransform: 'uppercase' }}>{user.username}</Navbar.Brand>
                        <Nav.Link href="/comprar">Compra</Nav.Link>
                        <Nav.Link href="/vender">Vender</Nav.Link>
                        <Nav.Link href="/graficos">Gráficos</Nav.Link>
                    </Nav>
                    <Nav>
                        <Navbar.Brand className="mr-sm-2" href="/logout"> Logout</Navbar.Brand>
                    </Nav>
                </Navbar.Collapse> : <Navbar.Brand>Login</Navbar.Brand>
            }</Navbar>
        </div>
    }
}

export default connect(state => ({ active: state.active, user: state.user }), {})(App)


