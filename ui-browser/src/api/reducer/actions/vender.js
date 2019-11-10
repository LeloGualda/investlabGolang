import vender from '../../middle/vender'
import actions from './actions'

export default function venderApi(acao) {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            vender(api,acao)
                .then(d => {
                    setTimeout( ()=> dispatch({ type: actions.vendeuAcoes}),2000)
                })
                .catch(error => {
                    
                })
        } catch (e) {
            console.error(e)
        }
    }
}