import * as React from 'react';
import Sidebar from '../../components/Sidebar/Sidebar';
import PostList from './PostList';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useIsMobileContext } from '../../providers/mobile/isMobile';

export default function Home() {
    const isMobile = useIsMobileContext();

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