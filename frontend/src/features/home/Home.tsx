import Sidebar from '../../components/Sidebar/Sidebar';
import { PostList } from '../../components/PostList/PostList';
import { Box } from '@mui/material';
import { apiInstance } from '../../lib/api/apiInstance';
import { Post } from '../../openapi';
import useMediaQuery from '@mui/material/useMediaQuery';

export const Home = async () => {
  const isMobile = useMediaQuery('(max-width:650px)');

  const fetchPosts = async () => {
    const api = apiInstance;
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
      return fetchedPosts;
    }
  };

  const posts = await fetchPosts();

  return (
    <>
      {!isMobile ? (
        <Box style={{ display: 'flex' }}>
          <Box style={{ flex: '0 0 30%', display: 'flex' }}>
            <Sidebar />
          </Box>
          <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
            {posts && <PostList propPosts={posts} />}
          </Box>
        </Box>
      ) : (
        <Box style={{ display: 'flex', justifyContent: 'center' }}>
          <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
            {posts && <PostList propPosts={posts} />}
          </Box>
        </Box>
      )}
    </>
  );
};
