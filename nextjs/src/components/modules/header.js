import React from 'react'
import { AppBar, Typography, Box } from '@mui/material'
import { IconButton } from '@mui/material'
import { Badge } from '@mui/material'
import { Toolbar } from '@mui/material'
import Link from 'next/link';
import { useContext } from 'react'
import AuthContext from '../../api/auth/authContext'

import { appBarItems } from '../../router/routes'
import LogoutIcon from '@mui/icons-material/Logout'

export default function Header() {

    const { logout, user } = useContext(AuthContext)

    return (

        <AppBar position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }} >
            <Toolbar>
                <Typography variant="h6" noWrap component="div">
                    <Link href={'/'}>Gleam</Link>
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Box sx={{ display: { xs: 'none', md: 'flex' } }}>
                    {appBarItems.map((navItem, i) => (
                        <Link href={navItem.path} key={i}>
                            <IconButton size="large" color="inherit">
                                {navItem.icon}
                            </IconButton>
                        </Link>
                    ))}
                    {user ? <>
                        <div>
                            <IconButton onClick={() => logout()} size="large" color="inherit">
                                <LogoutIcon />
                            </IconButton>
                        </div>
                    </> : null}
                </Box>
            </Toolbar>
        </AppBar>

    )
}


