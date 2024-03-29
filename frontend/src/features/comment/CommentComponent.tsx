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
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import {
  Post,
  Comment,
  ReqCreateCommentBody,
  CreateCommentLikeReqBody,
} from '../../openapi';
import { isValidUrl } from '../../utils/isValidUrl';
import Link from '@mui/material/Link';
import { Menu, MenuItem } from '@mui/material';
import { useAuthUserContext } from '../../lib/auth/auth';
import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';
import { CreatePostLikeReqBody, PostLike, CommentLike } from '../../openapi';
import { useNavigate } from 'react-router-dom';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { useParams } from 'react-router-dom';

const CommentSchema = z.object({
  content: z.string().nonempty('投稿内容は必須です').max(255, {
    message: '投稿内容は255文字以内で入力してください',
  }),
});

type FormData = z.infer<typeof CommentSchema>;

export function CommentComponent() {
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(CommentSchema),
  });
  const errorHandler = useErrorHandler();
  const [postAnchorEl, setPostAnchorEl] = React.useState<null | HTMLElement>(
    null,
  );
  const [commentAnchorEl, setCommentAnchorEl] =
    React.useState<null | HTMLElement>(null);
  const [selectedPostID, setSelectedPostID] = React.useState<number>(0);
  const [selectedCommentID, setSelectedCommentID] = React.useState<number>(0);
  const { user } = useAuthUserContext();
  const [likedPostIDs, setLikedPostIDs] = React.useState<number[]>([]);
  const [likedCommentIDs, setLikedCommentIDs] = React.useState<number[]>([]);
  const navigation = useNavigate();
  const { id } = useParams<{ id: string }>();
  const idNumber = id ? parseInt(id, 10) : 0;
  const [post, setPost] = React.useState<Post>();
  const [comments, setComments] = React.useState<Comment[]>([]);

  React.useEffect(() => {
    fetchPost();
    fetchComments();
    fetchLikedPostIDs();
    fetchLikedCommentIDs();
  }, []);

  const fetchPost = async () => {
    try {
      const api = await apiInstance;
      const res = await api.getPostByPostID(idNumber);
      if (res.data) {
        const post: Post = {
          post_id: res.data.post_id,
          user: res.data.user,
          content: res.data.content,
          rating: res.data.rating,
          image: res.data.image,
          created_at: res.data.created_at,
          book: res.data.book,
          likes: res.data.likes,
        };
        setPost(post);
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const fetchComments = async () => {
    try {
      const api = await apiInstance;
      const res = await api.getCommentsByPostID(idNumber);
      if (res.data) {
        const comments: Comment[] = res.data.map((item) => ({
          comment_id: item.comment_id,
          user: item.user,
          content: item.content,
          likes: item.likes,
          created_at: item.created_at,
        }));
        setComments(comments);
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const fetchLikedPostIDs = async () => {
    try {
      const api = await apiInstance;
      const res = await api.getLikedPostList();
      if (res.data && Array.isArray(res.data)) {
        const fetchedLikedPostIDs: number[] = res.data.map(
          (item) => item.post_id,
        );
        setLikedPostIDs(fetchedLikedPostIDs);
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const fetchLikedCommentIDs = async () => {
    try {
      const api = await apiInstance;
      const res = await api.getLikedCommentList();
      if (res.data && Array.isArray(res.data)) {
        const fetchedLikedCommentIDs: number[] = res.data.map(
          (item) => item.comment_id,
        );
        setLikedCommentIDs(fetchedLikedCommentIDs);
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

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
        const api = await apiInstance;
        await api.deletePost(selectedPostID);
        navigation(-1);
      } catch (error: unknown) {
        errorHandler(error);
      }
    }
    setPostAnchorEl(null);
    setSelectedPostID(0);
  };

  const DeleteComment = async () => {
    try {
      const api = await apiInstance;
      const res = await api.deleteComment(selectedCommentID);
      if (res.status === 204) {
        setComments((currentComments) =>
          currentComments.filter(
            (comment) => comment.comment_id !== selectedCommentID,
          ),
        );
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
    setCommentAnchorEl(null);
    setSelectedCommentID(0);
  };

  const handlePostLikeClick = async (postID: number) => {
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
        setLikedPostIDs((currentLikedPostIDs) => [
          ...currentLikedPostIDs,
          postID,
        ]);
        if (post) {
          if (post.likes) {
            post.likes.push(newLike);
          } else {
            post.likes = [newLike];
          }
        }
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleCommentLikeClick = async (commentID: number) => {
    try {
      const req: CreateCommentLikeReqBody = {
        comment_id: commentID,
      };
      const api = await apiInstance;
      const res = await api.createCommentLike(req);
      if (res.status === 201 && res.data) {
        const newLike: CommentLike = {
          comment_like_id: res.data.comment_like_id,
          user_id: user.user_id,
        };
        setLikedCommentIDs((currentLikedCommentIDs) => [
          ...currentLikedCommentIDs,
          commentID,
        ]);
        setComments((currentComments) =>
          currentComments.map((comment) =>
            comment.comment_id === commentID
              ? {
                  ...comment,
                  likes: Array.isArray(comment.likes)
                    ? [...comment.likes, newLike]
                    : [newLike],
                }
              : comment,
          ),
        );
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handlePostUnLikeClick = async (postID: number) => {
    try {
      const api = await apiInstance;
      const res = await api.deletePostLike(postID);
      if (res.status === 204) {
        setLikedPostIDs((currentLikedPostIDs) =>
          currentLikedPostIDs.filter((id) => id !== postID),
        );
        if (post) {
          if (post.likes) {
            post.likes = post.likes.filter(
              (like) => like.user_id !== user.user_id,
            );
          }
        }
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleCommentUnLikeClick = async (commentID: number) => {
    try {
      const api = await apiInstance;
      const res = await api.deleteCommentLike(commentID);
      if (res.status === 204) {
        setLikedCommentIDs((currentLikedCommentIDs) =>
          currentLikedCommentIDs.filter((id) => id !== commentID),
        );
        setComments((currentComments) =>
          currentComments.map((comment) =>
            comment.comment_id === commentID
              ? {
                  ...comment,
                  likes: comment.likes?.filter(
                    (like) => like.user_id !== user.user_id,
                  ),
                }
              : comment,
          ),
        );
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const onSubmit = async (data: FormData) => {
    if (!post) return;
    try {
      const req: ReqCreateCommentBody = {
        content: data.content,
        post_id: post.post_id,
      };
      const api = await apiInstance;
      const res = await api.createComment(req);
      if (res.status == 201 && res.data) {
        const newComment: Comment = {
          comment_id: res.data.comment_id,
          user: res.data.user,
          content: data.content,
          created_at: res.data.created_at,
          likes: [],
        };
        setComments((currentComments) => [...currentComments, newComment]);
        setValue('content', '');
      }
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
                  {post.user.user_id === user.user_id && (
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
        <Box component="form" noValidate onSubmit={handleSubmit(onSubmit)}>
          <Box
            sx={{
              display: 'flex',
            }}
          >
            <TextField
              {...register('content')}
              fullWidth
              id="content"
              label="Post your reply"
              name="content"
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
        </Box>
        {errors.content && (
          <span style={{ color: 'red' }}>{errors.content.message}</span>
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
                  {comment.user.user_id === user.user_id && (
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
