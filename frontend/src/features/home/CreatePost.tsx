import * as React from 'react';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import Button from '@mui/material/Button';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';
import { useAuthUserContext } from '../../lib/auth/auth';
import Avatar from '@mui/material/Avatar';
import { isValidUrl } from '../../utils/isValidUrl';
import Box from '@mui/material/Box';
import PostTextarea from './PostTextarea';
import Rating from '@mui/material/Rating';
import ImageIcon from '@mui/icons-material/Image';
import { z } from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { Book } from '../../openapi';
import ImportContactsIcon from '@mui/icons-material/ImportContacts';
import { BookSearchDialog } from './BookSearchDialog';
import { PostSchema } from '../../types/PostSchema';

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
  const [openCreatePostDialog, setOpenCreatePostDialog] = React.useState(false);
  const { user } = useAuthUserContext();
  const [imagePreview, setImagePreview] = React.useState<string | null>(null);

  const {
    register,
    handleSubmit,
    setValue,
    getValues,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(PostSchema),
  });

  React.useEffect(() => {
    if (formData?.content) {
      setValue('content', formData.content);
    }
  }, [formData?.content]);

  React.useEffect(() => {
    if (formData?.content) {
      setValue('rating', formData.rating);
    }
  }, [formData?.rating]);

  React.useEffect(() => {
    if (formData?.postImage) {
      setValue('postImage', formData.postImage);
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result as string);
      };
      reader.readAsDataURL(formData.postImage);
    }
  }, [formData?.postImage]);

  React.useEffect(() => {
    if (book) {
      setValue('ISBNcode', book.ISBNcode);
    }
  }, [book]);

  const errorHandler = useErrorHandler();

  const onSubmit = async (data: FormData) => {
    try {
      const api = await apiInstance;
      if (data.postImage) {
        await api.createPost(
          data.content,
          data.rating,
          data.ISBNcode,
          data.postImage,
        );
      } else {
        await api.createPost(data.content, data.rating, data.ISBNcode);
      }
      window.location.reload();
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];
      setValue('postImage', file, { shouldValidate: true });
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleImageRemove = () => {
    setValue('postImage', undefined, { shouldValidate: true });
    setImagePreview(null);
  };

  const handleRatingChange = (newValue: number) => {
    setValue('rating', newValue, { shouldValidate: true });
  };

  const handlePostTextChange = (
    event: React.ChangeEvent<HTMLTextAreaElement>,
  ) => {
    setValue('content', event.target.value, { shouldValidate: true });
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
        <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate>
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
                  isValidUrl(user.profile_image)
                    ? user.profile_image
                    : `data:image/png;base64,${user.profile_image}`
                }
              ></Avatar>
              <PostTextarea
                {...register('content')}
                onChange={handlePostTextChange}
                value={getValues('content')}
              />
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
                {errors.postImage && (
                  <span style={{ color: 'red', textAlign: 'center' }}>
                    {errors.postImage.message}
                  </span>
                )}
              </Box>
            )}
          </DialogContent>
          <DialogActions sx={{ justifyContent: 'start' }}>
            <Box>
              <Rating
                name="rating"
                value={Number(getValues('rating'))}
                onChange={(event, newValue) => {
                  handleRatingChange(Number(newValue));
                }}
              />
              <Box sx={{ display: 'flex', justifyContent: 'center' }}>
                {errors.rating && (
                  <span style={{ color: 'red' }}>{errors.rating.message}</span>
                )}
              </Box>
            </Box>
            {!book && (
              <BookSearchDialog
                formData={{
                  content: getValues('content'),
                  rating: getValues('rating'),
                  ISBNcode: getValues('ISBNcode'),
                  postImage: getValues('postImage'),
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
              disabled={
                !getValues('content') ||
                !getValues('rating') ||
                !getValues('ISBNcode')
              }
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
          <input
            {...register('ISBNcode')}
            type="hidden"
            value={book ? book.ISBNcode : ''}
          />
          <Box sx={{ display: 'flex', justifyContent: 'center' }}>
            {errors.content && (
              <span style={{ color: 'red' }}>{errors.content.message}</span>
            )}
          </Box>
          <Box sx={{ display: 'flex', justifyContent: 'center' }}>
            {errors.ISBNcode && (
              <span style={{ color: 'red' }}>{errors.ISBNcode.message}</span>
            )}
          </Box>
        </Box>
      </Dialog>
    </React.Fragment>
  );
};
