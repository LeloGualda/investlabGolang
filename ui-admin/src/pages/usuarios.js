import React, {Component} from 'react'
import {connect} from 'react-redux'

import Table from 'react-bootstrap/Table'
import InputGroup from 'react-bootstrap/InputGroup'
import FormControl from 'react-bootstrap/FormControl'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Button from 'react-bootstrap/Button'

import usersAction from '../api/reducer/actions/users'
import createUserAction from '../api/reducer/actions/createUser'

class Users extends Component {
    
    constructor(props){
        super(props)
        this.state = {
            username:'',
            password:''
        }
    }
    componentWillMount() {
        this
            .props
            .usersAction()
    }
    render() {

        const {users} = this.props
        const {username,password} = this.state

        return (
            <Container>
             <Row>
                    <InputGroup className="mb-3">
                        <InputGroup.Prepend>
                            <InputGroup.Text>NOME DO USUARIO</InputGroup.Text>
                        </InputGroup.Prepend>
                        <FormControl
                            max={10}
                            value={username}
                            onChange={({target}) => this.setState({username: target.value})}/>
                        <InputGroup.Prepend>
                            <InputGroup.Text>Senha</InputGroup.Text>
                        </InputGroup.Prepend>
                        <FormControl
                            value={password}
                            onChange={({target}) => this.setState({password: target.value})}/>
                        <Button onClick={ () => {
                            this.props.createUserAction({username,password})
                            setTimeout( () =>    this
                            .props
                            .usersAction() ,3000)
                        } } >
                            ADD
                        </Button>
                    </InputGroup>
                </Row>
   
                <Row>
                    <Table striped bordered hover>
                    <thead>
                            <tr>
                                <th>
                                    ID
                                </th>
                                <th>
                                    USER
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {users.map(({id, username}) => <tr key={username}>
                                <td>{id}</td>
                                <td>{username}</td>
                            </tr>)}
                        </tbody>
                    </Table>
                </Row>
            </Container>
        )
    }
}

export default connect(({users}) => ({users}), {usersAction,createUserAction})(Users)