import React from 'react'
import LoginForm from '../../src/components/modules/loginFrom'
import { Box } from '@mui/material'

function Login() {

    return (
        <Box
            sx={{
                my: 4,
                backgroundColor: 'secondary.light',
            }}
        >
            <LoginForm />
        </Box>
    )
}

// Signup.layout = "Fullpage"
export default Login