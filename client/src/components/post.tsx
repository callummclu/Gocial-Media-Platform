import { Title, Text, Card } from "@mantine/core"
import styled from "styled-components"
import ReactMarkdown from 'react-markdown'
import dayjs from "dayjs"
import relativeTime from "dayjs/plugin/relativeTime"


export const Post = ({title,username,content,created_at}:any) => {

    dayjs.extend(relativeTime)

    let formattedDate = dayjs(created_at).add(-1,'hour').fromNow()
    
    return (
        <Card radius="md" withBorder style={{marginTop:"30px"}}>
        <PostContainer>
            <div style={{display:"flex",alignItems:"center",gap:"10px"}}><Title order={2}>{title}</Title><Text className="created_at">{formattedDate.toString()}</Text></div>
            <Text className="username">{username}</Text>
            <Text className="content">
                <ReactMarkdown>{content}</ReactMarkdown>
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