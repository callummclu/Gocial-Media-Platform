import { Post } from "../Types/post";

export async function newPostPost(postParams:Post,token:string){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/post/${token}`, {
        method:"POST",
        body:JSON.stringify(postParams)
    })
}

export async function removePost(id:string, token:string,username:string){
    return fetch(`${process.env.REACT_APP_BACKEND_URI}/post/${username}/${id}/${token}`, {
        method:"GET"
    })
}

export function getPosts(searchParameters?:string,page:string="1",itemsPerPage:string="20"){
    let uri = (searchParameters && searchParameters.length>0 ? `${process.env.REACT_APP_BACKEND_URI}/post?searchParams=${searchParameters}&itemsPerPage=${itemsPerPage}&page=${page}` : `${process.env.REACT_APP_BACKEND_URI}/post`)
    const response = fetch(uri)
    return response
}

export function getPostsByUsername(username:string){
    let uri = `${process.env.REACT_APP_BACKEND_URI}/post/${username}`
    const response = fetch(uri)
    return response
}