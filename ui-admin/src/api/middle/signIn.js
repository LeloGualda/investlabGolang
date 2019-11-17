export default function (api,username,password){
    const response =  api.post("/session", { username, password })
    
    return response
}

export function userInfo(api){
    return api.get("/user/info")
}
