import { API_URL, COOKIE_SECURE } from "../../src/config/appConfig"
import cookie from 'cookie'

const signup = async (req, res) => {
    
    if (req.method === 'POST') {

        const { firstName, lastName, email, password } = req.body

        const apiRes = await fetch(`${API_URL}/signup`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                firstName,
                lastName,
                email,
                password
            }),
        })

        const resData = await apiRes.json()

        console.log(resData.data)

        // if (apiRes.ok) {
        //     // Set Cookie
        //     res.setHeader(
        //         'Set-Cookie',
        //         cookie.serialize('token', String(resData.data.token), {
        //             httpOnly: true,
        //             secure: { COOKIE_SECURE },
        //             maxAge: 60 * 60 * 24 * 7, // 1 week
        //             sameSite: 'strict',
        //             path: '/'
        //         })
        //     )

        //     res.status(200).json({ user: resData.data })
        // } else {
        //     res.status(500).json({ message: resData.message })
        // }
    } else {
        res.setHeader('Allow', ['POST'])
        res.status(405).json({ message: `Method ${req.method} not allowed` })
    }

}

export default signup