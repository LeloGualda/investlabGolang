import React, { Component } from 'react'
import post from '../middle/post'

export default class Upload extends Component {
    constructor(props) {
        super(props)
        this.state = {
            selectedFile: null,
            loaded: 0
        }
        this.onChangeHandler = this.onChangeHandler.bind(this)
        this.onClickHandler = this.onClickHandler.bind(this)
    }
    onChangeHandler = event => {
        this.setState({
            selectedFile: event.target.files[0],
            loaded: 0,
        })
    }
    onClickHandler = () => {
        const data = new FormData()
        data.append('file', this.state.selectedFile)
        console.log(this.state)
        post('upload',data)
    }
    render() {
        const { onChangeHandler,onClickHandler } = this
        return (
            <div>
                <input type="file" name="uploadfile" onChange={onChangeHandler} />
                <input type="submit" value="upload" onClick={onClickHandler} />
            </div>
        )
    }
}
