import React from 'react';
import Button from '@mui/material/Button';

const submitButtonStyles = {
  mt: 3,
  mb: 2,
  backgroundColor: '#FF7E73',
  textTransform: 'none',
  '&:hover': {
    backgroundColor: '#E56A67',
  },
};

function SubmitButton(props: { content: string }) {
  return (
    <Button type="submit" fullWidth variant="contained" sx={submitButtonStyles}>
      {props.content}
    </Button>
  );
}

export default SubmitButton;
