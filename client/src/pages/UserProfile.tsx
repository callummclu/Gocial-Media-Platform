import { Container, Title, Text, Divider, Group, Card, Avatar, Button } from "@mantine/core"
import { showNotification } from "@mantine/notifications"
import { useEffect, useState } from "react"
import { useParams, useSearchParams } from "react-router-dom"
import styled from "styled-components"
import { resolveModuleNameFromCache } from "typescript"
import { Post } from "../components/post"
import { PostContainer } from "./Home"

function UserProfile(props:any){

    const {username} = useParams()

    const [updatePosts, setUpdatePosts] = props.updatePosts
    const [userData,setUserData] = useState<any>(null)
    const [loggedIn,setLoggedIn] = props.loggedIn
    const [loggedInUserData, setLoggedInUserData] = props.loggedInUserData
    const [posts,setPosts] = useState<any>()
    useEffect(()=>{

        let uri = `${process.env.REACT_APP_BACKEND_URI}/post/${username}`

        fetch(uri)
            .then(async (res:any) => {
                let res_json = await res.json()
                setPosts(res_json)
            })
            fetch(`${process.env.REACT_APP_BACKEND_URI}/user/${loggedInUserData?.name}`)
            .then(async (res:any) => {
                let res_json = await res.json()
                setLoggedInUserData(res_json.data)
            })
    },[updatePosts])


    useEffect(()=>{
        fetch(`${process.env.REACT_APP_BACKEND_URI}/user/${username}`)
            .then(async (res:any) => {
                let res_json = await res.json()
                setUserData(res_json.data)
            })
    },[])
    const redirectToUser = (username:string) => {
        window.location.href = window.location.origin + "/users/" + username 
    }


    const checkUserNotSelfOrFriend = () => {
        let SameUser = (loggedInUserData?.name) == username
        let InFriends = ((loggedInUserData?.friends ?? []).includes(username))
        return !(SameUser || InFriends)
    }

    const checkUserIsSelf = () => {
        let sameUser = (loggedInUserData?.name) == username
        return (loggedIn && sameUser)
    }

    const addUser = () => {
        fetch(`${process.env.REACT_APP_BACKEND_URI}/user/invitation/${loggedInUserData?.name}/${username}/${localStorage.getItem("gocial_auth_token")}`,{
            method:"POST"
        }).then(async (res:any) => {
            let res_json = await res.json()

            if(res_json?.error ?? false){
                showNotification({
                    title:"Error",
                    message:"Already sent a request",
                    color:"red"
                  })
            } else {
                showNotification({
                    title:"Success",
                    message:"request sent successfully",
                    color:"green"
                  })
            }
        })
    }

    const acceptRequest = (req_user:string) => {
        fetch(`${process.env.REACT_APP_BACKEND_URI}/user/friends/${loggedInUserData?.name}/${req_user}/${localStorage.getItem("gocial_auth_token")}/accept`,{
            method:"POST"
        }).then(async (res:any) => {
            let res_json = await res.json()

            if(res_json?.error ?? false){
                showNotification({
                    title:"Error",
                    message:"Cannot add user",
                    color:"red"
                  })
                  
            } else {
                showNotification({
                    title:"Success",
                    message:"Added User",
                    color:"green"
                  })
            }
        })
    }

    return (
        <>
           <Container>
            <div style={{display:"flex",gap:"30px"}}>
                <Card style={{width:"70%"}}>
                    <div style={{display:"flex",alignItems:"center",justifyContent:"space-between"}}><Group mb={"xl"}><Avatar src={userData?.display_image} size={128} radius={100} /><div><Title ml="md">{username}</Title><Text ml="md">{userData?.description ?? ""}</Text></div></Group>{(loggedIn && checkUserNotSelfOrFriend()) && <Button onClick={addUser}>Add</Button>}</div>
                    <Divider/>
                    <PostContainer>
                    {(posts?.data) ? (posts.data).map((e:any)=><Post {...e} updatePosts={[updatePosts, setUpdatePosts]} loggedInUser={username} key={username+e.title}/>) :<Text m="xl">nothing to show.</Text>}
                    </PostContainer>
                </Card>
                <Card style={{width:"30%",height:"calc(100vh - 110px)", position:"sticky",top:"50px"}} withBorder >
                    <Title m="xl" order={2}>Friends</Title>
                    <FriendsContainer>{userData?.friends ? userData.friends.map((e:any)=><Card onClick={()=>redirectToUser(e)} className="user" p="lg" radius="md" key={e}><Group><Avatar radius="xl" /><Title order={3}>{e}</Title></Group></Card>) : <Card ml="md"><Text>no friends :(</Text></Card>}</FriendsContainer>
                    {(checkUserIsSelf() && (userData?.received_invitations && userData?.received_invitations.length > 0)) &&<><Divider m="md"/><Title m="xl" order={2}>Requests</Title><FriendsContainer>{userData?.received_invitations ? userData?.received_invitations.map((e:any)=><Card className="user" p="lg" radius="md" key={e}><Group><Avatar radius="xl" /><Title onClick={()=>redirectToUser(e)} order={3}>{e}</Title><Button onClick={()=>acceptRequest(e)}>Accept</Button></Group></Card>) : <Card ml="md"><Text>no friends :(</Text></Card>}</FriendsContainer></>}
                </Card>

            </div>
        </Container>
        </>
    )
}

export default UserProfile


const FriendsContainer = styled.div`

    & .user{
        &:hover{
            background:rgba(0,0,0,0.05);
            cursor:pointer;
        }
    }
`