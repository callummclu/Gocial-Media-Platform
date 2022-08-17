import { LogInUser, SignUpUser, User, CheckUser } from "../Types/auth";
import React, {
    createContext,
    ReactNode,
    useContext,
    useEffect,
    useMemo,
    useState,
  } from "react";

import * as SessionsApi from  '../api/sessions'
import * as UsersApi from '../api/users'

interface AuthContextType {
    user?: User;
    loading: boolean;
    error?: any;
    login: (loginParams:LogInUser) => void;
    signUp: (signupParams:SignUpUser) => void;
    logout: () => void;
    reload: () => void;
    loggedIn: boolean
}

const AuthContext = createContext<AuthContextType>({} as AuthContextType)

export function AuthProvider({children}:{children:ReactNode}):JSX.Element {
    const [user, setUser] = useState<User>()
    const [loggedIn, setLoggedIn] = useState<boolean>(false)
    const [error,setError] = useState<any>()
    const [loading, setLoading] = useState<boolean>(false)
    const [loadingInitial, setLoadingInitial] = useState<boolean>(true)

    useEffect(()=>{
        if (error) setError(null)
    },[window.location.pathname])

    useEffect(() => {
        UsersApi.checkAuth()
            .then(async (res:any) => {
                let res_json = await res.json()
                setLoggedIn(res_json.isAuthenticated)
                UsersApi.getUserDetails(res_json.username)
                    .then(async (res:any) => {
                        let res_json = await res.json()
                        setUser(res_json.data)
                    })
            })
            .catch((_errpr)=>{})
            .finally(()=> setLoadingInitial(false))
    },[])

    function reload(){
        setLoading(true)
        UsersApi.checkAuth()
            .then(async (res:any) => {
                let res_json = await res.json()
                setLoggedIn(res_json.isAuthenticated)
                UsersApi.getUserDetails(res_json.username)
                    .then(async (res:any) => {
                        let res_json = await res.json()
                        setUser(res_json.data)
                    })
            })
            .catch((_error)=>{})
            .finally(()=>setLoading(false))
    }

    function login({username,password}:LogInUser){
        setLoading(true)
        SessionsApi.login({username,password})
            .then(async (res:any) => {
                let res_json = await res.json()
                if (!Object.hasOwn(res_json,'error')){
                    localStorage.setItem("gocial_auth_token",res_json.token)
                    UsersApi.checkAuth()
                        .then(async (res:any) => {
                            let res_json = await res.json()
                            setLoggedIn(res_json.isAuthenticated)
                            UsersApi.getUserDetails(res_json.username)
                                .then(async (res:any) => {
                                    let res_json = await res.json()
                                    setUser(res_json.data)
                                })
                        })
                        .catch((_error)=>{})
                } else {
                    setError(res_json.error)
                }
            }).catch((error)=>{
                setError(error)
            }).finally(()=>setLoading(false))
    }

    function signUp(signupParams:SignUpUser){
        setLoading(true)

        UsersApi.signup(signupParams)
        .then(async (res:any) => {
            let res_json = await res.json()
            if (Object.hasOwn(res_json,'error')){
                setError(res_json.error)
            } else {
                return login({username:signupParams.username,password:signupParams.password})
            }
        }).catch((error)=>{
            setError(error)
        }).finally(()=> setLoading(false))
    }

    function logout(){
        SessionsApi.logOut().then(()=>setUser(undefined))
    }

    const memoedValue = useMemo(
        () => ({
            user,
            loading,
            login,
            signUp,
            logout,
            loggedIn,
            reload
        }),
        [user,loading,error]
    )

    return (
        <AuthContext.Provider value={memoedValue}>
            {!loadingInitial && children}
        </AuthContext.Provider>
    )
}

export default function useAuth() {
    return useContext(AuthContext)
}