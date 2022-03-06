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


    const authHeader = async () => {
        const user = JSON.parse(localStorage.getItem('user'));
        if (!user === null && !accessToken === null) {
            return { Authorization: 'Bearer ' + accessToken };
        } else {
            return {};
        }
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
                setAccessToken(response.data.access_token)
                setUser(response.data.userID)
                router.push('/')
            }, (error) => {
                console.log(error)
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
                console.log(error)
                setError(error)
                setError(null)
            })
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