import { useEffect, useRef, useState } from 'react';
import { BrowserRouter, Navigate, Route, RouteProps, Routes } from 'react-router-dom';
import { Navbar } from './components/navbar';
import Home from './pages/Home';
import Login from './pages/login';
import SearchResults from './pages/SearchResults';
import UserProfile from './pages/UserProfile';
import Signup from './pages/signup';
import { NewPost } from './components/newPost';
import UserSettings from './pages/userSettings';
import {Error} from './components/error'
import useAuth from './hooks/useAuth';
import ScrollToTop from './helpers/scrollToTop';
import { Footer } from './components/footer';


function App() {

  const {loggedIn,reload} = useAuth()
  const [updatePosts, setUpdatePosts] = useState(false)

  useEffect(()=>{
    reload()
  },[updatePosts])

  return(
    <>
    <BrowserRouter>
      <Navbar/>
      {loggedIn && <NewPost updatePosts={[updatePosts, setUpdatePosts]}/>}
      <ScrollToTop/>
      <Routes>
        <Route path="*" element={<Error/>}/>
        <Route path="" element={<Home updatePosts={[updatePosts, setUpdatePosts]}/>}/>
        <Route path="login" element={<Login />}/>
        <Route path="signup" element={<Signup />}/>
        <Route path="results" element={<SearchResults/>}/>
        <Route path="users/:username" element={<UserProfile updatePosts={[updatePosts, setUpdatePosts]}/>}/>
        <Route path="users/:username/settings" element={<UserSettings />}/>
      </Routes>
      <Footer/>
    </BrowserRouter>
    </>
  )
}

export default App;
