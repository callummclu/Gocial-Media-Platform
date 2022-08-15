import { Modal, Affix, Button, Transition, TextInput, Textarea } from "@mantine/core"
import { showNotification } from "@mantine/notifications";
import { useRef, useState } from "react";
import {HiOutlineDocumentAdd} from 'react-icons/hi'
import { newPostPost } from "../helpers/postHelper";
import { Post } from "../Types/post";

export const NewPost = (props:any) => {

    const [opened, setOpened] = useState(false);
    const [updatePosts, setUpdatePosts] = props.updatePosts

    const titleRef = useRef<HTMLInputElement>(null)
    const contentRef = useRef<HTMLTextAreaElement>(null)

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
              setOpened(false)
              setUpdatePosts(!updatePosts)

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
         <Modal
        opened={opened}
        onClose={() => setOpened(false)}
        title="Create a new Post!"
      >
        <form onSubmit={newPost}>
          <TextInput label="Title" ref={titleRef} type="name" required/>
          <Textarea label="Content" ref={contentRef} required/>
          <Button mt="sm" type="submit">Post</Button>
        </form>
      </Modal>
        {!opened && <Affix position={{ bottom: 20, right: 20 }}>
        <Button
            leftIcon={<HiOutlineDocumentAdd size={16} />}
            onClick={()=>setOpened(true)}
        >
            New Post
        </Button>

      </Affix>}
      </>
    )
}