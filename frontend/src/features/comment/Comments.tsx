import Sidebar from '../../components/Sidebar/Sidebar';
import { useIsMobileContext } from '../../providers/mobile/isMobile';
import { Box } from '@mui/material';
import { CommentComponent } from './CommentComponent';

export default function Comments() {
  const isMobile = useIsMobileContext();

  return (
    <>
      {!isMobile ? (
        <Box style={{ display: 'flex' }}>
          <Box style={{ flex: '0 0 30%', display: 'flex' }}>
            <Sidebar />
          </Box>
          <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
            <CommentComponent />
          </Box>
          <Box style={{ flex: '0 0 30%' }}></Box>
        </Box>
      ) : (
        <Box style={{ display: 'flex', justifyContent: 'center' }}>
          <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
            <CommentComponent />
          </Box>
        </Box>
      )}
    </>
  );
}
