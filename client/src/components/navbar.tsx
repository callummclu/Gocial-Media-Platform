import { Button, Text, Divider, Group, TextInput, Title, Avatar, Menu, Image, Indicator } from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import { useRef } from 'react'
import { useSearchParams } from 'react-router-dom'
import styled from 'styled-components'
import {FiSettings} from 'react-icons/fi'
import {BiLogOut} from 'react-icons/bi'
import { CgProfile } from 'react-icons/cg'
import useAuth from '../hooks/useAuth'

const NavIcon = (props:any) => {

    const {loggedIn, user, logout, reload} = useAuth()


    const goToProfile = () => {
        let userRedirect = `/users/${user?.username}`

        window.location.href = window.location.origin + userRedirect
    }

    const goToSettings = () => {
        let settingsRedirect = `/users/${user?.username}/settings`
        window.location.href = window.location.origin + settingsRedirect
    }

    const logOutHandler = async () => {
        logout()
        showNotification({
          title:"Success",
          message:"You've logged out successfully"
        })
        reload()
      }

    const isNotifications = () => {
        return (user?.received_invitations || []).length > 0
    }
      
    return (
        <>
            
            <Menu shadow="md" width={200}>
                <Menu.Target>

                    {isNotifications() ? 
                    <Indicator color="red" size={12} withBorder>
                        <Avatar src={user?.display_image} style={{cursor:"pointer"}} radius="xl"/>
                    </Indicator>
                    :
                    <Avatar src={user?.display_image} style={{cursor:"pointer"}} radius="xl"/>
                        }

                </Menu.Target>

                <Menu.Dropdown>
                    <Menu.Label>{user?.username}</Menu.Label>
                    <Menu.Item onClick={goToProfile} icon={<CgProfile size={14}/>}>My Profile</Menu.Item>
                    <Menu.Item onClick={goToSettings} icon={<FiSettings size={14} />}>Settings</Menu.Item>
                    <Menu.Divider />
                    <Menu.Label>Danger zone</Menu.Label>
                    <Menu.Item onClick={logOutHandler} icon={<BiLogOut size={14} />}>Logout</Menu.Item>,
                </Menu.Dropdown>
                </Menu>
        </>
    )
}

export const Navbar = (props:any) => {

    const {loggedIn} = useAuth()

    const searchQueryRef = useRef<HTMLInputElement>(null)
    const [searchParams,setSearchParams] = useSearchParams()


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
            {loggedIn ? <NavIcon/> : <Button color="green" onClick={logInHandler}>Login</Button>}
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