import { Title, Text, Card, Button } from "@mantine/core"
import styled from "styled-components"
import ReactMarkdown from 'react-markdown'
import dayjs from "dayjs"
import relativeTime from "dayjs/plugin/relativeTime"
import {removePost} from '../helpers/postHelper'


export const Post = (props:any) => {

    const [updatePosts, setUpdatePosts] = props.updatePosts

    dayjs.extend(relativeTime)

    let formattedDate = dayjs(props.created_at).add(-1,'hour').fromNow()


    const deletePost = () => {
        let token = localStorage.getItem("gocial_auth_token") || ""

        removePost(props.id,token,props.loggedInUser)
        setUpdatePosts(!updatePosts)

    }
    
    return (
        <Card radius="md" withBorder style={{marginTop:"30px"}}>
        <PostContainer>
            <div style={{display:"flex",alignItems:"center"}}>
            <div style={{width:"calc(100% - 100px)"}}>
                <div style={{display:"flex",alignItems:"center",gap:"10px"}}><Title order={2}>{props.title}</Title><Text className="created_at">{formattedDate.toString()}</Text></div>
                <Text className="username">{props.username}</Text>
            </div>
            {props.loggedInUser === props.username && <Button onClick={deletePost} color="red">Delete Post</Button>}
            </div>
            <Text className="content">
                <ReactMarkdown>{props.content}</ReactMarkdown>
            </Text>
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