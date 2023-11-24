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

export default function Home() {
    const errorHandler = useErrorHandler();

    const [posts, setPosts] = React.useState<Post[]>([]);

    React.useEffect(() => {
        const fetchPosts = async () => {
        
            try {
                const api = await apiInstance;
                const res = await api.getPosts();
                
                if (res.data && Array.isArray(res.data)) {
                    const fetchedPosts: Post[] = res.data.map(item => ({
                      post_id: item.post_id,
                      user: item.user,
                      content: item.content,
                      rating: item.rating,
                      image: item.image,
                      created_at: item.created_at,
                      book: item.book,
                    }));
                    setPosts(fetchedPosts);
                }
            } catch (error: unknown) {
                errorHandler(error);
            }
                
        };
    
        fetchPosts();
      }, []);

  return (
    <div style={{ display: 'flex', justifyContent: 'center', flexDirection: 'column', alignItems: 'center' }}>
        {posts.length > 0 ? (
        <>
            {posts.map(post => (
                <Card sx={{ width: '50%', minWidth: 600 }} key={post.post_id}>
                <CardHeader
                    avatar={
                    <Avatar src={post.user.profile_image} aria-label="recipe">
                        
                    </Avatar>
                    }
                    action={
                    <IconButton aria-label="settings">
                        <MoreVertIcon />
                    </IconButton>
                    }
                    title={post.user.name}
                    subheader={post.created_at}
                    sx={{
                        '& .MuiCardHeader-content': {
                            display: 'flex',
                            alignItems: 'center',
                        },
                        '& .MuiCardHeader-title': {
                            fontSize: '1.3rem',
                        },
                        '& .MuiCardHeader-subheader': {
                            marginLeft: '1em',
                        },
                    
                    }}
                />
                {post.image && (
                    <CardMedia
                        component="img"
                        height="400"
                        src={`data:image/png;base64,${post.image}`}
                    />
                )}
                <CardContent>
                    <Typography variant="body2" color="black" style={{ wordWrap: 'break-word' }}>
                        {post.content}
                    </Typography>
                </CardContent>
                <CardActions disableSpacing
                    sx={{
                        justifyContent: 'space-between',
                    }}
                >
                    <IconButton aria-label="add to favorites">
                    <FavoriteIcon />
                    </IconButton>
                    <Rating name="read-only" value={post.rating} readOnly />
                </CardActions>
                </Card>
            ))}
        </>
      ) : (
        <div>Loading...</div>
      )}
    </div>
  );
}