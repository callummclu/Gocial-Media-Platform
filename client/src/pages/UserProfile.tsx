import { Container, Title, Text, Divider, Group, Card, Avatar, Button, SegmentedControl } from "@mantine/core"
import { showNotification } from "@mantine/notifications"
import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { PuffLoader } from "react-spinners"
import styled from "styled-components"
import { Post } from "../components/post"
import { getLikedPostsByUsername, getPostsByUsername } from "../helpers/postHelper"
import { acceptFriendRequest, getSingleUser, sendInvitation } from "../helpers/userHelper"
import useAuth from "../hooks/useAuth"
import { User } from "../Types/auth"
import { PostContainer } from "./Home"

function UserProfile(props:any){

    const {username} = useParams()

    const {user,loggedIn} = useAuth()

    const [feedSection, setFeedSection] = useState("posts")


    const [updatePosts, setUpdatePosts] = props.updatePosts
    const [userData,setUserData] = useState<any>(null)
    const [posts,setPosts] = useState<any>()
    
    useEffect(()=>{
        if(feedSection == "posts"){
            getPostsByUsername(username as string)
             .then(async (res:any) => {
                 let res_json = await res.json()
                 setPosts(res_json)
             })
         } else {
             getLikedPostsByUsername(username as string)
                 .then(async (res:any) => {
                     let res_json = await res.json()
                     setPosts(res_json)
                 })
         }
    },[updatePosts,feedSection])


    useEffect(()=>{
        getSingleUser(username as string)
        .then(async (res:any) => {
            let res_json = await res.json()
            setUserData(res_json.data)
        })
    },[feedSection])

    const redirectToUser = (username:string) => {
        window.location.href = window.location.origin + "/users/" + username 
    }


    const checkUserNotSelfOrFriend = () => {
        let SameUser = (user?.username) === username
        let InFriends = (((user?.friends ?? [])).includes(username || "21321321321"))
        return !(SameUser || InFriends)
    }

    const checkUserIsSelf = () => {
        let sameUser = (user?.username) === username
        return (loggedIn && sameUser)
    }

    const addUser = () => {
        sendInvitation(user as User, username as string)
            .then(async (res:any) => {
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
        acceptFriendRequest(user as User, req_user)
            .then(async (res:any) => {
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
        setUpdatePosts(!updatePosts)
    }

    const switchFeed = (e:any) => {
        setFeedSection(e)
    }

    return (
        <>
           <Container>
            <div style={{display:"flex",gap:"30px"}}>
                <Card style={{width:"70%"}}>
                    <div style={{display:"flex",alignItems:"center",justifyContent:"space-between"}}><Group mb={"xl"}><Avatar src={userData?.display_image} size={128} radius={100} /><div><Title ml="md">{username}</Title><Text ml="md">{userData?.description ?? ""}</Text></div></Group>{(loggedIn && checkUserNotSelfOrFriend()) && <Button onClick={addUser}>Add</Button>}</div>

                    {loggedIn && <SegmentedControl
                    onChange={switchFeed}
                    mt="xl"
                    fullWidth
                    defaultValue={feedSection}
                    data={[
                        { label: 'Posts', value: 'posts' },
                        { label: 'Likes', value: 'likes' }
                    ]}
                    />}
                <PostContainer></PostContainer>
                    <PostContainer>
                    {(posts?.data) ? (posts.data).map((e:any)=><Post {...e} updatePosts={[updatePosts, setUpdatePosts]} key={username+e.title}/>) :<div style={{width:"100%",height:"calc(100vh - 110px)",display:"flex",alignItems:"center",justifyContent:"center"}}><PuffLoader color="gray" size={20}/></div>}
                    </PostContainer>
                </Card>
                <Card style={{width:"30%",height:"calc(100vh - 110px)", position:"sticky",top:"50px"}} withBorder >
                    <Title m="xl" order={2}>Friends ({(userData?.friends ?? []).length || 0})</Title>
                    {checkUserIsSelf() ?
                    <FriendsContainer>{user?.friends ? user.friends.map((e:any)=><Card onClick={()=>redirectToUser(e)} className="user" p="sm" radius="md" key={e}><Group><Avatar radius="xl" /><Title order={5}>{e}</Title></Group></Card>) : <Card ml="md"><Text>nothing to show</Text></Card>}</FriendsContainer>
                    :
                    <FriendsContainer>{userData?.friends ? userData.friends.map((e:any)=><Card onClick={()=>redirectToUser(e)} className="user" p="sm" radius="md" key={e}><Group><Avatar radius="xl" /><Title order={5}>{e}</Title></Group></Card>) : <Card ml="md"><Text>nothing to show</Text></Card>}</FriendsContainer>
                    }
                    {(checkUserIsSelf() 
                        && (user?.received_invitations && user?.received_invitations.length > 0)) 
                        &&<>
                        <Divider m="md"/>
                        <Title m="xl" order={2}>Requests</Title>
                        <FriendsContainer>
                            {user?.received_invitations.map((e:any)=><Card className="user" p="sm" radius="md" key={e}><Group><Avatar radius="xl" /><Title onClick={()=>redirectToUser(e)} order={5}>{e}</Title><Button onClick={()=>acceptRequest(e)}>Accept</Button></Group></Card>)}
                        </FriendsContainer></>}
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