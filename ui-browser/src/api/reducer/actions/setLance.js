import actions from './actions'

export default function setLance(lanceAtual) {
    return (dispatch) => {
        dispatch({ type: actions.setLance,lanceAtual})
    }

}