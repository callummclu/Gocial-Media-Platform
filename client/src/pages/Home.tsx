import { Container, Divider, Title, Text, SegmentedControl, Center, TextInput } from "@mantine/core"
import { useEffect, useRef, useState } from "react"
import { useSearchParams } from "react-router-dom"
import { ClimbingBoxLoader, PuffLoader } from "react-spinners"
import styled from "styled-components"
import { CreateNewPost } from "../components/createNewPost"
import { Post } from "../components/post"
import { getFeedByUsername, getPosts } from "../helpers/postHelper"
import useAuth from "../hooks/useAuth"

function Home(props:any){
    const [searchParams,setSearchParams] = useSearchParams()
    const [posts,setPosts] = useState<any>()
    const [updatePosts, setUpdatePosts] = props.updatePosts
    const [page, setPage] = useState<number>(1)
    const {user,loggedIn} = useAuth()
    const [feedSection, setFeedSection] = useState("explore")

    useEffect(()=>{
        let searchParameters = searchParams.get("searchParams")
        if(feedSection == "explore"){
            getPosts(searchParameters as string,"1","20")
            .then(async (res:any) => {
                let res_json = await res.json()
                setPosts(res_json)
            })
        } else {
            getFeedByUsername(user?.username as string)
            .then(async (res:any) => {
                let res_json = await res.json()
                setPosts(res_json)
            })

            setPosts(null)
        }
        
    },[searchParams,page,updatePosts,feedSection])
    
    const switchFeed = (e:any) => {
        setFeedSection(e)
    }

    return (
        <>
        <Container>
            <Container>
                {loggedIn && <CreateNewPost updatePosts={[updatePosts, setUpdatePosts]}/>}
                {loggedIn && <SegmentedControl
                    onChange={switchFeed}
                    mt="xl"
                    fullWidth
                    value={feedSection}
                    data={[
                        { label: 'Friends', value: 'friends' },
                        { label: 'Explore', value: 'explore' }
                    ]}
                    />}
                <PostContainer>
                
                    {(posts?.data) ? (posts.data).map((e:any)=><Post {...e} updatePosts={[updatePosts, setUpdatePosts]} key={user?.username+e.title+e.content}/>) :<div style={{width:"100%",height:"calc(100vh - 110px)",display:"flex",alignItems:"center",justifyContent:"center"}}><PuffLoader color="gray" size={20}/></div>}
                </PostContainer>
            </Container>
        </Container>
        </>
    )
}

export default Home

export const PostContainer = styled.div`
    margin-bottom:50px;
`