
import React, { Component } from 'react'
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux'
import comprar from '../api/reducer/actions/comprar'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Card from 'react-bootstrap/Card'
class App extends Component {

    componentWillMount() {
        this.props.comprar()
    }
    routeChange(item) {
        let path = `/comprar/${item}`;
        this.props.history.push(path);
      }

    render() {
        const { seriesCompra = [] } = this.props
        return <Container>
            <Row>
                <Col>
                    {seriesCompra.map(item => <Row key={item.codigo}>
                        <Card border="secondary Cardcompra" onDoubleClick={() => this.routeChange(item.codigo)} style={{ width: '18rem' }} >
                            <Card.Header>{item.codigo}</Card.Header>
                            <Card.Body>
                            <Card.Title> valor: {item.valor} </Card.Title>
                                <Card.Text>
                                    {item.data}
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

export default connect(state => ({ active: state.active, user: state.user, seriesCompra: state.seriesCompra }), { comprar })(withRouter(App))
