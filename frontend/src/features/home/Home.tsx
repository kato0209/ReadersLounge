import * as React from 'react';
import { styled } from '@mui/material/styles';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import Avatar from '@mui/material/Avatar';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import { red } from '@mui/material/colors';
import FavoriteIcon from '@mui/icons-material/Favorite';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import styles from './Home.css';
import Rating from '@mui/material/Rating';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { Post } from '../../openapi';
import Sidebar from './Sidebar';
import Box from '@mui/material/Box';
import PostList from './PostList';

export default function Home() {
    const [isMobile, setIsMobile] = React.useState(window.innerWidth <= 500);

    const handleResize = () => {
      setIsMobile(window.innerWidth <= 500);
    };

    React.useEffect(() => {
      window.addEventListener('resize', handleResize);
      return () => {
        window.removeEventListener('resize', handleResize);
      };
    }, []);

  return (
    <div style={{ display: 'flex'}}>
        {!isMobile && (
            <div style={{ flex: '0 0 30%', display: 'flex' }}>
                <Sidebar />
            </div>
        )}
        <div style={{ flex: 1, overflowX: 'hidden' }}>
            <PostList />
        </div>
    </div>
  );
}