import { useEffect, useRef, useState } from 'react';
import {checkAuth, logIn, logOut} from './helpers/authHelper'
import { LogInUser } from './Types/auth';
import { Button, Card, PasswordInput, TextInput, Title, Text} from '@mantine/core'
import { showNotification } from '@mantine/notifications';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Navbar } from './components/navbar';
import Home from './pages/Home';
import Login from './pages/login';
import SearchResults from './pages/SearchResults';
import UserProfile from './pages/UserProfile';
import Signup from './pages/signup';
import { NewPost } from './components/newPost';

function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  const [username,setUsername] = useState("")

  const [updatePosts, setUpdatePosts] = useState(false)
  const [loggedInUserData, setLoggedInUserData] = useState<any>(null)

  const checkLoggedIn = async () => {
    let isAuth = await checkAuth()
    setLoggedIn(isAuth.isAuthenticated)
    setUsername(isAuth.username)
    fetch(`${process.env.REACT_APP_BACKEND_URI}/user/${isAuth.username}`)
            .then(async (res:any) => {
                let res_json = await res.json()
                setLoggedInUserData(res_json.data)
            })
  }

  useEffect(()=>{
    checkLoggedIn()
  },[updatePosts])

  return(
    <BrowserRouter>
      <Navbar loggedIn={[loggedIn,setLoggedIn]} username={[username,setUsername]}/>
      {loggedIn && <NewPost updatePosts={[updatePosts, setUpdatePosts]}/>}
      <Routes>
        <Route path="" element={<Home updatePosts={[updatePosts, setUpdatePosts]} username={[username,setUsername]}/>}/>
        <Route path="login" element={<Login loggedIn={[loggedIn,setLoggedIn]}/>}/>
        <Route path="signup" element={<Signup loggedIn={[loggedIn,setLoggedIn]}/>}/>
        <Route path="results" element={<SearchResults/>}/>
        <Route path="users/:username" element={<UserProfile loggedInUserData={[loggedInUserData, setLoggedInUserData]} loggedIn={[loggedIn,setLoggedIn]} updatePosts={[updatePosts, setUpdatePosts]}/>}/>
      </Routes>
    </BrowserRouter>
  )
}

export default App;
