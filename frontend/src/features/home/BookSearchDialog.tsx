'use client';
import * as React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import CloseIcon from '@mui/icons-material/Close';
import Slide from '@mui/material/Slide';
import { TransitionProps } from '@mui/material/transitions';
import { SearchBook } from '../searchBook/SearchBook';
import { PostSchema } from '../../types/PostSchema';
import { z } from 'zod';

const Transition = React.forwardRef(function Transition(
  props: TransitionProps & {
    children: React.ReactElement;
  },
  ref: React.Ref<unknown>,
) {
  return <Slide direction="up" ref={ref} {...props} />;
});

type PostFormData = z.infer<typeof PostSchema>;

type BookSearchDialogProps = {
  formData?: PostFormData;
};

export const BookSearchDialog: React.FC<BookSearchDialogProps> = ({
  formData,
}) => {
  const [open, setOpen] = React.useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <React.Fragment>
      <Button
        sx={{
          borderRadius: '100px',
          backgroundColor: '#4d4d4d',
          color: '#fff',
          '&:hover': {
            backgroundColor: '#808080',
          },
        }}
        onClick={handleClickOpen}
      >
        本を選択する
      </Button>
      <Dialog
        fullScreen
        open={open}
        onClose={handleClose}
        TransitionComponent={Transition}
        sx={{
          '& .MuiDialog-paper': {
            backgroundColor: '#EFEBE5',
          },
        }}
      >
        <AppBar sx={{ position: 'relative', backgroundColor: '#FF7E73' }}>
          <Toolbar>
            <IconButton
              edge="start"
              color="inherit"
              onClick={handleClose}
              aria-label="close"
            >
              <CloseIcon />
            </IconButton>
            <Typography sx={{ ml: 2, flex: 1 }} variant="h6" component="div">
              本を選択してください
            </Typography>
          </Toolbar>
        </AppBar>
        <SearchBook formData={formData} />
      </Dialog>
    </React.Fragment>
  );
};
