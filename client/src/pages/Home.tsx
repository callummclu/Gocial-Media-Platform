import { Container, Divider, Title, Text } from "@mantine/core"

function Home(){
    return (
        <>
           <Container>
            <Title m="xl">Feed</Title>
            <Divider/>
            <Container>
                <Text m="xl">nothing to show.</Text>
            </Container>
        </Container>
        </>
    )
}

export default Home