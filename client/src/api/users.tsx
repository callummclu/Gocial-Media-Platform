import { SignUpUser } from "../Types/auth"

export async function signup(signupParams: SignUpUser){
    const response = fetch(`${process.env.REACT_APP_BACKEND_URI}/user`,{
        method:"POST", body:JSON.stringify(signupParams)
    })
    return response
}

export function checkAuth(){
    let token = localStorage.getItem("gocial_auth_token")
    const response = fetch(`${process.env.REACT_APP_BACKEND_URI}/auth/${token}`)
    return response
}

export function getUserDetails(username:string){
    const response = fetch(`${process.env.REACT_APP_BACKEND_URI}/user/${username}`)
    return response
}  
