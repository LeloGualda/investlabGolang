 const api = (api,{codigo,name,tipo}) =>  api.get(`/api/admin/addCarteira/query?codigo=${codigo}&tipo=${tipo}&nome=${name}`)
 export default api
