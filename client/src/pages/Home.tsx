import { Container, Divider, Title, Text } from "@mantine/core"
import { useEffect, useState } from "react"
import { useSearchParams } from "react-router-dom"
import { ClimbingBoxLoader, PuffLoader } from "react-spinners"
import styled from "styled-components"
import { CreateNewPost } from "../components/createNewPost"
import { Post } from "../components/post"
import useAuth from "../hooks/useAuth"

function Home(props:any){
    const [searchParams,setSearchParams] = useSearchParams()
    const [posts,setPosts] = useState<any>()
    const [updatePosts, setUpdatePosts] = props.updatePosts
    const [page, setPage] = useState<number>(1)
    const {user,loggedIn} = useAuth()

    useEffect(()=>{
        let searchParameters = searchParams.get("searchParams")

        let uri = (searchParameters && searchParameters.length>0 ? `${process.env.REACT_APP_BACKEND_URI}/post?searchParams=${searchParameters}&itemsPerPage=20&page=${page}` : `${process.env.REACT_APP_BACKEND_URI}/post`)

        fetch(uri)
            .then(async (res:any) => {
                let res_json = await res.json()
                setPosts(res_json)
            })
    },[searchParams,page,updatePosts])
    
    return (
        <>
        <Container>
            <Container>
                {loggedIn && <CreateNewPost updatePosts={[updatePosts, setUpdatePosts]}/>}
                <PostContainer>
                    {(posts?.data) ? (posts.data).map((e:any)=><Post {...e} updatePosts={[updatePosts, setUpdatePosts]} key={user?.username+e.title}/>) :<div style={{width:"100%",height:"calc(100vh - 110px)",display:"flex",alignItems:"center",justifyContent:"center"}}><PuffLoader color="gray" size={20}/></div>}
                </PostContainer>
            </Container>
        </Container>
        </>
    )
}

export default Home

export const PostContainer = styled.div`
    margin-bottom:50px;
`