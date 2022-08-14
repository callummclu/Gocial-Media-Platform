import { Button, Container, Divider, Group, TextInput, Title } from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import { useRef } from 'react'
import { useSearchParams } from 'react-router-dom'
import styled from 'styled-components'
import { logOut } from '../helpers/authHelper'

export const Navbar = (props:any) => {

    const searchQueryRef = useRef<HTMLInputElement>(null)
    const [searchParams,setSearchParams] = useSearchParams()
    const [loggedIn,setLoggedIn] = props.loggedIn 
    const [username,setUsername] = props.username

    const searchUsersSubmit = (e:any) => {
        e.preventDefault()
        if (window.location.pathname !== '/results'){
            window.location.href = window.location.origin + `/results?searchParams=${searchQueryRef.current?.value}`
        } else {
            setSearchParams({"searchParams":`${searchQueryRef.current?.value}`})
        }
    }

    const logOutHandler = async () => {
        setLoggedIn(await logOut())
        setUsername("")
        showNotification({
          title:"Success",
          message:"You've logged out successfully"
        })
      }

      const logInHandler = () => {
        window.location.href = window.location.origin + '/login'
      }

    return (
        <>
        <NavbarStyled>
            <Title onClick={()=>window.location.href = window.location.origin}>Gocial Media</Title> 
                <form onSubmit={searchUsersSubmit}>
                    <TextInput placeholder="search users..." ref={searchQueryRef}/>    
                </form>
            <Group style={{display:"flex","justifyContent":"center",height:"70px"}}>
            {username}
            {loggedIn ? <Button color="red" onClick={logOutHandler}>Logout</Button> : <Button color="green" onClick={logInHandler}>Login</Button>}
            </Group>
        </NavbarStyled>
        <Divider pb="xl"/>
        
        </>
    )
}

const NavbarStyled = styled.div`
    display:flex;
    align-items:center;
    justify-content:space-around;
    padding-left:10px;
    padding-right:10px;
    height:70px;


    & h1{
        cursor:pointer;
    }

    & form {
        display:flex;
        gap:10px;
    }
`