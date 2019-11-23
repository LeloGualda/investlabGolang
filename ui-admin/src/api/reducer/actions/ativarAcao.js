import ativarAction from '../../middle/ativarAcao'
import actions from './actions'

export default function (acao) {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            ativarAction(api,acao)
                // .then(() => 
                //     // dispatch({ type: actions.comprar})
                //  )
                .catch(error => {
                    //TODO
                })
        } catch (e) {
            console.error(e)
        }
    }
}