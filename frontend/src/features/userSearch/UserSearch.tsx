import * as React from 'react';
import Sidebar from '../../components/Sidebar/Sidebar';
import { useIsMobileContext } from '../../providers/mobile/isMobile';
import UserSearchComponent from './UserSearchComponent';
import { Box } from '@mui/material';

export default function UserSearch() {
  const isMobile = useIsMobileContext();

  return (
    <>
      {!isMobile ? (
        <Box style={{ display: 'flex' }}>
          <Box style={{ flex: '0 0 30%', display: 'flex' }}>
            <Sidebar />
          </Box>
          <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
            <UserSearchComponent />
          </Box>
        </Box>
      ) : (
        <Box style={{ display: 'flex', justifyContent: 'center' }}>
          <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
            <UserSearchComponent />
          </Box>
        </Box>
      )}
    </>
  );
}
