import React from 'react'
import SignupForm from '../../src/components/modules/signupFrom'
import { Box } from '@mui/material'

function Signup() {

    return (
        <Box
            sx={{
                my: 4,
                backgroundColor: 'secondary.light',
            }}
        >
            <SignupForm />
        </Box>
    )
}

// Signup.layout = "Fullpage"
export default Signup