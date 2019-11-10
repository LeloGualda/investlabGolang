
import React, { Component } from 'react'
import { withRouter } from 'react-router-dom';

import { connect } from 'react-redux'
import Container from 'react-bootstrap/Container'
import acoesUSer from '../api/reducer/actions/getAcoesUser'

import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Card from 'react-bootstrap/Card'

class App extends Component {

    componentWillMount() {
        this.props.acoesUSer()
    }
    routeChange(item) {
        // let path = `/comprar/${item}`;
        // this.props.history.push(path);
    }

    render() {
        const { acoesUsuario = [] } = this.props
        return <Container>
            <Row>
                <Col>
                    {acoesUsuario.map(item => <Row key={item.codigo}>
                        <Card border="secondary Cardcompra" style={{ width: '18rem' }} >
                            <Card.Header>{item.codigo}</Card.Header>
                            <Card.Body>
                                <Card.Title> Quantidade: {item.qtd} </Card.Title>
                                <Card.Text>
                                    A venda: {item.venda ? 'está a venda' : 'não está a venda'}
                                </Card.Text>
                            </Card.Body>
                        </Card>
                    </Row>)
                    }
                </Col>
            </Row>

        </Container>
    }
}

export default connect(state => ({ active: state.active, user: state.user, acoesUsuario: state.acoesUsuario }), { acoesUSer })(withRouter(App))
