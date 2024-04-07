'use client';
import * as React from 'react';
import { Box, Typography, Button, Avatar, Stack } from '@mui/material';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';
import { isValidUrl } from '../../utils/isValidUrl';
import AddAPhotoIcon from '@mui/icons-material/AddAPhoto';
import TextField from '@mui/material/TextField';
import { User } from '../../openapi';
import DialogTitle from '@mui/material/DialogTitle';
import { useFormState } from 'react-dom';
import { State } from './ProfileEditAction';
import { profileEdit } from './ProfileEditAction';

const initialState: State = {
  error: '',
  fieldErrors: {
    name: '',
    profileText: '',
    profileImage: '',
  },
};

type EditProfileProps = {
  user: User;
};

export const EditProfile: React.FC<EditProfileProps> = ({ user }) => {
  const [state, formAction] = useFormState(profileEdit, initialState);
  const [openUpdateProfileDialog, setOpenUpdateProfileDialog] =
    React.useState(false);
  const [imagePreview, setImagePreview] = React.useState<string | null>(null);
  const inputRef = React.useRef<HTMLInputElement>(null);
  const [profileImage, setProfileImage] = React.useState<File | undefined>(
    undefined,
  );

  const handleOpen = () => {
    setOpenUpdateProfileDialog(true);
  };

  const handleClose = () => {
    setOpenUpdateProfileDialog(false);
  };

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];
      setProfileImage(file);
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleImageButtonClick = () => {
    inputRef.current?.click();
  };

  return (
    <>
      <Button
        variant="outlined"
        size="small"
        sx={{
          color: 'black',
          borderColor: 'black',
          '&:hover': {
            borderColor: 'black',
            color: 'black',
            backgroundColor: 'rgba(0, 0, 0, 0.1)',
          },
          '@media (max-width: 1050px)': {
            maxWidth: 60,
          },
          '@media (max-width: 770px)': {
            maxWidth: 30,
          },
          '@media (max-width: 600px)': {
            maxWidth: 30,
          },
        }}
        onClick={handleOpen}
      >
        Edit profile
      </Button>
      <form action={formAction}>
        <Dialog
          open={openUpdateProfileDialog}
          onClose={handleClose}
          fullWidth
          sx={{
            '& .MuiDialog-paper': {
              backgroundColor: '#EFEBE5',
            },
          }}
        >
          <DialogTitle
            sx={{
              display: 'flex',
              justifyContent: 'space-between',
              alignItems: 'center',
            }}
          >
            <IconButton
              aria-label="close"
              onClick={handleClose}
              sx={{
                marginRight: 1,
              }}
            >
              <CloseIcon />
            </IconButton>
            <Typography sx={{ fontSize: '1.4rem', fontWeight: 'bold' }}>
              Edit Profile
            </Typography>
            <Button
              variant="contained"
              sx={{
                backgroundColor: '#FF7E73',
                color: '#fff',
                '&:hover': {
                  backgroundColor: '#E56A67',
                },
              }}
              type="submit"
            >
              Save
            </Button>
          </DialogTitle>
          <DialogContent>
            <Box
              component="form"
              noValidate
              sx={{
                flex: 'auto',
                marginTop: '2rem',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
              }}
            >
              <label htmlFor="image-upload">
                <input
                  ref={inputRef}
                  type="file"
                  value={profileImage ? profileImage.name : ''}
                  id="image-upload"
                  name="profileImage"
                  style={{ display: 'none' }}
                  onChange={handleImageChange}
                  accept="image/*"
                />
                <Button
                  sx={{ width: 128, height: 128, borderRadius: '50%' }}
                  onClick={handleImageButtonClick}
                >
                  <Stack
                    direction="row"
                    justifyContent="center"
                    alignItems="center"
                  >
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
                    ) : (
                      <Avatar
                        sx={{
                          width: 128,
                          height: 128,
                          border: '4px solid',
                          borderColor: 'background.paper',
                        }}
                        src={
                          isValidUrl(user.profile_image)
                            ? user.profile_image
                            : `data:image/png;base64,${user.profile_image}`
                        }
                      />
                    )}
                    <AddAPhotoIcon
                      sx={{
                        position: 'absolute',
                        color: 'rgba(255, 255, 255, 1)',
                        fontSize: 'large',
                      }}
                    />
                  </Stack>
                </Button>
              </label>
              <TextField
                margin="normal"
                fullWidth
                name="name"
                label="Name"
                defaultValue={user.name}
                id="name"
              />
              {state.fieldErrors?.name && (
                <span style={{ color: 'red' }}>{state.fieldErrors?.name}</span>
              )}
              <TextField
                margin="normal"
                fullWidth
                name="profileText"
                label="Profile"
                defaultValue={user.profile_text}
                id="profileText"
              />
              {state.fieldErrors?.profileText && (
                <span style={{ color: 'red' }}>
                  {state.fieldErrors?.profileText}
                </span>
              )}
            </Box>
          </DialogContent>
        </Dialog>
      </form>
    </>
  );
};
