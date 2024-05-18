'use client';
import * as React from 'react';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import FavoriteIcon from '@mui/icons-material/Favorite';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import Rating from '@mui/material/Rating';
import { useErrorHandler } from 'react-error-boundary';
import { Post } from '../../openapi';
import Box from '@mui/material/Box';
import { isValidUrl } from '../../utils/isValidUrl';
import Link from '@mui/material/Link';
import { Menu, MenuItem } from '@mui/material';
import UserAvatar from '../../components/Avatar/UserAvatar';
import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';
import { User } from '../../openapi';
import { useRouter } from 'next/navigation';
import axios from 'axios';
import { PostLike } from '../../openapi';

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
  },
};

type ListSectionProps = {
  propPosts: Post[];
  initialLikedPostIDs: number[];
};
export const ListSection: React.FC<ListSectionProps> = ({
  propPosts,
  initialLikedPostIDs,
}) => {
  const errorHandler = useErrorHandler();
  const [postAnchorEl, setPostAnchorEl] = React.useState<null | HTMLElement>(
    null,
  );
  const [selectedPostID, setSelectedPostID] = React.useState<number>(0);
  const [likedPostIDs, setLikedPostIDs] =
    React.useState<number[]>(initialLikedPostIDs);
  const [posts, setPosts] = React.useState<Post[]>(propPosts);

  const [user, setUser] = React.useState<User | null>(null);
  const router = useRouter();

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
    const handlePopState = () => {
      window.location.reload();
    };

    window.addEventListener('popstate', handlePopState);
  }, []);

  const handleSettingClick = (
    event: React.MouseEvent<HTMLElement>,
    postID: number,
  ) => {
    event.stopPropagation();
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
        await axios.get(`/api/delete-post?selectedPostId=${selectedPostID}`);
        setPosts((currentPosts) =>
          currentPosts.filter((post) => post.post_id !== selectedPostID),
        );
        handleSettingClose();
      } catch (error: unknown) {
        errorHandler(error);
      }
    }
  };

  const handleLikeClick = async (postID: number) => {
    try {
      const res = await axios.get(`/api/create-post-like?postID=${postID}`);
      if (res.status === 200 && res.data) {
        const newLike: PostLike = {
          post_like_id: res.data.postLike.post_like_id,
          user_id: user!.user_id,
        };
        setLikedPostIDs([...likedPostIDs, postID]);
        setPosts((currentPosts) =>
          currentPosts.map((post) =>
            post.post_id === postID
              ? {
                  ...post,
                  likes: Array.isArray(post.likes)
                    ? [...post.likes, newLike]
                    : [newLike],
                }
              : post,
          ),
        );
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleUnLikeClick = async (postID: number) => {
    try {
      const res = await axios.get(`/api/delete-post-like?postID=${postID}`);
      if (res.status === 200 && res.data) {
        setLikedPostIDs(likedPostIDs.filter((id) => id !== postID));
        setPosts((currentPosts) =>
          currentPosts.map((post) =>
            post.post_id === postID
              ? {
                  ...post,
                  likes: post.likes?.filter(
                    (like) => like.user_id !== user?.user_id,
                  ),
                }
              : post,
          ),
        );
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handlePostClick = async (postID: number) => {
    router.push(`/post/${postID}`);
  };

  return (
    <Box sx={PostListContainer}>
      {posts.length > 0 ? (
        <>
          {posts.map((post) => (
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
                },
              }}
              key={post.post_id}
              onClick={() => handlePostClick(post.post_id)}
            >
              <CardHeader
                avatar={
                  <UserAvatar
                    image={post.user.profile_image}
                    userID={post.user.user_id}
                  />
                }
                action={
                  <>
                    {post.user.user_id === user?.user_id && (
                      <>
                        <IconButton
                          onClick={(e) => handleSettingClick(e, post.post_id)}
                        >
                          <MoreVertIcon />
                        </IconButton>
                        <Menu
                          anchorEl={postAnchorEl}
                          open={Boolean(postAnchorEl)}
                          onClose={(e: React.MouseEvent) => {
                            e.stopPropagation();
                            handleSettingClose();
                          }}
                          sx={{
                            '& .MuiPaper-root': {
                              boxShadow: 'none',
                              border: 'none',
                            },
                          }}
                        >
                          <MenuItem
                            onClick={(event) => {
                              event.stopPropagation();
                              DeletePost();
                            }}
                          >
                            投稿を削除
                          </MenuItem>
                        </Menu>
                      </>
                    )}
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
                    },
                  },
                  '& .MuiCardHeader-subheader': {
                    marginLeft: '1em',
                  },
                }}
              />
              {post.image && (
                <Box sx={{ margin: '1rem' }}>
                  <CardMedia
                    component="img"
                    src={
                      isValidUrl(post.image)
                        ? post.image
                        : `data:image/png;base64,${post.image}`
                    }
                  />
                </Box>
              )}
              <CardContent
                onClick={(event) => event.stopPropagation()}
                sx={{ padding: '0px', margin: '1rem' }}
              >
                <Typography
                  variant="body2"
                  color="black"
                  style={{ wordWrap: 'break-word' }}
                >
                  {post.content}
                </Typography>
              </CardContent>
              <Box
                sx={{
                  display: 'flex',
                  border: '1px solid #BDBDBD',
                  justifyContent: 'space-between',
                  margin: '1rem',
                  borderRadius: '20px',
                  alignItems: 'center',
                }}
              >
                <CardContent
                  sx={{ flex: '1' }}
                  onClick={(event) => event.stopPropagation()}
                >
                  <Link href={post.book.item_url} underline="hover">
                    <Typography
                      component="div"
                      sx={{
                        fontSize: '1.3rem',
                        '@media (max-width: 500px)': { fontSize: '1.0rem' },
                      }}
                    >
                      {post.book.title}
                    </Typography>
                  </Link>
                  <Typography
                    variant="subtitle1"
                    color="text.secondary"
                    component="div"
                  >
                    著者：{post.book.author}
                  </Typography>
                </CardContent>
                <Box
                  sx={{
                    width: '30%',
                    margin: '1rem',
                    display: 'flex',
                    justifyContent: 'flex-end',
                    '@media (max-width: 500px)': {
                      margin: '0.2rem',
                      width: '35%',
                    },
                  }}
                >
                  <CardMedia
                    component="img"
                    sx={{
                      width: '60%',
                      '@media (max-width: 500px)': {
                        width: '100%',
                        margin: '0.4rem',
                      },
                    }}
                    image={post.book.image}
                  />
                </Box>
              </Box>
              <CardActions
                disableSpacing
                sx={{
                  justifyContent: 'space-between',
                }}
              >
                <Box onClick={(event) => event.stopPropagation()}>
                  {likedPostIDs.includes(post.post_id) ? (
                    <IconButton onClick={() => handleUnLikeClick(post.post_id)}>
                      <FavoriteIcon sx={{ color: '#FF69B4' }} />
                    </IconButton>
                  ) : (
                    <IconButton onClick={() => handleLikeClick(post.post_id)}>
                      <FavoriteBorderIcon />
                    </IconButton>
                  )}

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
};
