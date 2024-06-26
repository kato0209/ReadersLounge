import Sidebar from '../../components/Sidebar/Sidebar';
import UserSearchSection from './UserSearchSection';
import { Box } from '@mui/material';

export default function UserSearch() {
  return (
    <>
      <Box className="isMobile" style={{ display: 'flex', marginTop: '1rem' }}>
        <Box style={{ flex: '0 0 30%', display: 'flex' }}>
          <Sidebar />
        </Box>
        <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
          <UserSearchSection />
        </Box>
      </Box>
      <Box
        className="isPC"
        style={{ display: 'flex', justifyContent: 'center', marginTop: '1rem' }}
      >
        <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
          <UserSearchSection />
        </Box>
      </Box>
    </>
  );
}
