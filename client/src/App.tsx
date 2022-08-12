import { useEffect, useRef, useState } from 'react';
import {checkAuth, logIn, logOut} from './helpers/authHelper'
import { LogInUser } from './Types/auth';
import { Button, Card, PasswordInput, TextInput, Title, Text} from '@mantine/core'
import { showNotification } from '@mantine/notifications';

function App() {

  const [loggedIn, setLoggedIn] = useState(false)
  const usernameRef = useRef<HTMLInputElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)

  const checkLoggedIn = async () => {
    setLoggedIn(await checkAuth())
  }

  const logOutHandler = async () => {
    setLoggedIn(await logOut())
    showNotification({
      title:"Success",
      message:"You've logged out successfully"
    })
  }

  useEffect(()=>{
    checkLoggedIn()
  },[])

  async function formSubmitLogin(e:any){
    e.preventDefault()
    let LoginParams: LogInUser = {
      username: usernameRef.current!.value,
      password: passwordRef.current!.value
    }
    let loggedIn = await logIn(LoginParams)
    if (loggedIn === false) {
      console.log("error incorrect details")
      showNotification({
        title:"Error",
        message:"Incorrect Details Provided",
        color:"red"
      })
    } else {
      showNotification({
        title:"Congrats",
        message:"You've logged in",
        color:"green"
      })
    }
    setLoggedIn(loggedIn)
  }

  return (
    <>
    <div style={{"display":"flex","alignItems":"center","justifyContent":"center","height":"100vh","background":"#f5f5f5"}}>
      <Card p="xl" shadow="sm" radius="md" withBorder style={{"width":"370px"}}>
    
    {loggedIn ? <Button color="red" style={{width:"100%"}} onClick={logOutHandler}>Logout</Button> : (
    <>
    
        <Title mt={"sm"} order={1}>Log In</Title>
        <Text mb={"md"} color={"darkgray"}>Enter your details below</Text>
        <form onSubmit={formSubmitLogin}>
          <TextInput label="Username" ref={usernameRef} type="name" required/>
          <PasswordInput label="Password" ref={passwordRef} type="password" required/>
          <Button mt="sm" type="submit">log in</Button>
        </form>

    </>
    
    )}
        </Card>
    </div>
    </>
  );
}

export default App;
