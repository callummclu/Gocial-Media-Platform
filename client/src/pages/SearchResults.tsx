import { Card, Container, Title, Text, Pagination, Group, Divider } from "@mantine/core"
import { useEffect,useState } from "react"
import { useSearchParams } from "react-router-dom"
import internal from "stream"
import styled from "styled-components"

const SearchResults = () => {
    const [searchParams,setSearchParams] = useSearchParams()
    const [users,setUsers] = useState<any>()
    const [page, setPage] = useState<number>(1)

    useEffect(()=>{
        let searchParameters = searchParams.get("searchParams")
        fetch(`${process.env.REACT_APP_BACKEND_URI}/user?searchParams=${searchParameters}&itemsPerPage=20&page=${page}`)
            .then(async (res:any) => {
                let res_json = await res.json()
                setUsers(res_json)
            })
    },[searchParams,page])

    return (
        <>
        <Container >
            <Title m="xl">Search Results</Title>
            <Divider variant="dotted" labelPosition="center" label={`${users?.results ?? 0} results`}/>
            {users ? (users.data || users).map((e:any)=>(
            <Card m="xl" shadow="sm" p="lg" radius="md" withBorder key={e.username}><Text><a href={`/users/${e.username}`}>{e.username}</a> - {e.name} {e.surname}</Text></Card>)
            ): "no results"}
            <Group style={{display:"flex",flexDirection:"column",alignItems:"center","justifyContent":"center"}}>
                <Pagination page={page} total={users?.pages ?? 0} onChange={setPage}/>
            </Group>
        </Container>
        </>
    )
}

export default SearchResults