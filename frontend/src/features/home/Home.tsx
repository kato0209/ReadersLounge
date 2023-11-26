import * as React from 'react';
import Sidebar from './Sidebar';
import PostList from './PostList';
import useMediaQuery from '@mui/material/useMediaQuery';

export default function Home() {
    const isMobile = useMediaQuery('(max-width:500px)');

  return (
    <div style={{ display: 'flex'}}>
        {!isMobile && (
            <div style={{ flex: '0 0 30%', display: 'flex' }}>
                <Sidebar />
            </div>
        )}
        <div style={{ flex: 1, overflowX: 'hidden' }}>
            <PostList />
        </div>
    </div>
  );
}