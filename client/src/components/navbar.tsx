import { Button, Text, Divider, Group, TextInput, Title, Avatar, Menu, Image } from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import { useRef } from 'react'
import { useSearchParams } from 'react-router-dom'
import styled from 'styled-components'
import { logOut } from '../helpers/authHelper'
import {FiSettings} from 'react-icons/fi'
import {BiLogOut} from 'react-icons/bi'
import { CgProfile } from 'react-icons/cg'

const NavIcon = (props:any) => {
    const [loggedIn,setLoggedIn] = props.loggedIn 
    const [username,setUsername] = props.username


    let userRedirect = `/users/${username}`

    const goToProfile = () => {
        window.location.href = window.location.origin + userRedirect
    }

    const logOutHandler = async () => {
        setLoggedIn(await logOut())
        setUsername("")
        showNotification({
          title:"Success",
          message:"You've logged out successfully"
        })
      }

    return (
        <>
            
            <Menu shadow="md" width={200}>
                <Menu.Target>
                    <Avatar style={{cursor:"pointer"}} radius="xl"/>
                </Menu.Target>

                <Menu.Dropdown>
                    <Menu.Label>{username}</Menu.Label>
                    <Menu.Item onClick={goToProfile} icon={<CgProfile size={14}/>}>My Profile</Menu.Item>
                    <Menu.Item icon={<FiSettings size={14} />}>Settings</Menu.Item>
                    <Menu.Divider />
                    <Menu.Label>Danger zone</Menu.Label>
                    <Menu.Item onClick={logOutHandler} icon={<BiLogOut size={14} />}>Logout</Menu.Item>,
                </Menu.Dropdown>
                </Menu>
        </>
    )
}

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

      const logInHandler = () => {
        window.location.href = window.location.origin + '/login'
      }

      

    return (
        <>
        <NavbarStyled>
            <div style={{width: 100, cursor:"pointer" }}>
                <Image src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1280px-Go_Logo_Blue.svg.png" onClick={()=>window.location.href = window.location.origin}/> 
            </div>
                
            <Group style={{display:"flex","justifyContent":"center",height:"70px"}}>
            <form onSubmit={searchUsersSubmit}>
                    <TextInput placeholder="search users..." ref={searchQueryRef}/>    
                </form>
            {loggedIn ? <NavIcon loggedIn={[loggedIn,setLoggedIn]} username={[username,setUsername]}/> : <Button color="green" onClick={logInHandler}>Login</Button>}
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