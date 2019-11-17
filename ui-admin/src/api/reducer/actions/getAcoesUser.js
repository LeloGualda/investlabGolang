import acoesUser from '../../middle/acoesUser'
import actions from './actions'

export default function comprarAction() {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            acoesUser(api)
                .then(d => d.data)
                .then(json => {
                    dispatch({ type: actions.getAcoesUser, data: json })
                })
                .catch(error => {
                    // dispatch({ type: actions.comprarError, text: "não encontrado nenhuma ação disponivel para compra" })
                })
        } catch (e) {
            console.error(e)
        }
    }
}