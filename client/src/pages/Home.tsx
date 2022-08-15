import { Container, Divider, Title, Text } from "@mantine/core"
import { useEffect, useState } from "react"
import { useSearchParams } from "react-router-dom"
import styled from "styled-components"
import { Post } from "../components/post"

function Home(props:any){
    const [searchParams,setSearchParams] = useSearchParams()
    const [posts,setPosts] = useState<any>()
    const [updatePosts, setUpdatePosts] = props.updatePosts
    const [page, setPage] = useState<number>(1)
    const [username,setUsername] = props.username

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
                <PostContainer>
                    {(posts?.data) ? (posts.data).map((e:any)=><Post {...e} updatePosts={[updatePosts, setUpdatePosts]} loggedInUser={username} />) :<Text m="xl">nothing to show.</Text>}
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