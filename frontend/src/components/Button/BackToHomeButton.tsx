import React from 'react';
import Button from '@mui/material/Button';

function BackToHomeButton() {
  return (
    <Button
      fullWidth
      variant="contained"
      sx={{
        backgroundColor: '#FF7E73',
        '&:hover': { backgroundColor: '#E56A67', color: '#fff' },
      }}
      href="/"
    >
      ホームへ戻る
    </Button>
  );
}

export default BackToHomeButton;
