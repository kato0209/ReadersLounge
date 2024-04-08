import Sidebar from '../../components/Sidebar/Sidebar';
import { Box } from '@mui/material';
import { CommentSC } from './Comment';

export default function Comments({ postID }: { postID: number }) {
  return (
    <>
      <Box className="isMobile" style={{ display: 'flex' }}>
        <Box style={{ flex: '0 0 30%', display: 'flex' }}>
          <Sidebar />
        </Box>
        <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
          <CommentSC postID={postID} />
        </Box>
        <Box style={{ flex: '0 0 30%' }}></Box>
      </Box>
      <Box
        className="isPC"
        style={{ display: 'flex', justifyContent: 'center' }}
      >
        <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
          <CommentSC postID={postID} />
        </Box>
      </Box>
    </>
  );
}
