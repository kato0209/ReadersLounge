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

const PostSchema = z.object({
  content: z.string().nonempty('投稿内容は必須です').max(255, {
    message: "投稿内容は255文字以内で入力してください"
  }),
  rating: z.number().positive(),
  ISBNcode: z.string().nonempty('本が選択されていません'), 
  postImage: z.any().optional()
  .refine(
    (files) => {
      if (!files || files.length === 0) {
        return true;
      }
      ['image/jpeg', 'image/png'].includes(files?.[0]?.type),
    '.jpg, .jpeg, .pngのファイルを選択してください。'
    }
  )
});

type FormData = z.infer<typeof PostSchema>;

export default function CreatePost() {
  const [openCreatePostDialog, setOpenCreatePostDialog] = React.useState(false);
  const { user } = useAuthUserContext();
  const [image, setImage] = React.useState<File | null>(null);
  const [imagePreview, setImagePreview] = React.useState<string | null>(null);

  const { register, handleSubmit, setValue, getValues, formState: { errors } } = useForm<FormData>({
    resolver: zodResolver(PostSchema),
  });
  const errorHandler = useErrorHandler();

  const onSubmit = async (data: FormData) => {
    console.log("onSubmit");
    console.log(data);
    try {
      const api = await apiInstance;
      const res = await api.createPost(
        data.content,
        data.rating,
        data.ISBNcode,
        data.postImage
      );
    } catch (error: unknown) {
      errorHandler(error);
    }
  }
  

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];
      setImage(file);
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleRatingChange = (newValue: number) => {
    setValue('rating', newValue, { shouldValidate: true });
  };

  const handlePostTextChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    //setPostText(event.target.value);
    setValue('content', event.target.value, { shouldValidate: true });
  };
  

  const handleOpen  = () => {
    setOpenCreatePostDialog(true);
  };

  const handleClose = () => {
    setOpenCreatePostDialog(false);
  };

  const handleImageRemove = () => {
    setImage(null);
    setImagePreview(null);
  };

  return (
    <React.Fragment>
      {errors.postImage && <p>lklk</p>}
        <ListItem 
        button 
        sx={{ 
            borderRadius: '50px', 
            backgroundColor: '#FF7E73',
            color: '#fff',
            marginTop: '0.8rem',
            '&:hover': {
            backgroundColor: '#E56A67',
            },
        }}
        onClick={handleOpen}
        >
        <ListItemText primary="Post" sx={{ textAlign: 'center' }} />
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
                )}
              </DialogContent>
              <DialogActions sx={{ justifyContent: 'start' }}>
                  <Rating
                      name="rating"
                      value={Number(getValues('rating'))}
                      onChange={(event, newValue) => {
                        handleRatingChange(Number(newValue));
                      }}
                  />
                  <label htmlFor="image-upload">
                    <input
                      {...register("postImage")}
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
                    disabled={!getValues('content') || !getValues('rating') || !getValues('ISBNcode')}
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
            </Box>
        </Dialog>
    </React.Fragment>
  );
}