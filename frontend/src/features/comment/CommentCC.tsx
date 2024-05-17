'use client';
import * as React from 'react';
import { Box } from '@mui/material';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import UserAvatar from '../../components/Avatar/UserAvatar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import FavoriteIcon from '@mui/icons-material/Favorite';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import Rating from '@mui/material/Rating';
import { useErrorHandler } from 'react-error-boundary';
import { Post, Comment } from '../../openapi';
import { isValidUrl } from '../../utils/isValidUrl';
import Link from '@mui/material/Link';
import { Menu, MenuItem } from '@mui/material';
import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import { User } from '../../openapi';
import axios from 'axios';
import { useFormState } from 'react-dom';
import { createComment } from './CreateCommentAction';
import { State } from './CreateCommentAction';

const initialState: State = {
  error: '',
  fieldErrors: {
    content: '',
    postID: '',
  },
};

export function CommentCC({
  post,
  comments,
  likedPostIDs,
  likedCommentIDs,
}: {
  post: Post;
  comments: Comment[];
  likedPostIDs: number[];
  likedCommentIDs: number[];
}) {
  const [state, formAction] = useFormState(createComment, initialState);
  const errorHandler = useErrorHandler();
  const [postAnchorEl, setPostAnchorEl] = React.useState<null | HTMLElement>(
    null,
  );
  const [commentAnchorEl, setCommentAnchorEl] =
    React.useState<null | HTMLElement>(null);
  const [selectedPostID, setSelectedPostID] = React.useState<number>(0);
  const [selectedCommentID, setSelectedCommentID] = React.useState<number>(0);
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

  const handlePostSettingClick = (
    event: React.MouseEvent<HTMLElement>,
    postID: number,
  ) => {
    event.stopPropagation();
    setPostAnchorEl(event.currentTarget);
    setSelectedPostID(postID);
  };

  const handlePostSettingClose = () => {
    setPostAnchorEl(null);
    setSelectedPostID(0);
  };

  const handleCommentSettingClick = (
    event: React.MouseEvent<HTMLElement>,
    commentID: number,
  ) => {
    setCommentAnchorEl(event.currentTarget);
    setSelectedCommentID(commentID);
  };

  const handleCommentSettingClose = () => {
    setCommentAnchorEl(null);
    setSelectedCommentID(0);
  };

  const DeletePost = async () => {
    if (selectedPostID > 0) {
      try {
        await axios.get(`/api/delete-post?selectedPostId=${selectedPostID}`);
      } catch (error: unknown) {
        errorHandler(error);
      }
    }
  };

  const DeleteComment = async () => {
    if (selectedCommentID > 0) {
      try {
        await axios.get(
          `/api/delete-comment?selectedCommentId=${selectedCommentID}`,
        );
      } catch (error: unknown) {
        errorHandler(error);
      }
    }
  };

  const handlePostLikeClick = async (postID: number) => {
    try {
      await axios.get(`/api/create-post-like?postID=${postID}`);
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handlePostUnLikeClick = async (postID: number) => {
    try {
      await axios.get(`/api/delete-post-like?postID=${postID}`);
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleCommentLikeClick = async (commentID: number) => {
    try {
      await axios.get(`/api/create-comment-like?commentID=${commentID}`);
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleCommentUnLikeClick = async (commentID: number) => {
    try {
      await axios.get(`/api/delete-comment-like?commentID=${commentID}`);
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  return (
    <>
      {post && (
        <Box sx={{ border: '1px solid #BDBDBD' }}>
          <Card
            sx={{
              width: '100%',
              minWidth: '600',
              backgroundColor: '#EFEBE5',
              boxShadow: 'none',
              cursor: 'pointer',
              '&:hover': {
                color: 'inherit',
                backgroundColor: '#EAE6E0',
              },
              '@media (max-width: 500px)': {
                width: '100%',
              },
            }}
            key={post.post_id}
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
                        onClick={(e) => handlePostSettingClick(e, post.post_id)}
                      >
                        <MoreVertIcon />
                      </IconButton>
                      <Menu
                        anchorEl={postAnchorEl}
                        open={Boolean(postAnchorEl)}
                        onClose={(e: React.MouseEvent) => {
                          e.stopPropagation();
                          handlePostSettingClose();
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
                  <IconButton
                    onClick={() => handlePostUnLikeClick(post.post_id)}
                  >
                    <FavoriteIcon sx={{ color: '#FF69B4' }} />
                  </IconButton>
                ) : (
                  <IconButton onClick={() => handlePostLikeClick(post.post_id)}>
                    <FavoriteBorderIcon />
                  </IconButton>
                )}

                {post.likes?.length}
              </Box>
              <Rating name="read-only" value={post.rating} readOnly />
            </CardActions>
          </Card>
        </Box>
      )}

      <Box
        sx={{
          padding: '1rem',
          borderRight: '1px solid #BDBDBD',
          borderLeft: '1px solid #BDBDBD',
          borderBottom: '1px solid #BDBDBD',
        }}
      >
        <form action={formAction}>
          <Box
            sx={{
              display: 'flex',
            }}
          >
            <TextField
              fullWidth
              id="content"
              label="Post your reply"
              name="content"
            />
            <input
              type="hidden"
              id="postID"
              name="postID"
              value={post.post_id}
            />
            <Button
              type="submit"
              sx={{
                backgroundColor: '#FF7E73',
                color: '#fff',
                '&:hover': {
                  backgroundColor: '#E56A67',
                },
                '&.Mui-disabled': {
                  backgroundColor: '#FFA49D',
                  color: '#fff',
                },
              }}
            >
              Reply
            </Button>
          </Box>
        </form>
        {state.fieldErrors?.content && (
          <span style={{ color: 'red' }}>{state.fieldErrors?.content}</span>
        )}
      </Box>
      <Box
        sx={{
          borderRight: '1px solid #BDBDBD',
          borderLeft: '1px solid #BDBDBD',
        }}
      >
        {comments.length > 0 ? (
          <>
            {comments.map((comment) => (
              <Box
                key={comment.comment_id}
                sx={{
                  display: 'flex',
                  flexDirection: 'column',
                  padding: '0.5rem',
                  borderBottom: '1px solid #BDBDBD',
                  cursor: 'pointer',
                  '&:hover': {
                    color: 'inherit',
                    backgroundColor: '#EAE6E0',
                  },
                }}
              >
                <Box sx={{ display: 'flex', alignItems: 'center' }}>
                  <UserAvatar
                    image={comment.user.profile_image}
                    userID={comment.user.user_id}
                  />
                  <Box sx={{ margin: '0.5rem' }}>
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <Typography variant="h6" color="black">
                        {comment.user.name}
                      </Typography>
                      <Typography
                        color="gray"
                        sx={{ marginLeft: '0.5rem', fontSize: '0.9rem' }}
                      >
                        {comment.created_at}
                      </Typography>
                    </Box>
                    <Typography
                      variant="body2"
                      color="black"
                      style={{ wordWrap: 'break-word' }}
                    >
                      {comment.content}
                    </Typography>
                  </Box>
                  {comment.user.user_id === user?.user_id && (
                    <Box
                      sx={{
                        marginLeft: 'auto',
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'start',
                      }}
                    >
                      <IconButton
                        onClick={(e) =>
                          handleCommentSettingClick(e, comment.comment_id)
                        }
                      >
                        <MoreVertIcon />
                      </IconButton>
                      <Menu
                        anchorEl={commentAnchorEl}
                        open={Boolean(commentAnchorEl)}
                        onClose={() => {
                          handleCommentSettingClose();
                        }}
                        sx={{
                          '& .MuiPaper-root': {
                            boxShadow: 'none',
                            border: 'none',
                          },
                        }}
                      >
                        <MenuItem
                          onClick={() => {
                            DeleteComment();
                          }}
                        >
                          コメントを削除
                        </MenuItem>
                      </Menu>
                    </Box>
                  )}
                </Box>
                <Box>
                  {likedCommentIDs.includes(comment.comment_id) ? (
                    <IconButton
                      onClick={() =>
                        handleCommentUnLikeClick(comment.comment_id)
                      }
                    >
                      <FavoriteIcon sx={{ color: '#FF69B4' }} />
                    </IconButton>
                  ) : (
                    <IconButton
                      onClick={() => handleCommentLikeClick(comment.comment_id)}
                    >
                      <FavoriteBorderIcon />
                    </IconButton>
                  )}

                  {comment.likes?.length}
                </Box>
              </Box>
            ))}
          </>
        ) : (
          <Box></Box>
        )}
      </Box>
    </>
  );
}
