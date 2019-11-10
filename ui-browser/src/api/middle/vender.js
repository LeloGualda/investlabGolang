 const api = (api,venda) =>  api.get(`/api/vender/query?valor=${venda.valor}&codigo=${venda.codigo}&qtd=${venda.qtd}`)
 export default api
