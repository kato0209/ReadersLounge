import * as React from 'react';
import Sidebar from './Sidebar';
import PostList from './PostList';

export default function Home() {
    const [isMobile, setIsMobile] = React.useState(window.innerWidth <= 500);

    const handleResize = () => {
      setIsMobile(window.innerWidth <= 500);
    };

    React.useEffect(() => {
      window.addEventListener('resize', handleResize);
      return () => {
        window.removeEventListener('resize', handleResize);
      };
    }, []);

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