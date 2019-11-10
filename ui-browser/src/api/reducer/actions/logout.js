import actions from './actions'

export default function logout() {
    return (dispatch) => {
        console.log('thau volte logo :)')
        dispatch({ type: actions.logout})
    }

}