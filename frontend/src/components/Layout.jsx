import React from 'react';
import { styled, typography } from '@mui/system';
import { Typography, Drawer } from '@mui/material';
import { List, ListItem, ListItemIcon, ListItemText } from '@mui/material';



const drawerWidth = 240;

const Root = styled('div')({
    display: 'flex',
});

const PageStyle = styled('div')({
  backgroundColor: 'aliceblue',
  padding: 8,
});

function Layout({children}) {
    return (
        <Root>
            <Drawer
                sx={{
                width: drawerWidth,
                flexShrink: 0,
                '& .MuiDrawer-paper': {
                    width: drawerWidth,
                    boxSizing: 'border-box',
                },
                }}
                variant="permanent"
                anchor="left"
            >
                <div>
                    <Typography variant="h5">
                        Gleam
                    </Typography>
                </div>

                

            </Drawer>
            
            <PageStyle>
            {children}
            </PageStyle>
        </Root>
    )
}

export default Layout;
