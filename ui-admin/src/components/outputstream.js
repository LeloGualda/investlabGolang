import React, { Component } from 'react'
import './sass/outputstream.scss'

export default class Output extends Component {
  render() {
    const {lines} = this.props || []

    return (
      <div className="outputstream">
          {lines.map( line => <div className="line"> {line} </div>)}
      </div>
    )
  }
}
