 const buscar = (api,key) =>  api.get(`/api/admin/search/query?keyword=${key}`)
 export default buscar
