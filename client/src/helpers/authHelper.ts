import { LogInUser, SignUpUser } from "../Types/auth";

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

export async function signup(signupParams: SignUpUser){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/user`,{
        method:"POST",
        body:JSON.stringify(signupParams)
    }).then(async (res:any) => {
        let res_json = await res.json()
        if (Object.hasOwn(res_json,'error')){
            return false
        } else {
            return logIn({username:signupParams.username,password:signupParams.password})
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
            return {isAuthenticated:res_json.isAuthenticated,username:res_json.username}
        })
}

export async function logOut(){
    localStorage.removeItem("gocial_auth_token")
    return false
}