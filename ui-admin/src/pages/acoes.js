import React, {Component} from 'react'
import {connect} from 'react-redux'

import Table from 'react-bootstrap/Table'
import InputGroup from 'react-bootstrap/InputGroup'
import FormControl from 'react-bootstrap/FormControl'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Button from 'react-bootstrap/Button'

import comprar from '../api/reducer/actions/comprar'
import buscarAcoes from '../api/reducer/actions/search'

class Acoes extends Component {
    constructor(props) {
        super(props)
        this.state = {
            kerword: ''
        }
    }
    componentWillMount() {
        this
            .props
            .comprar()
    }
    render() {
        const {
            seriesCompra = [],
            acoes = {}
        } = this.props
        const {keyword} = this.state
        const {
            BestMatches = []
        } = acoes

        return (
            <Container>
                <Row>
                    <InputGroup className="mb-3">
                        <InputGroup.Prepend>
                            <InputGroup.Text>Buscar nova ação</InputGroup.Text>
                        </InputGroup.Prepend>
                        <FormControl
                            value={keyword}
                            onChange={({target}) => this.setState({keyword: target.value})}/>
                        <Button onClick={() => this.props.buscarAcoes(keyword)}>
                            Buscar
                        </Button>
                    </InputGroup>
                </Row>
                <Row>
                    RESULTADO DA CONSULTA:
                </Row>
                <Row>
                    <Table striped bordered hover>
                        <thead>
                            <tr>
                                <th>
                                    CODIGO
                                </th>
                                <th>
                                    NOME
                                </th>
                                <th>
                                    ADICIONAR
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {BestMatches.map(acao => <tr key={acao['1. symbol']}>
                                <td>{acao['1. symbol']}</td>
                                <td>{acao['2. name']}</td>
                                <td><Button> add</Button></td>
                            </tr>)}
                        </tbody>
                    </Table>
                </Row>
                <Row>
                    SUAS AÇÕES:
                </Row>
                <Row>
                    <Table striped bordered hover>
                    <thead>
                            <tr>
                                <th>
                                    CODIGO
                                </th>
                                <th>
                                    ULTIMA ATUALIZAÇÃO
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {seriesCompra.map(({codigo, data}) => <tr key={codigo}>
                                <td>{codigo}</td>
                                <td>{data}</td>
                            </tr>)}
                        </tbody>
                    </Table>
                </Row>
            </Container>
        )
    }
}

export default connect(({seriesCompra, buscaAcoes}) => ({seriesCompra, acoes: buscaAcoes}), {comprar, buscarAcoes})(Acoes)