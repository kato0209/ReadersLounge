import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import AccountCircle from '@mui/icons-material/AccountCircle';
import MenuItem from '@mui/material/MenuItem';
import Menu from '@mui/material/Menu';
import { FaBookOpen } from 'react-icons/fa';
import useLogout from '../../features/logout/logout';
import { Link } from 'react-router-dom';
import useMediaQuery from '@mui/material/useMediaQuery';
import { CreatePost } from '../../features/home/CreatePost';
import HomeIcon from '@mui/icons-material/Home';
import NotificationsIcon from '@mui/icons-material/Notifications';
import MailIcon from '@mui/icons-material/Mail';
import SearchIcon from '@mui/icons-material/Search';
import { useAuthUserContext } from '../../lib/auth/auth';
import { Avatar, Button } from '@mui/material';
import { isValidUrl } from '../../utils/isValidUrl';
import PersonSearchIcon from '@mui/icons-material/PersonSearch';

export default function AppHeader() {
  const [profileAnchorEl, setProfileAnchorEl] = React.useState<null | HTMLElement>(null);
  const [MenuAnchorEl, setMenuAnchorEl] = React.useState<null | HTMLElement>(null);
  const { user } = useAuthUserContext();

  const handleProfile = (event: React.MouseEvent<HTMLElement>) => {
    setProfileAnchorEl(event.currentTarget);
  };

  const handleProfileClose = () => {
    setProfileAnchorEl(null);
  };

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setMenuAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setMenuAnchorEl(null);
  };

  const handleLogout = useLogout();
  

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar 
        position="fixed" 
        sx={{ 
          backgroundColor: '#FF7E73', 
          boxShadow: 'none', 
          '& .MuiPaper-root': {
            padding: '0px!important',
          },
        }}>
        <Toolbar sx={{ alignItems: 'center', height: '3rem', minHeight: '28px !important' }}>
          <>
            <IconButton
              size="large"
              edge="start"
              color="inherit"
              aria-label="menu"
              onClick={handleMenu}
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={MenuAnchorEl}
              anchorOrigin={{
                vertical: 'top',
                horizontal: 'right',
              }}
              keepMounted
              transformOrigin={{
                vertical: 'top',
                horizontal: 'right',
              }}
              sx={{ 
                '& .MuiPaper-root': {
                  paddingLeft: '5px', paddingRight: '5px'
                },
              }}
              open={Boolean(MenuAnchorEl)}
              onClose={handleMenuClose}
            >
              <MenuItem component={Link} to="/" sx={{display: 'flex', '&:hover': { color: 'black'}}}>
                <HomeIcon sx={{marginRight:'0.5rem'}}/>
                Home
              </MenuItem>
              <MenuItem component={Link} to="/user-search" sx={{display: 'flex', '&:hover': { color: 'black'}}}>
                <PersonSearchIcon sx={{marginRight:'0.5rem'}}/>
                ユーザー検索
              </MenuItem>
              <MenuItem component={Link} to="/chat-room-list" sx={{display: 'flex', '&:hover': { color: 'black'}}}>
                <MailIcon sx={{marginRight:'0.5rem'}}/>
                Messages
              </MenuItem>
              <MenuItem component={Link} to="/search-book" sx={{display: 'flex', '&:hover': { color: 'black'}}}>
                <SearchIcon sx={{marginRight:'0.5rem'}} />
                本を探す
              </MenuItem>
              <CreatePost displayString='Post'/>
            </Menu>
          </>
          <Box 
            component={Link} 
            to="/" 
            sx={{ 
              display: 'flex', 
              alignItems: 'center', 
              color: 'inherit',
              '&:hover': {
                color: '#f0f0f0'
              },
              }}
          >
            <FaBookOpen/>
            <Typography variant="h6" component="div" sx={{ ml: 0.5 }}>
              ReadersLounge
            </Typography>
          </Box>
          <Box style={{ flexGrow: 1 }}></Box>
          <>
            <Button
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleProfile}
              color="inherit"
              sx={{borderRadius: '50%'}}
            >
              <Avatar 
                sx={{
                  border: "1px solid rgba(0, 0, 0, 0.2)",
                  boxShadow: "0px 2px 4px rgba(0, 0, 0, 0.1)",
                  '@media (max-width: 500px)': {
                    width: '30px',
                    height: '30px',
                }
                }} 
                src={isValidUrl(user.profile_image) ? user.profile_image : `data:image/png;base64,${user.profile_image}` } 
              />
            </Button>
            <Menu
              id="menu-appbar"
              anchorEl={profileAnchorEl}
              anchorOrigin={{
                vertical: 'top',
                horizontal: 'right',
              }}
              keepMounted
              transformOrigin={{
                vertical: 'top',
                horizontal: 'right',
              }}
              open={Boolean(profileAnchorEl)}
              onClose={handleProfileClose}
            >
              <MenuItem component={Link} to={`/user-profile/${user.user_id}`} sx={{display: 'flex', '&:hover': { color: 'black'}}}>プロフィール</MenuItem>
              <MenuItem onClick={handleLogout}>ログアウト</MenuItem>
            </Menu>
          </>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
