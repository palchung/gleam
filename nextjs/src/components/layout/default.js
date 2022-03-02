import React from 'react'
import Box from '@mui/material/Box'
import { Container } from '@mui/material'
import Header from '../modules/header'

const Default = ({ children }) => {
    return (
        <Box sx={{ display: 'flex' }}>
            <Header />
            <Container>
                <Box sx={{
                    my: 10,
                    backgroundColor: '#92a8d1',
                    width: '100%',
                    height: '100%',
                }}>
                    {children}
                </Box>
            </Container>
        </Box>
    )
}

export default Default

