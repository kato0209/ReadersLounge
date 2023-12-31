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
import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';
import { CreatePostLikeReqBody, PostLike } from '../../openapi';

const PostListContainer = {
    display: 'flex', 
    justifyContent: 'center', 
    flexDirection: 'column', 
    alignItems: 'start',
    flex: 1,
    borderTop: '1px solid #BDBDBD',
    borderRight: '1px solid #BDBDBD',
    borderLeft: '1px solid #BDBDBD',
    '@media (max-width: 500px)': {
        alignItems: 'center',
    }
};

type PostListProps = {
    propPosts: Post[];
};
export const PostList: React.FC<PostListProps> = ({ propPosts }) => {
    const errorHandler = useErrorHandler();
    const [postAnchorEl, setPostAnchorEl] = React.useState<null | HTMLElement>(null);
    const [selectedPostID, setSelectedPostID] = React.useState<number>(0);
    const { user } = useAuthUserContext();
    const [likedPostIDs, setLikedPostIDs] = React.useState<number[]>([]);
    const [posts, setPosts] = React.useState<Post[]>([]);

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
        setPostAnchorEl(null);
        setSelectedPostID(0);
    }

    const fetchLikedPostIDs = async () => {
        try {
            const api = await apiInstance;
            const res = await api.getLikedPostList();
            if (res.data && Array.isArray(res.data)) {
                const fetchedLikedPostIDs: number[] = res.data.map(item => item.post_id);
                setLikedPostIDs(fetchedLikedPostIDs);
            }
        } catch (error: unknown) {
            errorHandler(error);
        }
    }
    
    React.useEffect(() => {
        setPosts(propPosts);
        fetchLikedPostIDs();
    }, [propPosts]);

    const handleLikeClick = async (postID: number) => {
        try {
            const req: CreatePostLikeReqBody = {
                post_id: postID,
            };
            const api = await apiInstance;
            const res = await api.createPostLike(req);
            if (res.status === 201 && res.data) {
                const newLike: PostLike = {
                    post_like_id: res.data.post_like_id,
                    user_id: user.user_id,
                };
                setLikedPostIDs(currentLikedPostIDs => [...currentLikedPostIDs, postID]);
                setPosts(currentPosts =>
                    currentPosts.map(post =>
                        post.post_id === postID
                            ? {
                                ...post,
                                likes: Array.isArray(post.likes) ? [...post.likes, newLike] : [newLike]
                              }
                            : post
                    )
                );
            }
        } catch (error: unknown) {
            errorHandler(error);
        }
    };

    const handleUnLikeClick = async (postID: number) => {
        try {
            const api = await apiInstance;
            const res = await api.deletePostLike(postID);
            if (res.status === 204) {
                setLikedPostIDs(currentLikedPostIDs => 
                    currentLikedPostIDs.filter(id => id !== postID)
                );
                setPosts(currentPosts => 
                    currentPosts.map(post => 
                        post.post_id === postID ? { ...post, likes: post.likes?.filter(like => like.user_id !== user.user_id) } : post
                    )
                );
            }
        } catch (error: unknown) {
            errorHandler(error);
        }
    }

  return (
        <Box sx={PostListContainer}>
            {posts.length > 0 ? (
            <>
                {posts.map(post => (
                    <Card 
                        sx={{ 
                            width: '100%', 
                            minWidth: '600',
                            backgroundColor: '#EFEBE5', 
                            boxShadow: 'none',  
                            cursor: 'pointer',
                            borderBottom: '1px solid #BDBDBD',
                            '&:hover': {
                                color: 'inherit',
                                backgroundColor: '#EAE6E0',
                            },
                            '@media (max-width: 500px)': {
                                width: '100%',
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
                        <Box>
                            {likedPostIDs.includes(post.post_id) 
                            ? (
                                <IconButton onClick={() => handleUnLikeClick(post.post_id)}>
                                    <FavoriteIcon sx={{color: "#FF69B4"}} />
                                </IconButton>) : (
                                <IconButton onClick={() => handleLikeClick(post.post_id)}>
                                    <FavoriteBorderIcon />
                                </IconButton>
                                )
                            }
                            
                            {post.likes?.length}
                        </Box>
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