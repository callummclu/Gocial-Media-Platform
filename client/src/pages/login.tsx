import { useRef  } from 'react';
import useAuth from '../hooks/useAuth';
import { LogInUser } from '../Types/auth';
import { Button, Card, PasswordInput, TextInput, Title, Text} from '@mantine/core'

function Login(props:any) {
  const usernameRef = useRef<HTMLInputElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)
  const { login, loggedIn } = useAuth()

  async function formSubmitLogin(e:any){
    e.preventDefault()
    let LoginParams: LogInUser = {
      username: usernameRef.current!.value,
      password: passwordRef.current!.value
    }
    await login(LoginParams)
  }

  return (
    <>
    <div style={{"display":"flex","alignItems":"center","justifyContent":"center","height":"calc(100vh - 150px)"}}>
      <Card p="xl" shadow="sm" radius="md" withBorder style={{"width":"370px"}}>
    
    {loggedIn ? window.location.href = window.location.origin : (
    <>
    
        <Title mt={"sm"} order={1}>Log In</Title>
        <Text mb={"md"} color={"darkgray"}>Enter your details below</Text>
        <form onSubmit={formSubmitLogin}>
          <TextInput label="Username" ref={usernameRef} type="name" required/>
          <PasswordInput label="Password" ref={passwordRef} type="password" required/>
          <Text mt="sm">Don't have an account? sign up <a href="/signup">here</a></Text>
          <Button mt="sm" type="submit">log in</Button>
        </form>

    </>
    
    )}
        </Card>
    </div>
    </>
  );
}

export default Login;
