import temporal from '../../middle/temporal'
import actions from './actions'

export default function (codigo) {
    return (dispatch, getState) => {
        const { api } = getState();
        try {
            temporal(api,codigo)
                .then(d => d.data)
                .then(json => dispatch({ type: actions.getTemporalAcao,data:json,codigo}))
                .catch(error => {
                    //TODO
                })
        } catch (e) {
            console.error(e)
        }
    }
}