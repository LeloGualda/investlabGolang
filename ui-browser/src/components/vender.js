import React, {Component} from 'react'
import {withRouter} from 'react-router-dom';

import {connect} from 'react-redux'
import Container from 'react-bootstrap/Container'
import acoesUSer from '../api/reducer/actions/getAcoesUser'
import temporal from '../api/reducer/actions/temporal'
import lance from '../api/reducer/actions/lance'

import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Card from 'react-bootstrap/Card'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Plot from "react-plotly.js"
import Confirm from './confirm'

class App extends Component {
    constructor(props) {
        super(props)
        this.state = {
            item: {},
            lanceCompra: 1,
            lanceQuantidade: 1,
            lanceVender: true,
            show: false
        }
        this.setLance = this
            .setLance
            .bind(this)
        this.colocarAVenda = this
            .colocarAVenda
            .bind(this)
        this.close = this
            .close
            .bind(this)
        this.confirm = this
            .confirm
            .bind(this)
    }
    componentWillMount() {
        this
            .props
            .acoesUSer()
    }
    routeChange(item) {
        // let path = `/comprar/${item}`; this.props.history.push(path);
    }
    close() {
        this.setState({show: false})
    }
    colocarAVenda(e) {
        e.preventDefault()
        this.setState({show: true})
    }
    setLance(item) {
        this.setState({item})
        this
            .props
            .temporal(item.codigo)
    }
    confirm() {

        const objLance = {
            lance: this.state.lanceCompra,
            qtd: this.state.lanceQuantidade,
            venda: !this.state.item.venda,
            codigo: this.state.item.codigo
        }

        this
            .props
            .lance(objLance)
    }
    render() {
        const {
            acoesUsuario = [],
            seriesTemporal
        } = this.props

        const {
            codigo = null,
            qtd = 1
        } = this.state.item

        const serieAtual = seriesTemporal[codigo] || []
        const valorAtual = serieAtual.length > 1
            ? serieAtual[0].valor
            : 0
        // console.log(this.state)
        return <Container>
            <Confirm
                title={`VENDER: ${codigo}`}
                body={!this.state.item.venda
                ? ` DESEJA VENDER ${this
                    .state
                    .lanceQuantidade} AÇÕES <br/> CODIGO ${codigo} <br/> NO VALOR ATUAL DE ${Math
                    .round((valorAtual + valorAtual * (this.state.lanceCompra / 100)) * 100) / 100} <br/> COM VARIAÇÃO ACIMA DA AÇÃO DE ${this.state.lanceCompra}% `
                : ` RETIRAR AÇÃO DA VENDA ? `}
                show={this.state.show}
                confirm={this.confirm}
                close={this.close}/>
            <Row>
                <Col>
                    <Form onSubmit={this.colocarAVenda}>
                        <Form.Group>
                            <Form.Label>
                                AÇÃO:
                                <b>
                                    {codigo}
                                </b>
                            </Form.Label>
                            <br/>
                            <Form.Group
                                style={{
                                display: 'flex',
                                justifyContent: 'space-around'
                            }}>
                                <Form.Label>
                                    VALOR:
                                    <b>
                                        {serieAtual[0] && valorAtual || 0}
                                    </b>
                                </Form.Label>

                                <Form.Label>
                                    VENDA:
                                    <b>
                                        {Math.round((valorAtual + valorAtual * (this.state.lanceCompra / 100)) * 100) / 100}
                                    </b>
                                </Form.Label>

                                <Form.Label>
                                    LUCRO:
                                    <b>
                                        {(Math.round((-valorAtual + (valorAtual + valorAtual * (this.state.lanceCompra / 100))) * 100) / 100)}
                                    </b>
                                </Form.Label>
                            </Form.Group>
                            <Form.Label>
                                QUANTIDADE (MAX {qtd}) : {" "}
                                <Form.Control
                                    onChange={({target}) => {
                                    this.setState({lanceQuantidade: target.value})
                                }}
                                disabled={this.state.item.venda || false}

                                    value={this.state.lanceQuantidade}
                                    type="number"
                                    label="quantidade"
                                    style={{
                                    width: '120px',
                                    display: 'inline'
                                }}
                                    max={qtd}
                                    min={1}
                                    placeholder="quantidade"/>
                            </Form.Label>
                            <br/>

                            <Form.Label>
                                Selecione sua oferta em % {this.state.lanceCompra}
                            </Form.Label>

                            <Form.Control
                                onChange={({target}) => {
                                this.setState({lanceCompra: target.value})
                            }}
                                disabled={this.state.item.venda || false}
                                type="range"
                                value={this.state.lanceCompra}
                                max={100}
                                min={1}
                                placeholder="quantidade"/>

                            <Form.Label>
                            {!this.state.item.venda ?'AINDA NÃO ESTÁ A VENDA' : 'JÁ ESTÁ A VENDA'}
                            </Form.Label>
                            <br/>
                            <Button variant="primary" disabled={!codigo} type="submit">
                               {!this.state.item.venda ?' COLOCAR A VENDA' : 'RETIRAR AÇÃO'}
                            </Button>
                        </Form.Group>
                    </Form>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Plot
                        data={[{
                            x: serieAtual.map(({data}) => data),
                            y: serieAtual.map(({valor}) => valor),
                            type: "scatter",
                            mode: "lines",
                            marker: {
                                color: "blue"
                            }
                        }
                    ]}
                        layout={{
                        width: "200px",
                        height: "100%",
                        title: codigo
                    }}/>
                </Col>
            </Row>
            <Row  className="board-cards">
                <Col>
                    {acoesUsuario.map(item => <Row key={item.codigo}>
                        <Card
                            border="secondary Cardcompra"
                            onDoubleClick={() => this.setLance(item)}
                            style={{
                            width: '18rem'
                        }}>
                            <Card.Header>{item.codigo}</Card.Header>
                            <Card.Body>
                                <Card.Title>
                                    Quantidade: {item.qtd}
                                </Card.Title>
                                <Card.Text>
                                    A venda: {item.venda
                                        ? item.lance + "%"
                                        : 'não está a venda'}
                                    <br/>
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

export default connect(state => ({active: state.active, user: state.user, acoesUsuario: state.acoesUsuario, seriesTemporal: state.temporal}), {acoesUSer, temporal, lance})(withRouter(App))
