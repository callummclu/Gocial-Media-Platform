import { Card, Container, Title, Text, Pagination, Group, Divider, Avatar } from "@mantine/core"
import { useEffect,useState } from "react"
import { useSearchParams } from "react-router-dom"
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

    const redirectToUser = (username:string) => {
        window.location.href = window.location.origin + "/users/" + username 
    }

    return (
        <UserResultContainer>
        <Container >
            <Title m="xl">Search Results</Title>
            <Divider variant="dotted" labelPosition="center" label={`${users?.results ?? 0} results`}/>
            <br/>
            {users ? (users.data || users).map((e:any)=>(
            <><Card onClick={()=>redirectToUser(e.username)} className="user" p="lg" radius="md" key={e.username}><Group><Avatar radius="xl" /><div><Title order={3}>{e.username}</Title><Text>Test Description</Text></div></Group></Card></>)
            ): "no results"}
            <br/>
            <Group style={{display:"flex",flexDirection:"column",alignItems:"center","justifyContent":"center"}}>
                <Pagination page={page} total={users?.pages ?? 0} onChange={setPage}/>
            </Group>
        </Container>
        </UserResultContainer>
    )
}

export default SearchResults

const UserResultContainer = styled.div`
    & .user{
        &:hover{
            background:rgba(0,0,0,0.05);
            cursor:pointer;
        }
    }
`