 const lance = (api,{lance,qtd,venda,codigo}) =>  api.get(`/api/lance/query?lance=${lance}&codigo=${codigo}&qtd=${qtd}&vender=${venda}`)
 export default lance
