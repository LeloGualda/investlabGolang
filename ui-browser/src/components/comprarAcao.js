import React, { Component } from 'react'
import { connect } from 'react-redux'

import { withRouter } from 'react-router-dom'
import { Form, Button, Card, Row, Col, Container } from 'react-bootstrap'
import Confirm from './confirm'

import comprarAcao from '../api/reducer/actions/comprarAcao'
import setLance from '../api/reducer/actions/setLance'
import venderAcao from '../api/reducer/actions/vender'


class App extends Component {

    constructor(props) {
        super(props)
        this.codigo = ""
        this.state ={
            quantidadeCompras:1,
            show:false,
            title:'',
            body:''
        }
        this.venda = this.venda.bind(this)
        this.confirm = this.confirm.bind(this)
        this.close = this.close.bind(this)
    }
    venda(e){
        e.preventDefault()
        const {valor} = this.props.lanceAtual
        const {quantidadeCompras} = this.state
        this.setState({show:true,title:`AÇÃO ${this.codigo}`,body:`
        Deseja comprar ${quantidadeCompras} açõe(s)
            no valor de ${valor * quantidadeCompras}
        ?`} )
    }
    confirm(){
        const {valor,codigo} = this.props.lanceAtual
        const {quantidadeCompras} = this.state

        this.props.venderAcao({valor,codigo,qtd:quantidadeCompras})
        document.location.href  = "/vender"
    }
    componentWillMount() {
        this.codigo = this.props.match.params.codigo
        this.props.comprarAcao(this.codigo)
    }

    close(){
        this.setState({show:false})
    }
    fSetLance(lance) {
        this.props.setLance(lance)
     }
    render() {
        const { lances = [],lanceAtual = ({valor:null}),vendeuAcao } = this.props

        const {valor  = 0,qtd = 0} =  lanceAtual || {}

        if(vendeuAcao){
          return  this.routeChange()
        }

        return <Container>
            <Confirm {...this.state} confirm={this.confirm} close={this.close}/>
            <Row>
                <Col>
                    <Container>
                        <Form  onSubmit={this.venda}>
                            <Form.Group controlId="formBasicEmail">
                                <Form.Label>Ação: <b> {this.codigo} </b> </Form.Label> <br/>
                                <Form.Label>Valor: <b> {valor || "selecione um lance"} </b> </Form.Label>
                                <Form.Control
                                onChange={ ({target}) => {
                                    this.setState({quantidadeCompras:target.value})
                                } }
                                value={this.state.quantidadeCompras}
                                 type="number" max={qtd || 1} min={1} placeholder="quantidade" />
                            </Form.Group>
                            <Button variant="primary"  disabled={!valor} type="submit">
                                Comprar
                            </Button>
                        </Form>
                    </Container>
                </Col>
            </Row>
            <Row>
                <Col>
                    {lances.map((lance,ids) => <Card key={`of${ids}`} onDoubleClick={() => this.fSetLance(lance)} border="secondary Cardcompra" style={{ width: '18rem' }} >
                            <Card.Header>{lance.codigo}</Card.Header>
                            <Card.Body>
                                <Card.Title>Valor R$ {lance.valor}</Card.Title>
                                <Card.Text>
                                    Quantidade : {lance.qtd}
                                </Card.Text></Card.Body>
                        </Card>)}
                </Col>
            </Row>
        </Container>
    }
}

export default connect(state => ({ active: state.active, user: state.user, lances: state.lances,lanceAtual:state.lanceAtual,vendeuAcao:state.vendeuAcao }), { comprarAcao,setLance,venderAcao })(withRouter(App))


