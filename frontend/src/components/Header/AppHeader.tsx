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


export default function AppHeader() {
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogout = useLogout();
  

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="fixed" sx={{ backgroundColor: '#FF7E73' }}>
        <Toolbar sx={{ alignItems: 'center', height: '3rem', minHeight: '28px !important' }}>
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
          >
            <MenuIcon />
          </IconButton>
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
          <div style={{ flexGrow: 1 }}></div>
          <>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleMenu}
              color="inherit"
            >
              <AccountCircle />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorEl}
              anchorOrigin={{
                vertical: 'top',
                horizontal: 'right',
              }}
              keepMounted
              transformOrigin={{
                vertical: 'top',
                horizontal: 'right',
              }}
              open={Boolean(anchorEl)}
              onClose={handleClose}
            >
              <MenuItem onClick={handleClose}>Profile</MenuItem>
              <MenuItem onClick={handleLogout}>ログアウト</MenuItem>
            </Menu>
          </>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
