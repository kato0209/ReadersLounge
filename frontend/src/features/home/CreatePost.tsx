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
import { zfd } from "zod-form-data";
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { get } from 'http';
import { error } from 'console';

const IMAGE_TYPES = ['image/jpeg', 'image/png'];

const PostSchema = z.object({
  content: z.string().nonempty('投稿内容は必須です').max(255, {
    message: "投稿内容は255文字以内で入力してください"
  }),
  rating: z.number().positive(),
  ISBNcode: z.string().nonempty('本が選択されていません'), 
  postImage: z.instanceof(File).optional()
  .refine((file) => {
    return (
      file === undefined ||
      (IMAGE_TYPES.includes(file.type) &&
        file.name.split('.').pop()?.toLowerCase() !== 'jpg')
    );
  }, {
    message: '.jpegもしくは.pngのみ可能です',
  })
});

type FormData = z.infer<typeof PostSchema>;

export default function CreatePost() {
  const [openCreatePostDialog, setOpenCreatePostDialog] = React.useState(false);
  const { user } = useAuthUserContext();
  const [imagePreview, setImagePreview] = React.useState<string | null>(null);

  const { register, handleSubmit, setValue, getValues, formState: { errors } } = useForm<FormData>({
    resolver: zodResolver(PostSchema),
  });
  const errorHandler = useErrorHandler();

  const onSubmit = async (data: FormData) => {
    
    try {
      const api = await apiInstance;

      if (data.postImage) {
      const res = await api.createPost(
        data.content,
        data.rating,
        data.ISBNcode,
        data.postImage
      );
    } else {
      const res = await api.createPost(
        data.content,
        data.rating,
        data.ISBNcode
      );
    }
    window.location.reload();
    } catch (error: unknown) {
      errorHandler(error);
    }
  }
  

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

  const handlePostTextChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setValue('content', event.target.value, { shouldValidate: true });
  };
  

  const handleOpen  = () => {
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
        <ListItemText primary="Post" sx={{ textAlign: 'center', paddingHorizontal: '0.5rem' }} />
        </ListItem>

        <Dialog 
          open={openCreatePostDialog} 
          onClose={handleClose} 
          fullWidth
          sx={{
            '& .MuiDialog-paper': {
              backgroundColor: '#EFEBE5',
            }
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
                <Box sx={{ display: 'flex', alignItems: 'start', gap: 2 }}>
                    <Avatar 
                        src={isValidUrl(user.profile_image) ? user.profile_image : `data:image/png;base64,${user.profile_image}` }>                  
                    </Avatar>
                    <PostTextarea {...register("content")} onChange={handlePostTextChange} value={getValues('content')}/>
                </Box>
                {imagePreview && (
                <Box sx={{ display: 'flex', flexDirection: 'column', justifyContent: 'center' }}>
                    <Box sx={{ display: 'flex', justifyContent: 'center', marginTop: '1rem', position: 'relative' }}>
                      <img src={imagePreview} style={{ width: '50%', height: 'auto' }}/>
                      
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
                    {errors.postImage && <span style={{ color: 'red', textAlign: 'center' }}>{errors.postImage.message}</span>}
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
                    <Box sx={{display: 'flex', justifyContent: 'center'}}>
                    {errors.rating && <span style={{ color: 'red' }}>{errors.rating.message}</span>}
                    </Box>
                  </Box>
                  <label htmlFor="image-upload">
                    <input
                      type="file"
                      id="image-upload"
                      style={{ display: 'none' }}
                      onChange={handleImageChange}
                      accept="image/*"
                    />
                    <IconButton component='span'>
                      <ImageIcon />
                    </IconButton>
                  </label>
                  <Button 
                    type="submit"
                    disabled={ !getValues('content') || !getValues('rating') || !getValues('ISBNcode')}
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
                      }
                    }}
                  >
                      Post
                  </Button>
              </DialogActions>
              <input {...register("ISBNcode")} type="hidden" value="9784472405433" />
              <Box sx={{display: 'flex', justifyContent: 'center'}}>
                {errors.content && <span style={{ color: 'red' }}>{errors.content.message}</span>}
              </Box>
              <Box sx={{display: 'flex', justifyContent: 'center'}}>
              {errors.ISBNcode && <span style={{ color: 'red' }}>{errors.ISBNcode.message}</span>}
              </Box>
            </Box>
        </Dialog>
    </React.Fragment>
  );
}