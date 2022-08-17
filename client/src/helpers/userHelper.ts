import { User } from "../Types/auth"

export function getUser(searchParameters:string, page:number){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/user?searchParams=${searchParameters}&itemsPerPage=20&page=${page}`)
}

export function acceptFriendRequest(user:User, requestUser:string){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/user/friends/${user?.username}/${requestUser}/${localStorage.getItem("gocial_auth_token")}/accept`,{
        method:"POST"
    })
}

export function getSingleUser(username:string){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/user/${username}`)
}

export function sendInvitation(user:User, username:string){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/user/invitation/${user?.username}/${username}/${localStorage.getItem("gocial_auth_token")}`,{
        method:"POST"
    })
}