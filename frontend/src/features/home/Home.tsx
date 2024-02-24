import { useState, useEffect } from 'react';
import Sidebar from '../../components/Sidebar/Sidebar';
import { PostList } from '../../components/PostList/PostList';
import { useIsMobileContext } from '../../providers/mobile/isMobile';
import { Box } from '@mui/material';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { Post } from '../../openapi';

export default function Home() {
  const isMobile = useIsMobileContext();
  const errorHandler = useErrorHandler();
  const [posts, setPosts] = useState<Post[]>([]);

  const fetchPosts = async () => {
    try {
      const api = await apiInstance;
      const res = await api.getPosts();

      if (res.data && Array.isArray(res.data)) {
        const fetchedPosts: Post[] = res.data.map((item) => ({
          post_id: item.post_id,
          user: item.user,
          content: item.content,
          rating: item.rating,
          image: item.image,
          created_at: item.created_at,
          book: item.book,
          likes: item.likes,
        }));
        setPosts(fetchedPosts);
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  return (
    <>
      {!isMobile ? (
        <Box style={{ display: 'flex' }}>
          <Box style={{ flex: '0 0 30%', display: 'flex' }}>
            <Sidebar />
          </Box>
          <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
            <PostList propPosts={posts} />
          </Box>
        </Box>
      ) : (
        <Box style={{ display: 'flex', justifyContent: 'center' }}>
          <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
            <PostList propPosts={posts} />
          </Box>
        </Box>
      )}
    </>
  );
}
