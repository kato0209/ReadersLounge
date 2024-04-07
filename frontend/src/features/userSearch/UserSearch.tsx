import Sidebar from '../../components/Sidebar/Sidebar';
import UserSearchSection from './UserSearchSection';
import { Box } from '@mui/material';
import useMediaQuery from '@mui/material/useMediaQuery';

export default function UserSearch() {
  const isMobile = useMediaQuery('(max-width:650px)');

  return (
    <>
      {!isMobile ? (
        <Box style={{ display: 'flex' }}>
          <Box style={{ flex: '0 0 30%', display: 'flex' }}>
            <Sidebar />
          </Box>
          <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
            <UserSearchSection />
          </Box>
        </Box>
      ) : (
        <Box style={{ display: 'flex', justifyContent: 'center' }}>
          <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
            <UserSearchSection />
          </Box>
        </Box>
      )}
    </>
  );
}
