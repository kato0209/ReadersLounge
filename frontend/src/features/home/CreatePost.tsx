'use client';
import * as React from 'react';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import Button from '@mui/material/Button';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';
import Avatar from '@mui/material/Avatar';
import { isValidUrl } from '../../utils/isValidUrl';
import Box from '@mui/material/Box';
import PostTextarea from './PostTextarea';
import Rating from '@mui/material/Rating';
import ImageIcon from '@mui/icons-material/Image';
import { useErrorHandler } from 'react-error-boundary';
import { Book } from '../../openapi';
import ImportContactsIcon from '@mui/icons-material/ImportContacts';
import { BookSearchDialog } from './BookSearchDialog';
import { User } from '../../openapi';
import axios from 'axios';
import { post } from './SubmitPost';
import { State } from './SubmitPost';
import { useFormState } from 'react-dom';
import { PostSchema } from '../../types/PostSchema';
import { z } from 'zod';

export const initialState: State = {
  error: '',
  fieldErrors: {
    content: '',
    rating: '',
    ISBNcode: '',
    postImage: '',
  },
};

type FormData = z.infer<typeof PostSchema>;
type CreatePostProps = {
  displayString: string;
  book?: Book;
  formData?: FormData;
};

export const CreatePost: React.FC<CreatePostProps> = ({
  displayString,
  book,
  formData,
}) => {
  const [state, formAction] = useFormState(post, initialState);
  const [openCreatePostDialog, setOpenCreatePostDialog] = React.useState(false);
  const [imagePreview, setImagePreview] = React.useState<string | null>(null);
  const [content, setContent] = React.useState<string>(formData?.content || '');
  const [rating, setRating] = React.useState<number>(formData?.rating || 0);
  const [ISBNcode] = React.useState<string>(formData?.ISBNcode || '');
  const [postImage, setPostImage] = React.useState<File | undefined>(
    formData?.postImage || undefined,
  );
  const [user, setUser] = React.useState<User | null>(null);
  const errorHandler = useErrorHandler();

  async function fetchLoginUser() {
    try {
      const res = await axios.get(`/api/fetch-login-user`);
      return res.data;
    } catch (error: unknown) {
      errorHandler(error);
    }
  }

  React.useEffect(() => {
    fetchLoginUser().then((data) => {
      setUser(data);
    });
  }, []);

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];
      setPostImage(file);
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleImageRemove = () => {
    setPostImage(undefined);
    setImagePreview(null);
  };

  const handleRatingChange = (newValue: number) => {
    setRating(newValue);
  };

  const handlePostTextChange = (
    event: React.ChangeEvent<HTMLTextAreaElement>,
  ) => {
    setContent(event.target.value);
  };

  const handleOpen = () => {
    setOpenCreatePostDialog(true);
  };

  const handleClose = () => {
    setOpenCreatePostDialog(false);
  };

  return (
    <React.Fragment>
      <ListItem
        button
        sx={{
          borderRadius: '100px',
          backgroundColor: '#FF7E73',
          color: '#fff',
          marginTop: '0.8rem',
          '&:hover': {
            backgroundColor: '#E56A67',
          },
        }}
        onClick={handleOpen}
      >
        <ListItemText
          primary={displayString}
          sx={{ textAlign: 'center', paddingHorizontal: '0.5rem' }}
        />
      </ListItem>

      <Dialog
        open={openCreatePostDialog}
        onClose={handleClose}
        fullWidth
        sx={{
          '& .MuiDialog-paper': {
            backgroundColor: '#EFEBE5',
          },
        }}
      >
        <IconButton
          aria-label="close"
          onClick={handleClose}
          sx={{
            position: 'absolute',
            left: 8,
            top: 8,
            color: (theme) => theme.palette.grey[500],
          }}
        >
          <CloseIcon />
        </IconButton>
        <form action={formAction}>
          <DialogContent
            sx={{
              flex: 'auto',
              marginTop: '2rem',
            }}
          >
            {book && (
              <Box
                sx={{
                  marginBottom: '1rem',
                  justifyContent: 'center',
                  display: 'flex',
                }}
              >
                <ImportContactsIcon />
                {book.title}
              </Box>
            )}
            <Box sx={{ display: 'flex', alignItems: 'start', gap: 2 }}>
              <Avatar
                src={
                  isValidUrl(user?.profile_image)
                    ? user?.profile_image
                    : `data:image/png;base64,${user?.profile_image}`
                }
              ></Avatar>
              <PostTextarea onChange={handlePostTextChange} value={content} />
            </Box>
            {imagePreview && (
              <Box
                sx={{
                  display: 'flex',
                  flexDirection: 'column',
                  justifyContent: 'center',
                }}
              >
                <Box
                  sx={{
                    display: 'flex',
                    justifyContent: 'center',
                    marginTop: '1rem',
                    position: 'relative',
                  }}
                >
                  <img
                    src={imagePreview}
                    style={{ width: '50%', height: 'auto' }}
                  />

                  <IconButton
                    aria-label="close"
                    onClick={handleImageRemove}
                    sx={{
                      position: 'absolute',
                      right: 0,
                      top: 0,
                      color: (theme) => theme.palette.grey[500],
                    }}
                  >
                    <CloseIcon />
                  </IconButton>
                </Box>
                {state.fieldErrors?.postImage && (
                  <span style={{ color: 'red', textAlign: 'center' }}>
                    {state.fieldErrors.postImage}
                  </span>
                )}
              </Box>
            )}
          </DialogContent>
          <DialogActions sx={{ justifyContent: 'start' }}>
            <Box>
              <Rating
                name="rating"
                value={Number(rating)}
                onChange={(event, newValue) => {
                  handleRatingChange(Number(newValue));
                }}
              />
              <Box sx={{ display: 'flex', justifyContent: 'center' }}>
                {state.fieldErrors?.rating && (
                  <span style={{ color: 'red' }}>
                    {state.fieldErrors.rating}
                  </span>
                )}
              </Box>
            </Box>
            {!book && (
              <BookSearchDialog
                formData={{
                  content: content,
                  rating: rating,
                  ISBNcode: ISBNcode,
                  postImage: postImage,
                }}
              />
            )}
            <label htmlFor="image-upload">
              <input
                type="file"
                id="image-upload"
                style={{ display: 'none' }}
                onChange={handleImageChange}
                accept="image/*"
              />
              <IconButton component="span">
                <ImageIcon />
              </IconButton>
            </label>
            <Button
              type="submit"
              disabled={!content || !rating || !ISBNcode}
              sx={{
                borderRadius: '50px',
                backgroundColor: '#FF7E73',
                color: '#fff',
                marginTop: '0.8rem',
                marginLeft: 'auto!important',
                '&:hover': {
                  backgroundColor: '#E56A67',
                },
                '&.Mui-disabled': {
                  backgroundColor: '#FFA49D',
                  color: '#fff',
                },
              }}
            >
              Post
            </Button>
          </DialogActions>
          <input type="hidden" value={book ? book.ISBNcode : ''} />
          <Box sx={{ display: 'flex', justifyContent: 'center' }}>
            {state.fieldErrors?.content && (
              <span style={{ color: 'red' }}>{state.fieldErrors.content}</span>
            )}
          </Box>
          <Box sx={{ display: 'flex', justifyContent: 'center' }}>
            {state.fieldErrors?.ISBNcode && (
              <span style={{ color: 'red' }}>{state.fieldErrors.ISBNcode}</span>
            )}
          </Box>
        </form>
      </Dialog>
    </React.Fragment>
  );
};
