import React, {
    createContext,
    ReactNode,
    useContext,
    useEffect,
    useMemo,
    useState,
  } from "react";

interface PostContextType {
    loading: boolean;
    error?: any;
    // newPost: () => any
    // userPosts: (username:string) => any
    // getAllPosts: () => any 
    // deletePost: () => void
}

const PostContext = createContext<PostContextType>({} as PostContextType)

export function PostProvider({children}:{children:ReactNode}):JSX.Element {

    const [error,setError] = useState<any>()
    const [loading, setLoading] = useState<boolean>(false)
    const [loadingInitial, setLoadingInitial] = useState<boolean>(true)

    // new post
    // get postByUser
    // get AllPosts
    // delete Post

    const memoedValue = useMemo(
        () => ({
            error,
            loading
        }),
        [loading,error]
    )

    return (
        <PostContext.Provider value={memoedValue}>
            {!loadingInitial && children}
        </PostContext.Provider>
    )
}

export default function usePost() {
    return useContext(PostContext)
}