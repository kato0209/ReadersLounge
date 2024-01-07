import * as React from 'react';
import { Box, Card, CardContent, Typography, Button, Avatar, Stack, CardMedia } from '@mui/material';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';
import { useAuthUserContext } from '../../lib/auth/auth';
import { isValidUrl } from '../../utils/isValidUrl';
import AddAPhotoIcon from '@mui/icons-material/AddAPhoto';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { set, useForm } from 'react-hook-form';
import TextField from '@mui/material/TextField';
import { User } from '../../openapi';
import { useErrorHandler } from 'react-error-boundary';
import { apiInstance } from '../../lib/api/apiInstance';
import DialogTitle from '@mui/material/DialogTitle';

type EditProfileProps  = {
    user: User;
    fetchUser: () => Promise<void>;
};


const IMAGE_TYPES = ['image/jpeg', 'image/png'];

const ProfileSchema = z.object({
  name: z.string().max(20, {
    message: "名前は20文字以内で入力してください"
  }),
  profileText: z.string().max(255, {
    message: "プロフィールは255文字以内で入力してください"
  }),
  profileImage: z.instanceof(File).optional()
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

type FormData = z.infer<typeof ProfileSchema>;

export const EditProfile: React.FC<EditProfileProps> = ({ user, fetchUser }) => {
    const { register, handleSubmit, setValue, getValues, formState: { errors } } = useForm<FormData>({
        resolver: zodResolver(ProfileSchema),
    });
    const [openUpdateProfileDialog, setOpenUpdateProfileDialog] = React.useState(false);
    const { user: loginUser } = useAuthUserContext();
    const errorHandler = useErrorHandler();
    const [imagePreview, setImagePreview] = React.useState<string | null>(null);
    const inputRef = React.useRef<HTMLInputElement>(null);

    const handleOpen  = () => {
        setOpenUpdateProfileDialog(true);
    };
    
    const handleClose = () => {
        setOpenUpdateProfileDialog(false);
    };

    const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files[0]) {
          const file = event.target.files[0];
          setValue('profileImage', file, { shouldValidate: true });
          const reader = new FileReader();
          reader.onloadend = () => {
            setImagePreview(reader.result as string);
          };
          reader.readAsDataURL(file);
        }
    };

    const handleImageButtonClick = () => {
        inputRef.current?.click();
    }

    const onSubmit = async (data: FormData) => {
        try {
            const api = await apiInstance;
            if (data.profileImage) {
                await api.updateUser(
                    data.name,
                    data.profileImage,
                    data.profileText,
                );
              } else {
                await api.updateUser(
                    data.name,
                    undefined,
                    data.profileText,
                );
              }
            handleClose();
            await fetchUser();
        } catch (error: unknown) {
          errorHandler(error);
        }
    }


    return (
        <>
            <Button 
                variant="outlined" 
                sx={{color: "black", 
                    borderColor: "black",
                    "&:hover": {
                        borderColor: "black", 
                        color: 'black', 
                        backgroundColor: "rgba(0, 0, 0, 0.1)" 
                    }  
                }}
                onClick={handleOpen}
            >
                Edit profile
            </Button>
            <Dialog 
                open={openUpdateProfileDialog} 
                onClose={handleClose} 
                fullWidth
                sx={{
                    '& .MuiDialog-paper': {
                    backgroundColor: '#EFEBE5',
                    }
                }}  
            >
                <DialogTitle sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                    <IconButton
                        aria-label="close"
                        onClick={handleClose}
                        sx={{
                            marginRight: 1,
                        }}
                    >
                        <CloseIcon />
                    </IconButton>
                    <Typography sx={{fontSize: "1.4rem", fontWeight: "bold"}}>Edit Profile</Typography>
                    <Button 
                        variant="contained" 
                        sx={{
                            backgroundColor: '#FF7E73',
                            color: '#fff',
                            '&:hover': {
                                backgroundColor: '#E56A67',
                            },
                        }}
                        onClick={handleSubmit(onSubmit)}
                    >
                        Save
                    </Button>
                </DialogTitle>
                <DialogContent
                >
                    <Box 
                        component="form" 
                        noValidate 
                        sx={{
                            flex: 'auto',
                            marginTop: '2rem',
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                        }}>
                        <label htmlFor="image-upload">
                            <input
                                ref={inputRef}
                                type="file"
                                id="image-upload"
                                style={{ display: 'none' }}
                                onChange={handleImageChange}
                                accept="image/*"
                            />
                            <Button 
                                sx={{ width: 128, height: 128, borderRadius: "50%" }}
                                onClick={handleImageButtonClick}
                            >
                                <Stack direction="row" justifyContent="center" alignItems="center">
                                    {imagePreview ? (
                                        <Avatar
                                            sx={{ 
                                                width: 128, 
                                                height: 128, 
                                                border: '4px solid', 
                                                borderColor: 'background.paper',
                                            }}
                                            src={imagePreview}
                                        />
                                    ): (
                                        <Avatar
                                            sx={{ 
                                                width: 128, 
                                                height: 128, 
                                                border: '4px solid', 
                                                borderColor: 'background.paper',
                                            }}
                                            src={isValidUrl(user.profile_image) ? user.profile_image : `data:image/png;base64,${user.profile_image}`}
                                        />
                                    )}
                                    <AddAPhotoIcon 
                                        sx={{ 
                                            position: 'absolute', 
                                            color: 'rgba(255, 255, 255, 1)', 
                                            fontSize: 'large'
                                        }} 
                                    />
                                </Stack>
                            </Button>
                        </label>
                        <TextField
                            {...register("name")}
                            margin="normal"
                            fullWidth
                            name="name"
                            label="Name"
                            defaultValue={user.name}
                            id='name'
                        />
                        {errors.name && <span style={{ color: 'red' }}>{errors.name.message}</span>}
                        <TextField
                            {...register("profileText")}
                            margin="normal"
                            fullWidth
                            name="profileText"
                            label="Profile"
                            defaultValue={user.profile_text}
                            id='profileText'
                        />
                        {errors.profileText && <span style={{ color: 'red' }}>{errors.profileText.message}</span>}
                    </Box>
                </DialogContent>
        </Dialog>
        </>
    );
}