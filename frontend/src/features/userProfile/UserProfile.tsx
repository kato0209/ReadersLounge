import Sidebar from '../../components/Sidebar/Sidebar';
import UserProfileComponent from './UserProfileComponent';
import { Box } from '@mui/material';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useSearchParams } from 'next/navigation';
import { apiInstance } from '../../lib/api/apiInstance';
import { Connection, Post, User } from '../../openapi';

export default async function UserProfile() {
  const isMobile = useMediaQuery('(max-width:650px)');

  const searchParams = useSearchParams();
  const id = searchParams.get('id');
  const userID = id ? parseInt(id, 10) : 0;

  const fetchUser = async (userID: number): Promise<User> => {
    try {
      const api = apiInstance;
      const res = await api.getUser(userID);
      if (res.data) {
        const targetUser: User = {
          user_id: res.data.user_id,
          name: res.data.name,
          profile_image: res.data.profile_image,
          profile_text: res.data.profile_text,
        };
        return targetUser;
      } else {
        return Promise.reject('User not found');
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const fetchfollowerConnections = async (
    userID: number,
  ): Promise<Connection[]> => {
    try {
      const api = apiInstance;
      const res = await api.getFollowerConnections(userID);
      if (res.data && res.data.length > 0) {
        const followerConnections: Connection[] = res.data.map((connection) => {
          return {
            connection_id: connection.connection_id,
            target_user_id: connection.target_user_id,
            target_user_name: connection.target_user_name,
            target_user_profile_image: connection.target_user_profile_image,
          };
        });
        return followerConnections;
      } else {
        return [];
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const fetchfollowingConnections = async (
    userID: number,
  ): Promise<Connection[]> => {
    try {
      const api = apiInstance;
      const res = await api.getFollowingConnections(userID);
      if (res.data && res.data.length > 0) {
        const followingConnections: Connection[] = res.data.map(
          (connection) => {
            return {
              connection_id: connection.connection_id,
              target_user_id: connection.target_user_id,
              target_user_name: connection.target_user_name,
              target_user_profile_image: connection.target_user_profile_image,
            };
          },
        );
        return followingConnections;
      } else {
        return [];
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const fetchPostsOfUser = async (userID: number): Promise<Post[]> => {
    try {
      const api = apiInstance;
      const res = await api.getPostsOfUser(userID);
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
      } else {
        return [];
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const user = await fetchUser(userID);
  const followerConnections = await fetchfollowerConnections(userID);
  const followingConnections = await fetchfollowingConnections(userID);
  const posts = await fetchPostsOfUser(userID);

  return (
    <>
      {!isMobile ? (
        <Box style={{ display: 'flex' }}>
          <Box style={{ flex: '0 0 30%', display: 'flex' }}>
            <Sidebar />
          </Box>
          <Box style={{ flex: 1, overflowX: 'hidden' }}>
            <UserProfileComponent
              user={user}
              followerConnections={followerConnections}
              followingConnections={followingConnections}
              posts={posts}
            />
          </Box>
        </Box>
      ) : (
        <Box style={{ display: 'flex', justifyContent: 'center' }}>
          <Box style={{ flex: '0 0 100%', overflowX: 'hidden' }}>
            <UserProfileComponent
              user={user}
              followerConnections={followerConnections}
              followingConnections={followingConnections}
              posts={posts}
            />
          </Box>
        </Box>
      )}
    </>
  );
}
