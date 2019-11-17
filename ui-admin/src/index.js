import React, { Component } from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import store from './api/reducer/store'
import Router from './router'
import Navbar from './components/Navbar'

import 'bootstrap/dist/css/bootstrap.min.css'
import './components/sass/outputstream.scss'

class App extends Component {
	render() {
		return <Provider store={store}>
			<Navbar />
			<Router style={{ padding: '10px 20px' }} />
		</Provider>
	}
}

ReactDOM.render(<App />, document.getElementById('react'))