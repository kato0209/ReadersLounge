import Drawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import HomeIcon from '@mui/icons-material/Home';
import MailIcon from '@mui/icons-material/Mail';
import { CreatePost } from '../../features/home/CreatePost';
import SearchIcon from '@mui/icons-material/Search';
import Link from 'next/link';
import PersonSearchIcon from '@mui/icons-material/PersonSearch';

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
        <Link href="/" passHref>
          <ListItem
            button
            component="a"
            sx={{
              borderRadius: '50px',
              color: 'black',
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
        </Link>
        <Link href="/user-search" passHref>
          <ListItem
            button
            component="a"
            sx={{
              borderRadius: '50px',
              color: 'black',
              '&:hover': {
                color: 'black',
              },
            }}
          >
            <ListItemIcon>
              <PersonSearchIcon />
            </ListItemIcon>
            <ListItemText primary="User Search" />
          </ListItem>
        </Link>
        <Link href="/chat-room-list" passHref>
          <ListItem
            button
            component="a"
            sx={{
              borderRadius: '50px',
              color: 'black',
              '&:hover': {
                color: 'black',
              },
            }}
          >
            <ListItemIcon>
              <MailIcon />
            </ListItemIcon>
            <ListItemText primary="Messages" />
          </ListItem>
        </Link>
        <Link href="/search-book" passHref>
          <ListItem
            button
            component="a"
            sx={{
              borderRadius: '50px',
              color: 'black',
              '&:hover': {
                color: 'black',
              },
            }}
          >
            <ListItemIcon>
              <SearchIcon />
            </ListItemIcon>
            <ListItemText primary="Book Search" />
          </ListItem>
          <CreatePost displayString="Post" />
        </Link>
      </List>
    </Drawer>
  );
}
