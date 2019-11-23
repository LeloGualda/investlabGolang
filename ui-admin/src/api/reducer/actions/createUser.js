import createUsercaoAction from '../../middle/createUser'
import actions from './actions'

export default function (user) {
    return (dispatch, getState) => {
        const { api } = getState();

        try {
            createUsercaoAction(api,user)
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