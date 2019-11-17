import lanceAction from '../../middle/lance'
import actions from './actions'

export default function (lance) {
    return (dispatch, getState) => {
        const { api } = getState();
        try {
            lanceAction(api,lance)
                .then(json => {
                    setTimeout(() => {
                        document.location.reload()
                    }, 2000);
                })
                .catch(error => {
                    console.log('error')
                })
        } catch (e) {
            console.error(e)
        }
    }
}