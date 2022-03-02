import { API_URL, COOKIE_SECURE } from "../../src/config/appConfig"
import cookie from 'cookie'

export default async (req, res) => {
    if (req.method === 'POST') {

        const { email, password } = req.body

        const apiRes = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'applicable/json'
            },
            body: JSON.stringify({
                email,
                password
            })
        })

        const resData = await apiRes.json()

        console.log(resData.jwt)

        if (apiRes.ok) {

            //set cookie
            res.setHeader(
                'Set-Cookie',
                cookie.serialize('token', String(resData.data.token), {
                    httpOnly: true,
                    secure: { COOKIE_SECURE },
                    maxAge: 60 * 60 * 24 * 7, // 1 week
                    sameSite: 'strict',
                    path: '/'
                })
            )

            res.status(200).json({ user: resData.user })
        } else {
            res.status(data.statusCode).json({ message: resData.message })
        }

    } else {
        res.setHeader('Allow', [POST])
        res.status(405).json({ message: `Method ${req.method} not allowed` })
    }
}