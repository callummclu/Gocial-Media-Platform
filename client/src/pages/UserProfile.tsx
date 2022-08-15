import { Container, Title, Text, Divider } from "@mantine/core"
import { useEffect, useState } from "react"
import { useParams, useSearchParams } from "react-router-dom"
import { Post } from "../components/post"
import { PostContainer } from "./Home"

function UserProfile(){

    const {username} = useParams()

    const [posts,setPosts] = useState<any>()

    useEffect(()=>{

        let uri = `${process.env.REACT_APP_BACKEND_URI}/post/${username}`

        fetch(uri)
            .then(async (res:any) => {
                let res_json = await res.json()
                setPosts(res_json)
            })
    },[])


    return (
        <>
           <Container>
            <Title m="xl">{username}</Title>
            <Divider/>
            <PostContainer>
            {(posts?.data) ? (posts.data.reverse()).map((e:any)=><Post {...e}/>) :<Text m="xl">nothing to show.</Text>}
            </PostContainer>
        </Container>
        </>
    )
}

export default UserProfile