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