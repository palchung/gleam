import React from 'react'
import { appBarItems } from '../../router/routes'
import Link from '../../../src/Link'
import Box from '@mui/material/Box'
import { Container, Drawer, Toolbar, Typography } from '@mui/material'
import { List, ListItem, ListItemIcon, ListItemText } from '@mui/material'
import Header from '../modules/header'

const drawerWidth = 240

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

