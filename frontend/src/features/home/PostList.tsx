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
import Sidebar from '../../components/Sidebar/Sidebar';
import Box from '@mui/material/Box';
import { isValidUrl } from '../../utils/isValidUrl';
import Link from '@mui/material/Link';
import { Menu, MenuItem } from '@mui/material';
import UserAvatar from '../../components/Avatar/UserAvatar';
import { useAuthUserContext } from '../../lib/auth/auth';

const PostListContainer = {
    display: 'flex', 
    justifyContent: 'center', 
    flexDirection: 'column', 
    alignItems: 'start',
    flex: 1,
    '@media (max-width: 500px)': {
        alignItems: 'center',
    }
};

export default function PostList() {
    const errorHandler = useErrorHandler();
    const [postAnchorEl, setPostAnchorEl] = React.useState<null | HTMLElement>(null);
    const [selectedPostID, setSelectedPostID] = React.useState<number>(0);
    const { user } = useAuthUserContext();

    const handleSettingClick = (event: React.MouseEvent<HTMLElement>, postID: number) => {
        setPostAnchorEl(event.currentTarget);
        setSelectedPostID(postID);
    };

    const handleSettingClose = () => {
        setPostAnchorEl(null);
        setSelectedPostID(0);
    };

    const DeletePost = async () => {
        if (selectedPostID > 0) {
            try {
                const api = await apiInstance;
                const res = await api.deletePost(selectedPostID);
                setPosts(currentPosts => currentPosts.filter(post => post.post_id !== selectedPostID));
            } catch (error: unknown) {
                errorHandler(error);
            }
        }
    }

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
        <Box sx={PostListContainer}>
            {posts.length > 0 ? (
            <>
                {posts.map(post => (
                    <Card 
                        sx={{ 
                            width: '60%', 
                            minWidth: '600',
                            backgroundColor: '#EFEBE5', 
                            boxShadow: 'none',  
                            border: '1px solid #BDBDBD',
                            cursor: 'pointer',
                            '&:hover': {
                                color: 'inherit',
                                backgroundColor: '#EAE6E0',
                            },
                            '@media (max-width: 500px)': {
                                width: '80%',
                            }
                        }} 
                        key={post.post_id}>
                    <CardHeader
                        avatar={
                            <UserAvatar image={post.user.profile_image} userID={post.user.user_id}/>
                        }
                        action={
                        <>
                            {post.user.user_id === user.user_id && 
                                <>
                                    <IconButton onClick={(e) => handleSettingClick(e, post.post_id)}>
                                        <MoreVertIcon />
                                    </IconButton>
                                    <Menu
                                        anchorEl={postAnchorEl}
                                        open={Boolean(postAnchorEl)}
                                        onClose={handleSettingClose}
                                        sx={{
                                            '& .MuiPaper-root': {
                                                boxShadow: 'none', 
                                                border: 'none'
                                            }
                                        }}
                                    >
                                        <MenuItem onClick={() => {
                                            DeletePost();
                                        }}
                                        >
                                            投稿を削除
                                        </MenuItem>
                                    </Menu>
                                </>
                            }
                        </>
                        }
                        title={post.user.name}
                        subheader={post.created_at}
                        sx={{
                            '& .MuiCardHeader-content': {
                                display: 'flex',
                                alignItems: 'center',
                            },
                            '& .MuiCardHeader-title': {
                                fontSize: '1.2rem',
                                '@media (max-width: 500px)': {
                                    fontSize: '1rem',
                                }
                            },
                            '& .MuiCardHeader-subheader': {
                                marginLeft: '1em',
                            },
                        
                        }}
                    />
                    {post.image && (
                        <Box sx={{margin: '1rem'}}>
                            <CardMedia
                                component="img"
                                src={isValidUrl(post.image) ? post.image : `data:image/png;base64,${post.image}` }
                            />
                        </Box>
                    )}
                    <CardContent>
                        <Typography variant="body2" color="black" style={{ wordWrap: 'break-word' }}>
                            {post.content}
                        </Typography>
                    </CardContent>
                    <Box sx={{ 
                            display: 'flex', 
                            border: '1px solid #BDBDBD', 
                            justifyContent: 'space-between',
                            margin: '1rem',
                            borderRadius: '20px',
                            alignItems: 'center',
                        }}
                    >
                        <CardContent sx={{ flex: '1' }}>
                            <Link href={post.book.item_url} underline="hover">
                                <Typography component="div" sx={{ fontSize: '1.3rem', '@media (max-width: 500px)':{fontSize: '1.0rem'}}}>
                                    {post.book.title}
                                </Typography>
                            </Link>
                            <Typography variant="subtitle1" color="text.secondary" component="div">
                                著者：{post.book.author}
                            </Typography>
                        </CardContent>
                        <Box sx={{ 
                            width: '30%', 
                            margin: '1rem', 
                            display: 'flex', 
                            justifyContent: 'flex-end',
                            '@media (max-width: 500px)': {
                            margin: '0.2rem', 
                            width: '35%',
                            }
                        }}
                        >
                        <CardMedia
                            component="img"
                            sx={{ 
                            width: '60%',
                            '@media (max-width: 500px)':{
                                width: '100%',
                                margin: '0.4rem',
                            }
                            }}
                            image={post.book.image}
                        />
                        </Box>
                    </Box>
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
            <></>
        )}
        </Box>
  );
}