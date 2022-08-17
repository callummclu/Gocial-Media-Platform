import { useEffect, useRef, useState } from 'react';
import { Button, Card, PasswordInput, TextInput, Title, Text} from '@mantine/core'
import { showNotification } from '@mantine/notifications';
import { SignUpUser } from '../Types/auth';
import useAuth from '../hooks/useAuth';

function Signup(props:any) {
  const usernameRef = useRef<HTMLInputElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)
  const firstNameRef = useRef<HTMLInputElement>(null)
  const emailRef = useRef<HTMLInputElement>(null)
  const surnameRef = useRef<HTMLInputElement>(null)
  const confirmPasswordRef = useRef<HTMLInputElement>(null)
  const {signUp,loggedIn} = useAuth()

  async function formSubmitSignup(e:any){
    e.preventDefault()
    let SignupParmas: SignUpUser = {
      username: usernameRef.current!.value,
      password: passwordRef.current!.value,
      name: firstNameRef.current!.value,
      surname: surnameRef.current!.value,
      email: emailRef.current!.value
    }
    await signUp(SignupParmas)
  }

  return (
    <>
    <div style={{"display":"flex","alignItems":"center","justifyContent":"center","height":"calc(100vh - 150px)"}}>
      <Card p="xl" shadow="sm" radius="md" withBorder style={{"width":"370px"}}>
    
    {loggedIn ? window.location.href = window.location.origin : (
    <>
    
        <Title mt={"sm"} order={1}>Signup</Title>
        <Text mb={"md"} color={"darkgray"}>Enter your details below</Text>
        <form onSubmit={formSubmitSignup}>

        <TextInput label="Username" ref={usernameRef} type="name" required/>
        <TextInput label="Email" ref={emailRef} type="name" required/>
          <TextInput label="First Name" ref={firstNameRef} type="name" required/>
          <TextInput label="Surname" ref={surnameRef} type="name" required/>
          <PasswordInput label="Password" ref={passwordRef} type="password" required/>
          <PasswordInput label="Confirm Password" ref={confirmPasswordRef} type="password" required/>
          <Button mt="sm" type="submit">Sign up</Button>
        </form>

    </>
    
    )}
        </Card>
    </div>
    </>
  );
}

export default Signup;
