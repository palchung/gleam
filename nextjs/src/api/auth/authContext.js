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

    // useEffect(() => checkedUserLoggedIn(), [])

    //Sign Up user
    const signup = async ({ firstName, lastName, email, password }) => {

        await axios.post(`${API_URL}/signup`, {
            firstName: firstName,
            lastName: lastName,
            email: email,
            password: password
        })
            .then((response) => {

                console.log('access token: ' + response.data.access_token)
                setAccessToken(response.data.access_token)
                console.log(response.data)

                // setUser(null)
                // router.push('/')
            }, (error) => {
                console.log(error)
                setError(error)
                setError(null)
            })

    }


    // const signup = async ({ firstName, lastName, email, password }) => {

    //     const fetcher = (...args) => fetch(...args).then(res => res.json())

    //     const res = await fetch(`/api/signup`, {
    //         method: 'POST',
    //         headers: {
    //             'Content-Type': 'application/json',
    //         },
    //         body: JSON.stringify({
    //             firstName,
    //             lastName,
    //             email,
    //             password
    //         })
    //     })

    //     const resData = await res.json()

    //     if (res.ok) {
    //         setUser(resData.user)
    //         router.push('/')
    //     } else {
    //         setError(resData.message)
    //         setError(null)
    //     }
    // }

    //Login user
    const login = async ({ email, password }) => {

        const res = await fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email,
                password
            })
        })
        const data = await res.json()

        if (res.ok) {
            setUser(data.user)
            router.push('/')
        } else {
            setError(data.message)
            setError(null)
        }
    }

    //Logout
    const logout = async () => {
        const res = await fetch(`/api/logout`, {
            method: 'POST',
        })

        if (res.ok) {
            setUser(null)
            router.push('/')
        }
    }

    // check user log in
    const checkedUserLoggedIn = async () => {
        const res = await fetch('/api/user')
        const data = await res.json()

        if (res.ok) {
            setUser(data.user)
        } else {
            setUser(null)
        }
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