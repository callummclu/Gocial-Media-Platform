import { LogInUser } from "../Types/auth";

export async function logIn(loginParams: LogInUser){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/auth/login`,{
        method:"POST",
        body:JSON.stringify(loginParams)
    }).then(async (res:any) => {
        let res_json = await res.json()
        if (Object.hasOwn(res_json,'error')){
            return false
        } else {
            localStorage.setItem("gocial_auth_token",res_json.token)
            return true
        }
    }).catch(()=>{
        return false
    })
}

export function checkAuth(){
    let token = localStorage.getItem("gocial_auth_token")
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/auth/${token}`,)
        .then(async (res:any) => {
            let res_json = await res.json()
            return res_json.isAuthenticated
        })
}

export async function logOut(){
    localStorage.removeItem("gocial_auth_token")
    return false
}