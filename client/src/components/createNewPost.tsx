import { Button, Card,Text, Textarea, TextInput, Title } from "@mantine/core"
import { showNotification } from "@mantine/notifications"
import { PostContainer } from "../pages/Home"

export const CreateNewPost = () => {


    return (
        <>
        <Card radius="md" withBorder style={{marginTop:"30px",overflow:"hidden"}}>
            <Title p="xs" order={4}>New Post</Title>
            <TextInput placeholder="Title" disabled p="xs"></TextInput>
            <Textarea placeholder="Content" disabled p="xs"></Textarea>

            <Button m="xs" disabled>Post</Button>

        </Card>
        </>
    )
}
