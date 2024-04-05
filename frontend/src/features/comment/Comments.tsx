import Sidebar from '../../components/Sidebar/Sidebar';
import { Box } from '@mui/material';
import { CommentSC } from './Comment';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useSearchParams } from 'next/navigation';

export default function Comments() {
  const isMobile = useMediaQuery('(max-width:650px)');
  const searchParams = useSearchParams();
  const postID = Number(searchParams.get('id'));

  return (
    <>
      {!isMobile ? (
        <Box style={{ display: 'flex' }}>
          <Box style={{ flex: '0 0 30%', display: 'flex' }}>
            <Sidebar />
          </Box>
          <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
            <CommentSC postID={postID} />
          </Box>
          <Box style={{ flex: '0 0 30%' }}></Box>
        </Box>
      ) : (
        <Box style={{ display: 'flex', justifyContent: 'center' }}>
          <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
            <CommentSC postID={postID} />
          </Box>
        </Box>
      )}
    </>
  );
}
