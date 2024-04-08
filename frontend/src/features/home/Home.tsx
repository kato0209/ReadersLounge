import Sidebar from '../../components/Sidebar/Sidebar';
import { PostList } from '../../components/PostList/PostList';
import { Box } from '@mui/material';
import { apiInstance } from '../../lib/api/apiInstance';
import { Post } from '../../openapi';
import { getAllCookies } from '../../utils/getCookies';

export const Home = async () => {
  const fetchPosts = async () => {
    try {
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getPosts({ headers: { Cookie: cookie } });
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
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const posts = await fetchPosts();

  return (
    <>
      <Box className="isMobile" style={{ display: 'flex' }}>
        <Box style={{ flex: '0 0 30%', display: 'flex' }}>
          <Sidebar />
        </Box>
        <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
          {posts && <PostList propPosts={posts} />}
        </Box>
      </Box>
      <Box
        className="isPC"
        style={{ display: 'flex', justifyContent: 'center' }}
      >
        <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
          {posts && <PostList propPosts={posts} />}
        </Box>
      </Box>
    </>
  );
};
