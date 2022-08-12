import { Card, Container, Title, Text } from "@mantine/core"
import { useEffect,useState } from "react"
import { useSearchParams } from "react-router-dom"

const SearchResults = () => {
    const [searchParams,setSearchParams] = useSearchParams()
    const [users,setUsers] = useState<any>()

    useEffect(()=>{
        let searchParameters = searchParams.get("searchParams")
        fetch(`${process.env.REACT_APP_BACKEND_URI}/user?searchParams=${searchParameters}`)
            .then(async (res:any) => {
                let res_json = await res.json()
                setUsers(res_json.data)
            })
    },[searchParams])

    return (
        <>
        <Container>
            <Title m="xl">Search Results</Title>
            {users ? users.map((e:any)=>(
            <Card m="xl" shadow="sm" p="lg" radius="md" withBorder key={e.username}><Text>{e.username} - {e.name} {e.surname}</Text></Card>)
            ): "no results"}
        </Container>
        </>
    )
}

export default SearchResults