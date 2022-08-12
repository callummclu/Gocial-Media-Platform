import { Container, Title } from "@mantine/core"
import { useParams } from "react-router-dom"

function UserProfile(){

    const {username} = useParams()

    return (
        <>
           <Container>
            <Title m="xl">{username}</Title>
            
        </Container>
        </>
    )
}

export default UserProfile