import atualizarAcaoAction from '../../middle/atualizarAcao'
import actions from './actions'

export default function (codigo) {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            atualizarAcaoAction(api,codigo)
                .then(() => dispatch({ type: actions.atualizarAcao}))
                .catch(error => {
                    //TODO
                })
        } catch (e) {
            console.error(e)
        }
    }
}