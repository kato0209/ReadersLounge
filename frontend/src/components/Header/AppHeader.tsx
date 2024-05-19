'use client';
import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import MenuItem from '@mui/material/MenuItem';
import Menu from '@mui/material/Menu';
import { FaBookOpen } from 'react-icons/fa';
import { CreatePost } from '../../features/home/CreatePost';
import HomeIcon from '@mui/icons-material/Home';
import MailIcon from '@mui/icons-material/Mail';
import SearchIcon from '@mui/icons-material/Search';
import { Avatar, Button } from '@mui/material';
import { isValidUrl } from '../../utils/isValidUrl';
import PersonSearchIcon from '@mui/icons-material/PersonSearch';
import Link from 'next/link';
import { User } from '../../openapi';
import axios from 'axios';
import { useErrorHandler } from 'react-error-boundary';
import { useRouter } from 'next/navigation';

export default function AppHeader() {
  const router = useRouter();
  const errorHandler = useErrorHandler();
  const [profileAnchorEl, setProfileAnchorEl] =
    React.useState<null | HTMLElement>(null);
  const [MenuAnchorEl, setMenuAnchorEl] = React.useState<null | HTMLElement>(
    null,
  );
  const [user, setUser] = React.useState<User | null>(null);

  async function fetchLoginUser() {
    try {
      const res = await axios.get(`/api/fetch-login-user`);
      return res.data;
    } catch (error: unknown) {
      errorHandler(error);
    }
  }

  React.useEffect(() => {
    fetchLoginUser().then((res) => {
      setUser(res.data);
    });
  }, []);

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

  const handleLogout = async () => {
    try {
      await axios.post(`/api/logout`);
      router.push('/login');
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

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
        }}
      >
        <Toolbar
          sx={{
            alignItems: 'center',
            height: '3rem',
            minHeight: '28px !important',
          }}
        >
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
                  paddingLeft: '5px',
                  paddingRight: '5px',
                },
              }}
              open={Boolean(MenuAnchorEl)}
              onClose={handleMenuClose}
            >
              <Link href="/home" passHref>
                <MenuItem
                  sx={{
                    display: 'flex',
                    color: 'black',
                    '&:hover': { color: 'black' },
                  }}
                >
                  <HomeIcon sx={{ marginRight: '0.5rem' }} />
                  Home
                </MenuItem>
              </Link>
              <Link href="/user-search" passHref>
                <MenuItem
                  sx={{
                    display: 'flex',
                    color: 'black',
                    '&:hover': { color: 'black' },
                  }}
                >
                  <PersonSearchIcon sx={{ marginRight: '0.5rem' }} />
                  ユーザー検索
                </MenuItem>
              </Link>
              <Link href="/chat-room-list" passHref>
                <MenuItem
                  sx={{
                    display: 'flex',
                    color: 'black',
                    '&:hover': { color: 'black' },
                  }}
                >
                  <MailIcon sx={{ marginRight: '0.5rem' }} />
                  Messages
                </MenuItem>
              </Link>
              <Link href="/search-book" passHref>
                <MenuItem
                  sx={{
                    display: 'flex',
                    color: 'black',
                    '&:hover': { color: 'black' },
                  }}
                >
                  <SearchIcon sx={{ marginRight: '0.5rem' }} />
                  本を探す
                </MenuItem>
              </Link>
              <CreatePost displayString="Post" />
            </Menu>
          </>
          <Link href="/home" passHref>
            <Box
              sx={{
                display: 'flex',
                alignItems: 'center',
                color: 'white',
                '&:hover': {
                  color: '#f0f0f0',
                },
              }}
            >
              <FaBookOpen />
              <Typography variant="h6" component="div" sx={{ ml: 0.5 }}>
                ReadersLounge
              </Typography>
            </Box>
          </Link>
          <Box style={{ flexGrow: 1 }}></Box>
          <>
            <Button
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleProfile}
              color="inherit"
              sx={{ borderRadius: '50%' }}
            >
              <Avatar
                sx={{
                  border: '1px solid rgba(0, 0, 0, 0.2)',
                  boxShadow: '0px 2px 4px rgba(0, 0, 0, 0.1)',
                  '@media (max-width: 500px)': {
                    width: '30px',
                    height: '30px',
                  },
                }}
                src={
                  isValidUrl(user?.profile_image)
                    ? user?.profile_image
                    : `data:image/png;base64,${user?.profile_image}`
                }
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
              <Link href={`/user-profile/${user?.user_id}`} passHref>
                <MenuItem
                  component="a"
                  sx={{
                    display: 'flex',
                    color: 'black',
                    '&:hover': { color: 'black' },
                  }}
                >
                  プロフィール
                </MenuItem>
              </Link>
              <MenuItem onClick={handleLogout}>ログアウト</MenuItem>
            </Menu>
          </>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
