import { Title, Text, Card, Button, Menu, Chip, Divider, UnstyledButton } from "@mantine/core"
import styled from "styled-components"
import ReactMarkdown from 'react-markdown'
import dayjs from "dayjs"
import relativeTime from "dayjs/plugin/relativeTime"
import {removePost, toggleLikedPost} from '../helpers/postHelper'
import {BsThreeDots} from 'react-icons/bs'
import {BiTrash} from 'react-icons/bi'
import {AiFillLike,AiOutlineLike} from 'react-icons/ai'
import useAuth from "../hooks/useAuth"
import { useEffect, useState } from "react"
import { showNotification } from "@mantine/notifications"

export const Post = (props:any) => {

    const [updatePosts, setUpdatePosts] = props.updatePosts
    const {loggedIn,user} = useAuth()
    const [likeToggle, setLikesToggle] = useState(false)

    dayjs.extend(relativeTime)

    let formattedDate = dayjs(props.created_at).add(-1,'hour').fromNow()


    const deletePost = () => {
        let token = localStorage.getItem("gocial_auth_token") || ""

        removePost(props.id,token,user?.username ?? "")
        setUpdatePosts(!updatePosts)

    }

    const togglePostLike = () => {
        let token = localStorage.getItem("gocial_auth_token") || ""
        loggedIn ? (
        toggleLikedPost(props.id,token,user?.username ?? "")
        .then(()=>{
            setUpdatePosts(!updatePosts)
            setLikesToggle(!likeToggle)
        })
        ) : (
            showNotification({
                title:"Error - Log in",
                message:"You need to be logged in to like posts.",
                color:"red"
              })
        )
    }

    const redirectToUser = (username:string) => {
        window.location.href = window.location.origin + "/users/" + username 
    }


    useEffect(()=>{
        let likes_arr = props.likes
        setLikesToggle((likes_arr ?? []).includes(user?.username))
    },[])
    
    return (
        <Card radius="md" withBorder style={{marginTop:"30px",overflow:"hidden"}}>
        <PostContainer>
        {user?.username === props.username && <div style={{position:"absolute",top:25,right:25}}>
            <Menu position="left-start" width={130}>
                <Menu.Target>
                    <div>
                        <BsThreeDots style={{cursor:"pointer"}} color="gray"/>
                    </div>
                </Menu.Target>

                <Menu.Dropdown style={{right:"0"}}>
                    <Menu.Item color="red" onClick={deletePost} icon={<BiTrash size={14} />}>&nbsp;&nbsp;Delete Post</Menu.Item>
                </Menu.Dropdown>
            </Menu>

            </div>}
            <div style={{display:"flex",alignItems:"center"}}>
            <div style={{width:"calc(100% - 100px)"}}>
                <div style={{display:"flex",alignItems:"center",gap:"10px",flexWrap:"wrap"}}><Title order={4}>{props.title}</Title><Text style={{lineHeight:0}} pt="sm" pb="sm" className="created_at">{formattedDate.toString()}</Text></div>
                <Text className="username" onClick={()=>redirectToUser(props.username)}>{props.username}</Text>
            </div>
            </div>
            <Text className="content">
                <ReactMarkdown>{props.content}</ReactMarkdown>
            </Text>
            <Divider/>
            <Text pt="sm" style={{display:"flex",alignItems:"center",fontSize:"12px",gap:"10px"}}><UnstyledButton onClick={togglePostLike}>{likeToggle ? <AiFillLike size={24}/> : <AiOutlineLike size={24}/>}</UnstyledButton>{(props.likes)?.length ?? 0}</Text>
        </PostContainer>
        </Card>
    )
}

const PostContainer = styled.div`

    & *{
        margin:0
    }


    & .username{

        color:gray;
        &::before{
            content:"@";
        }
    }

    & .created_at{
        font-size:12px;
        color:gray;
    }

    & .content{
        padding-top:10px;
        padding-bottom:10px;

        & h1{font-size: 24px;}
        & h2{font-size: 20px;}
        & h3{font-size: 18px;}
        & h4{font-size: 16px;}
        & h5{font-size: 14px;}
        & h6{font-size: 12px;}
        & p{font-size: 14px;}
    }

`