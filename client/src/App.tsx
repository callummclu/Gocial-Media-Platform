import React, { useEffect, useRef, useState } from 'react';
import {checkAuth, logIn, logOut} from './helpers/authHelper'
import { LogInUser } from './Types/auth';
function App() {

  const [loggedIn, setLoggedIn] = useState(false)
  const usernameRef = useRef<HTMLInputElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)

  const checkLoggedIn = async () => {
    setLoggedIn(await checkAuth())
  }

  const logOutHandler = async () => {
    setLoggedIn(await logOut())
  }

  useEffect(()=>{
    checkLoggedIn()
  },[])

  async function formSubmitLogin(){
    let LoginParams: LogInUser = {
      username: usernameRef.current!.value,
      password: passwordRef.current!.value
    }
    setLoggedIn(await logIn(LoginParams))
  }

  return (
    <>
    {loggedIn ? <button onClick={logOutHandler}>Logout</button> : (
    <>
      <label>username</label>
      <input ref={usernameRef} type="name"/>
      <label>password</label>
      <input ref={passwordRef} type="password"/>
      <button onClick={formSubmitLogin}>log in</button>
    </>
    )}
    {loggedIn && <h1>You are logged In</h1>}
    </>
  );
}

export default App;
