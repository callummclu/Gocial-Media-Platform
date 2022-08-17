import React, {
    createContext,
    ReactNode,
    useContext,
    useEffect,
    useMemo,
    useState,
  } from "react";

import * as PostApi from '../api/posts'
import { Post } from "../Types/post";

interface PostContextType {
    loading: boolean;
    error?: any;
    newPost: (postParmas:Post) => any;
    getPosts: (searchParams?:string, itemsPerPage?:string, page?:string) => any;
    posts: Post[]
    deletePost: (id:string,username:string) => void;
    getPostsByUsername: (username:string) => void;
}

const PostContext = createContext<PostContextType>({} as PostContextType)

export function PostProvider({children}:{children:ReactNode}):JSX.Element {

    const [posts, setPosts] = useState<Post[]>([])
    const [error,setError] = useState<any>()
    const [loading, setLoading] = useState<boolean>(false)

    function newPost(postParmas:Post){
        setLoading(true)
        let token:string = localStorage.getItem("gocial_auth_token") as string
        PostApi.newPostPost(postParmas,token)
            .then(async (res:any) => {
                let res_json = await res.json()
                if (!Object.hasOwn(res_json, 'error')){
                    // if url is home getPosts
                    // if url is /users/{username} getPostsByUsername
                } else {
                    setError(res_json.error)
                }
            })
            .catch((error)=>{})
            .finally(()=>setLoading(false))
    }
    function getPosts(searchParams?:string, itemsPerPage?:string, page?:string){
        setLoading(true)
        PostApi.getPosts()
            .then(async (res:any) => {
                let res_json = await res.json()
                setPosts(res_json.data)
            })
            .catch((error)=>{})
            .finally(()=>{
                setLoading(false)
            })
    }

    function getPostsByUsername(username:string){
        setLoading(true)
        PostApi.getPostsByUsername(username)
            .then(async (res:any) => {
                let res_json = await res.json()
                setPosts(res_json.data)
            })
            .catch((error)=>{})
            .finally(()=>{
                setLoading(false)
            })
    }

    function deletePost(id:string,username:string){
        setLoading(true)
        let token:string = localStorage.getItem("gocial_auth_token") as string

        PostApi.removePost(id,token,username)
            .then(async (res:any) => {
                let res_json = await res.json()
                if (Object.hasOwn(res_json, 'error')){
                    setError(res_json.error)
                } else {
                    // if url is home getPosts
                    // if url is /users/{username} getPostsByUsername
                }
            }).catch(()=>{
                return false
            }).finally(()=>{
                setLoading(false)
            })


    }

    const memoedValue = useMemo(
        () => ({
            error,
            loading,
            newPost,
            getPosts,
            posts,
            deletePost,
            getPostsByUsername

        }),
        [loading,error]
    )

    return (
        <PostContext.Provider value={memoedValue}>
            {children}
        </PostContext.Provider>
    )
}

export default function usePost() {
    return useContext(PostContext)
}