import { Container, Title, Text, Divider, Group, Card, Avatar } from "@mantine/core"
import { useEffect, useState } from "react"
import { useParams, useSearchParams } from "react-router-dom"
import { Post } from "../components/post"
import { PostContainer } from "./Home"

function UserProfile(props:any){

    const {username} = useParams()

    const [updatePosts, setUpdatePosts] = props.updatePosts

    const [posts,setPosts] = useState<any>()
    const [friends,setFriends] = useState<any>()

    useEffect(()=>{

        let uri = `${process.env.REACT_APP_BACKEND_URI}/post/${username}`

        fetch(uri)
            .then(async (res:any) => {
                let res_json = await res.json()
                setPosts(res_json)
            })
    },[updatePosts])


    return (
        <>
           <Container>
            <div style={{display:"flex",gap:"30px"}}>
                <Card style={{width:"70%"}}>
                    <Group mb={"xl"}><Avatar size={128} radius={100} /><div><Title ml="md">{username}</Title><Text ml="md">{"Test Description"}</Text></div></Group>
                    <Divider/>
                    <PostContainer>
                    {(posts?.data) ? (posts.data).map((e:any)=><Post {...e} updatePosts={[updatePosts, setUpdatePosts]} loggedInUser={username} />) :<Text m="xl">nothing to show.</Text>}
                    </PostContainer>
                </Card>
                <Card style={{width:"30%",height:"calc(100vh - 110px)", position:"sticky",top:"50px"}} withBorder >
                    <Title m="xl" order={2}>Friends</Title>
                    {friends ? <Card>friend</Card> : <Card ml="md"><Text>no friends :(</Text></Card>}
                </Card>

            </div>
        </Container>
        </>
    )
}

export default UserProfile