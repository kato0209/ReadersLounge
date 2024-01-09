import * as React from 'react';
import Sidebar from '../../components/Sidebar/Sidebar';
import UserProfileComponent from './UserProfileComponent';
import { useIsMobileContext } from '../../providers/mobile/isMobile';
import { Box } from '@mui/material';

export default function UserProfile() {
    const isMobile = useIsMobileContext();
    
  return (
    <>
        {!isMobile ? (
            <Box style={{ display: 'flex'}}>
                <Box style={{ flex: '0 0 30%', display: 'flex' }}>
                    <Sidebar />
                </Box>
                <Box style={{ flex: 1, overflowX: 'hidden' }}>
                    <UserProfileComponent />
                </Box>
            </Box>
        ): (
        <Box style={{ display: 'flex', justifyContent: "center"}}>
            <Box style={{ flex: '0 0 100%', overflowX: 'hidden' }}>
                <UserProfileComponent />
            </Box>
        </Box>
    )}
    </>
  );
}