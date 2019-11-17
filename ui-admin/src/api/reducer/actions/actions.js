import axios from "axios";
import { w3cwebsocket as W3CWebSocket } from "websocket";

const TOKEN_KEY = "@investlab-Token";
const api = axios.create({
    baseURL: "/"
});
const getToken = () => localStorage.getItem(TOKEN_KEY)
const newclient = () => new W3CWebSocket(`ws://${document.location.host}/logger`);

api.interceptors.request.use(async config => {
    const token = getToken();
    if (token) {
        config.headers.Authorization = `Basic ${token}`;
    }
    return config;
});

export default {
    logout: "LOGOUT",
    logoin: "LOGIN",
    userInfo: 'GET_USER_INFO',
    loadUserInfo: 'LOAD_USER_INFO',
    errorAuth: 'ERROR_AUTH',
    resolveAuth: 'RESOLVE_AUTH',
    comprar: 'COMPRAR',
    comprarError: 'COMPRAR_ERROR',
    comprarOfertas: 'COMPRAR_OFERTAS',
    setLance: "SET_LANCE",
    getAcoesUser: "GET_ACOES_USER",
    vendeuAcoes: "VENDEU_ACOES",
    getSaldo: "GET_SALDO_USER",
    getTemporalAcao:"GET_TEMPORAL_ACAO",
    initialState: {
        api,
        newclient,
        vendeuAcao: false,
        user: {
            username: '',
            id: 0
        },
        erros: {
            auth: null
        },
        temporal:{},
        saldo: 0,
        active: false,
        getToken,
        loadUserInfo: true,
        isAuthenticated: () => localStorage.getItem(TOKEN_KEY) !== null,
        logout: () => { localStorage.removeItem(TOKEN_KEY) },
        login: token => { localStorage.setItem(TOKEN_KEY, token) }
    }
}