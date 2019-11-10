import React, { Component } from "react";
import { BrowserRouter, Route, Switch, Redirect } from "react-router-dom";
import {connect} from 'react-redux'
import Login from './components/login'
import Comprar from './components/comprar'
import ComprarAcoes from './components/comprarAcao'
import Vender from './components/vender'
import Graficos from './components/graficos'
import logout from './api/reducer/actions/logout'

class Router extends Component {
  render() {
    const {active,load,user} = this.props
    if(!load) return <div> carregando....</div>
    return <BrowserRouter>{
          active ? <Switch>
          <Route exact path="/" component={() => <h1>Bem vindo ao InvestaLab 2.0 {user.username} </h1>} />
          <Route exact path="/comprar" component={() => <Comprar/>} />
          <Route  path="/comprar/:codigo" component={ComprarAcoes} />
          <Route exact path="/vender" component={() => <Vender/>} />
          <Route exact path="/graficos" component={() => <Graficos/>} />
          <Route exact path="/logout" component={() => { this.props.logout(); return <Redirect to="/" /> }} />
          <Route path="*" component={() => <h1>Page not found</h1>} />
          </Switch> : <Login/>
        }
      </BrowserRouter>
  }
}
export default connect(state => ({active:state.active,user:state.user,load:state.loadUserInfo}), {logout})(Router)