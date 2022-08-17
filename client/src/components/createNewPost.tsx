import { Button, Card,Text, Textarea, TextInput, Title } from "@mantine/core"
import { showNotification } from "@mantine/notifications"
import { useRef } from "react"
import { newPostPost } from "../helpers/postHelper"
import { PostContainer } from "../pages/Home"
import { Post } from "../Types/post"

export const CreateNewPost = (props:any) => {

    const titleRef = useRef<HTMLInputElement>(null)
    const contentRef = useRef<HTMLTextAreaElement>(null)
    const [updatePosts, setUpdatePosts] = props.updatePosts

    const newPost = async (e:any) => {
        e.preventDefault()
        let postParams: Post = {
            title: titleRef.current!.value,
            content: contentRef.current!.value
        }

        let token = localStorage.getItem("gocial_auth_token") || ""

        let data = await newPostPost(postParams,token)

        if (data){
            showNotification({
                title:"Congrats",
                message:"You've made a post",
                color:"green"
              })
              setUpdatePosts(!updatePosts)
              titleRef.current!.value = ""
              contentRef.current!.value = ""

        } else {
            showNotification({
                title:"Error",
                message:"Something happened, try again.",
                color:"red"
              })
        }
    }
    return (
        <>
        <Card radius="md" withBorder style={{marginTop:"30px",overflow:"hidden"}}>
            <Title p="xs" order={4}>New Post</Title>
            <form onSubmit={newPost}>
                <TextInput ref={titleRef} placeholder="Title" p="xs"></TextInput>
                <Textarea ref={contentRef} placeholder="Content" p="xs"></Textarea>

                <Button type="submit" m="xs">Post</Button>
            </form>
        </Card>
        </>
    )
}
