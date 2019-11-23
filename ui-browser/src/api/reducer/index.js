import actions from './actions/actions'

function reducer(state, action) {
  if (typeof state === 'undefined') {
    return { ...actions.initialState, loadUserInfo: false }
  }
  const reset = () => ({ ...actions.initialState, erros: state.erros })

  switch (action.type) {
    case actions.logout:
      state.logout()
      return reset()
    case actions.logoin:
      return { ...state, active: true }

    case actions.userInfo:
      if (action.user.id > 0) {
        return { ...state, user: action.user, active: true }
      }
      return reset()

    case actions.loadUserInfo:
      return { ...state, loadUserInfo: true }
    case actions.errorAuth:
      return {
        ...state, erros: {
          ...state.erros,
          auth: action.text
        }
      }
    case actions.resolveAuth:
      return {
        ...state, erros: {
          ...state.erros,
          auth: null
        }
      }
    case actions.comprar:
      return {
        ...state, seriesCompra: action.data
      }
    case actions.comprarOfertas:
      return {
        ...state, lances: action.data
      }
    case actions.setLance:
      return {
        ...state, lanceAtual: action.lanceAtual
      }
    case actions.getAcoesUser:
      return {
        ...state, acoesUsuario: action.data
      }
    case actions.getSaldo:
      return {
        ...state, saldo: action.saldo
      }
    case actions.getTemporalAcao:
      return {
        ...state, temporal: {
          ...state.temporal,
          [action.codigo]: action.data
        }
      }
    case actions.vendeuAcoes:
      return {
        ...state, vendeuAcao: true, lanceAtual: null, lances: null
      }

    case actions.getUser:
      return {
        ...state, users:action.data
      }
  
    default:
      return state
  }
}

export default reducer