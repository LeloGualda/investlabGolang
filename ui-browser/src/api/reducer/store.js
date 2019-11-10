import thunk from 'redux-thunk';
import { createStore, applyMiddleware } from 'redux';

import reducer from './index'
import actions from './actions/actions'
import { userInfo } from '../middle/signIn'

const store = createStore(reducer, applyMiddleware(thunk));
store.getState().getToken() ?
    userInfo(store.getState().api)
        .then(e => e.data)
        .then(json => { store.dispatch({ type: actions.userInfo, user: json }) })
        .catch(() => store.dispatch({ type: actions.logout }))
        .finally(() => {
            store.dispatch({ type: actions.loadUserInfo })
        }) : store.dispatch({ type: actions.loadUserInfo })

export default store