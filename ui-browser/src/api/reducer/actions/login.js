import signIn, { userInfo } from '../../middle/signIn'
import actions from './actions'

export default function sigIn(username, password) {
    return (dispatch, getState) => {
        const { api, login } = getState();
        signIn(api, username, password)
            .then(e => {
                dispatch({type: actions.logoin})
                login(btoa(username + ":" + password))
                userInfo(api).then(d => d.data)
                    .then(json => {
                        dispatch({ type: actions.resolveAuth})
                        dispatch({ type: actions.userInfo, user: json })
                    })
                    .catch(error => {
                        dispatch({ type: actions.errorAuth,text:"Falha na autenticação"})
                        dispatch({ type: actions.logout })
                    })
            })
            .catch(error => {
                dispatch({ type: actions.errorAuth,text:"Falha na autenticação"})
                dispatch({ type: actions.logout})
            })
    }
}