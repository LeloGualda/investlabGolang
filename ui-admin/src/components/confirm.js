import React, { Component } from 'react'
import { Modal, Button } from 'react-bootstrap'

export default class Confirm extends Component {
    render() {
        const {title = "",body = "",show = false,confirm,close} = this.props

        return <Modal show={show} onHide={close}>
            <Modal.Header closeButton>
                <Modal.Title>{title}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <p>
                <div dangerouslySetInnerHTML={{__html:body}} />
                </p>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={close}>Sair</Button>
                <Button variant="primary" onClick={confirm}>Confirmar</Button>
            </Modal.Footer>
        </Modal>
    }
}


