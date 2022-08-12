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

function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  
  const checkLoggedIn = async () => {
    setLoggedIn(await checkAuth())
  }

  useEffect(()=>{
    checkLoggedIn()
  },[])

  return(
    <BrowserRouter>
      <Navbar loggedIn={[loggedIn,setLoggedIn]}/>
      <Routes>
        <Route path="" element={<Home/>}/>
        <Route path="login" element={<Login loggedIn={[loggedIn,setLoggedIn]}/>}/>
        <Route path="results" element={<SearchResults/>}/>
      </Routes>
    </BrowserRouter>
  )
}

export default App;
