import * as React from 'react';
import Drawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import HomeIcon from '@mui/icons-material/Home';
import NotificationsIcon from '@mui/icons-material/Notifications';
import MailIcon from '@mui/icons-material/Mail';

export default function Sidebar() {

  return (
    <Drawer
        variant="permanent"
        sx={{
        backgroundColor: '#EFEBE5',
        '& .MuiDrawer-paper': {
            backgroundColor: '#EFEBE5',
            marginTop: '3rem',
            width: '30%',
            alignItems: 'center',
        },
        }}
        anchor="left"
    >
      <List sx={{ width: '60%' }}>
        <ListItem button sx={{ borderRadius: '50px' }}>
          <ListItemIcon>
            <HomeIcon />
          </ListItemIcon>
          <ListItemText primary="Home" />
        </ListItem>
        <ListItem button sx={{ borderRadius: '50px' }}>
          <ListItemIcon>
            <NotificationsIcon />
          </ListItemIcon>
          <ListItemText primary="Notifications" />
        </ListItem>
        <ListItem button sx={{ borderRadius: '50px' }}>
          <ListItemIcon>
            <MailIcon />
          </ListItemIcon>
          <ListItemText primary="Messages" />
        </ListItem>
        <ListItem 
          button 
          sx={{ 
            borderRadius: '50px', 
            backgroundColor: '#FF7E73',
            color: '#fff',
            marginTop: '0.8rem',
            '&:hover': {
              backgroundColor: '#E56A67',
            },
          }}>
          <ListItemText primary="Post" sx={{ textAlign: 'center' }} />
        </ListItem>
      </List>
    </Drawer>
  );
}
