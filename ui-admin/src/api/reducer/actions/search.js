import searchAction from '../../middle/search'
import actions from './actions'

export default function (keyword) {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            searchAction(api,keyword)
                .then(d => d.data)
                .then(json => dispatch({ type: actions.buscaApi,data:json}))
                .catch(error => {
                    //TODO
                })
        } catch (e) {
            console.error(e)
        }
    }
}