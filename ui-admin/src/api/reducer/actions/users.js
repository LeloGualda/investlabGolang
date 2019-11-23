import users from '../../middle/users'
import actions from '../../../../../ui-browser/src/api/reducer/actions/actions'

export default function () {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            users(api)
                .then(d => d.data)
                .then(json => {
                    dispatch({ type: actions.getUser, data: json })
                })
                .catch(error => {
                    // dispatch({ type: actions.comprarError, text: "não encontrado nenhuma ação disponivel para compra" })
                })
        } catch (e) {
            console.error(e)
        }
    }
}