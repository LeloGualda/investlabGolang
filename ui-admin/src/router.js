import React, { Component } from "react";
import { BrowserRouter, Route, Switch, Redirect } from "react-router-dom";
import {connect} from 'react-redux'
import Login from './components/login'
import logout from './api/reducer/actions/logout'
import Usuarios from './pages/usuarios'
import Acoes from './pages/acoes'


class Router extends Component {
  render() {
    const {active,load,user} = this.props
    if(!load) return <div> carregando....</div>
    return <BrowserRouter>{
          active ? <Switch>
          <Route exact path="/admin" component={() => <h1>AREA DO ADMIN {user.username} </h1>} />
          <Route exact path="/admin/usuarios" component={() => <Usuarios/>} />
          <Route exact path="/admin/acoes" component={() => <Acoes/>} />
          <Route exact path="/logout" component={() => { this.props.logout(); return <Redirect to="/" /> }} />
          <Route path="*" component={() => <h1>Page not found</h1>} />
          </Switch> : <Login/>
        }
      </BrowserRouter>
  }
}
export default connect(state => ({active:state.active,user:state.user,load:state.loadUserInfo}), {logout})(Router)