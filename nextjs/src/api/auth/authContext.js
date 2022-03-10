import { useState, useEffect, createContext } from 'react'
import { useRouter } from 'next/router'
import axios from 'axios'
import { API_URL } from '../../config/appConfig'


const AuthContext = createContext()

export const AuthProvider = ({ children }) => {

    const [user, setUser] = useState(null)
    const [accessToken, setAccessToken] = useState(null)
    const [error, setError] = useState(null)
    const [isLoading, setIsLoading] = useState(false)

    const router = useRouter()

    useEffect(() => checkedUserLoggedIn(), [])
    useEffect(() => getCSRF(), [])

    useEffect(() => {
        if (!user === null) {
            console.log("useEffect user : " + accessToken)
            axios.defaults.headers.common['Authorization'] = 'Bearer ' + accessToken
        } else {
            axios.defaults.headers.common['Authorization'] = {};
        }
    }, [user])


    const getCSRF = async () => {
        await axios.get(`${API_URL}/csrf`)
            .then((response) => {
                axios.defaults.headers.common['X-CSRF-TOKEN'] = response.data.csrf
            }, (error) => {
                console.log(error)
            })
    }

    //Sign Up user
    const signup = async ({ firstName, lastName, email, password }) => {

        await axios.post(`${API_URL}/signup`, {
            firstName: firstName,
            lastName: lastName,
            email: email,
            password: password
        })
            .then((response) => {
                console.log("after sign up:" + response.data.access_token)
                setAccessToken(response.data.access_token)
                setUser(response.data.userID)
                router.push('/')
            }, (error) => {

                setError(error)
                setError(null)
            })

    }


    //Login user
    const login = async ({ email, password }) => {

        await axios.post(`${API_URL}/login`, {
            email: email,
            password: password
        })
            .then((response) => {
                setAccessToken(response.data.access_token)
                setUser(response.data.userID)
                router.push('/')
            }, (error) => {

                setError(error)
                setError(null)
            })
    }

    //Logout
    const logout = async () => {

        // remove user
        // remove accessToken
        // remove axios header


        await axios.get(`${API_URL}/login`)
            .then((response) => {
                setAccessToken(null)
                setUser(null)
            }, (error) => {
                setError(error)
                setError(null)
            })






        axios.defaults.headers.common['Authorization'] = {};


        // const res = await fetch(`/api/logout`, {
        //     method: 'POST',
        // })

        // if (res.ok) {
        //     setUser(null)
        //     router.push('/')
        // }
    }

    // check user log in
    const checkedUserLoggedIn = async () => {


        // use refresh token to get access token
        await axios.get(`${API_URL}/refresh`)
            .then((response) => {
                setAccessToken(response.data.access_token)
                setUser(response.data.userID)
            }, (error) => {
                setError(error)
                setError(null)
            })

    }

    return (
        <AuthContext.Provider value={{
            signup,
            login,
            logout,
            checkedUserLoggedIn,
            isLoading,
            user,
            error
        }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext