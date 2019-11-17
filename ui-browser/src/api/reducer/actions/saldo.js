import saldo from '../../middle/saldo'
import actions from './actions'

export default function () {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            saldo(api)
                .then(d => d.data)
                .then(json => dispatch({ type: actions.getSaldo,saldo:json.valor || 0}))
                .catch(error => {
                    //TODO
                })
        } catch (e) {
            console.error(e)
        }
    }
}