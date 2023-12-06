import * as React from 'react';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';

export default function UserProfile() {

  return (
    <Container component="main" maxWidth="xs">
        <Box
          sx={{
            marginTop: '8rem',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
            <h1>プロフィール</h1>
        </Box> 
    </Container>
  );
}