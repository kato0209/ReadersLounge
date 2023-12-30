import * as React from 'react';
import Drawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import HomeIcon from '@mui/icons-material/Home';
import NotificationsIcon from '@mui/icons-material/Notifications';
import MailIcon from '@mui/icons-material/Mail';
import { CreatePost }  from '../../features/home/CreatePost';
import SearchIcon from '@mui/icons-material/Search';
import { Link } from 'react-router-dom';

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
            borderRight: 'none',
        },
        }}
        anchor="left"
    >
      <List sx={{ width: '60%' }}>
        <ListItem 
          button 
          component={Link} 
          to="/"
          sx={{ 
            borderRadius: '50px',
            '&:hover': {
              color: 'black',
            },
          }}
        >
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
        <ListItem 
          button 
          component={Link} 
          sx={{ 
            borderRadius: '50px',
            '&:hover': {
              color: 'black',
            },
          }} 
          to="/chat-room-list"
        >
          <ListItemIcon>
            <MailIcon />
          </ListItemIcon>
          <ListItemText primary="Messages" />
        </ListItem>
        <ListItem 
          button 
          component={Link} 
          sx={{ 
            borderRadius: '50px',
            '&:hover': {
              color: 'black',
            },
          }} 
          to="/search-book"
        >
          <ListItemIcon>
            <SearchIcon/>
          </ListItemIcon>
          <ListItemText primary="本を探す" />
        </ListItem>
        <CreatePost displayString='Post'/>
      </List>
    </Drawer>
  );
}
