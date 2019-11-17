import comprarAcoes from '../../middle/comprarAcoes'
import actions from './actions'

export default function comprarAction(acao) {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            comprarAcoes(api,acao)
                .then(d => d.data)
                .then(json => {
                    dispatch({ type: actions.comprarOfertas, data: json })
                })
                .catch(error => {
                    // dispatch({ type: actions.comprarError, text: "não encontrado nenhuma ação disponivel para compra" })
                })
        } catch (e) {
            console.error(e)
        }
    }
}