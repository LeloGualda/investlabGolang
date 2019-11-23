import addAcaoAction from '../../middle/addAcao'
import actions from './actions'

export default function (acao) {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            addAcaoAction(api,acao)
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