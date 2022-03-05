import axios from 'axios'
import { API_URL } from '../config/appConfig'

const api = axios.create({
    baseURL: `${API_URL}`,
    headers: {
        'Content-Type': 'application/json'
    }
})

// request interceptors





//refresh Token function

// response interceptors
