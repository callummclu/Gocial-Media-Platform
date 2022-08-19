import { Post } from "../Types/post";

export async function newPostPost(postParams:Post,token:string){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/post/${token}`, {
        method:"POST",
        body:JSON.stringify(postParams)
    }).then(async (res:any) => {
        let res_json = await res.json()
        if (Object.hasOwn(res_json, 'error')){
            return false
        } else {
            return true
        }
    }).catch(()=>{
        return false
    })
}

export async function removePost(id:string, token:string,username:string){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/post/${username}/${id}/${token}`, {
        method:"GET"
    }).then(async (res:any) => {
        let res_json = await res.json()
        if (Object.hasOwn(res_json, 'error')){
            return false
        } else {
            return true
        }
    }).catch(()=>{
        return false
    })
}

export function getPosts(searchParameters?:string,page:string="1",itemsPerPage:string="20"){
    let uri = (searchParameters && searchParameters.length>0 ? `${process.env.REACT_APP_BACKEND_URI}/post?searchParams=${searchParameters}&itemsPerPage=${itemsPerPage}&page=${page}` : `${process.env.REACT_APP_BACKEND_URI}/post`)
    return fetch(uri)
}

export function getPostsByUsername(username:string){
    let uri = `${process.env.REACT_APP_BACKEND_URI}/post/${username}`
    return fetch(uri)
}

export function getFeedByUsername(username:string){
    let uri = `${process.env.REACT_APP_BACKEND_URI}/feed/friends/${username}`
    return fetch(uri)
}

export function toggleLikedPost(id:string, token:string,username:string){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/post/like/${id}/${username}/${token}`, {
        method:"POST"
    })
}

export function getLikedPostsByUsername(username:string){
    let uri = `${process.env.REACT_APP_BACKEND_URI}/post/like/${username}`
    return fetch(uri)
}