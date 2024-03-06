import type { ReactNode } from 'react';
import AppHeader from '../../components/Header/AppHeader';
import Box from '@mui/material/Box';

export default function SignupLayout({ children }: { children: ReactNode }) {
  return (
    <Box sx={{ pt: '3rem' }}>
      <AppHeader />
      {children}
    </Box>
  );
}
