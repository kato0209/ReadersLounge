import Sidebar from '../../components/Sidebar/Sidebar';
import UserProfileComponent from './UserProfileComponent';
import { Box } from '@mui/material';
import { apiInstance } from '../../lib/api/apiInstance';
import { Connection, Post, User } from '../../openapi';
import { getAllCookies } from '../../utils/getCookies';
import { PostList } from '../../components/PostList/PostList';

export default async function UserProfile({ userID }: { userID: number }) {
  const fetchUser = async (userID: number): Promise<User> => {
    try {
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getUser(userID, { headers: { Cookie: cookie } });
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
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getFollowerConnections(userID, {
        headers: { Cookie: cookie },
      });
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
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getFollowingConnections(userID, {
        headers: { Cookie: cookie },
      });
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
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getPostsOfUser(userID, {
        headers: { Cookie: cookie },
      });
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
      <Box className="isMobile" style={{ display: 'flex' }}>
        <Box style={{ flex: '0 0 30%', display: 'flex' }}>
          <Sidebar />
        </Box>
        <Box style={{ flex: 1, overflowX: 'hidden' }}>
          <UserProfileComponent
            user={user}
            initialFollowerConnections={followerConnections}
            initialFollowingConnections={followingConnections}
            posts={posts}
            postListComponent={<PostList propPosts={posts} />}
          />
        </Box>
      </Box>
      <Box
        className="isPC"
        style={{ display: 'flex', justifyContent: 'center' }}
      >
        <Box style={{ flex: '0 0 100%', overflowX: 'hidden' }}>
          <UserProfileComponent
            user={user}
            initialFollowerConnections={followerConnections}
            initialFollowingConnections={followingConnections}
            posts={posts}
            postListComponent={<PostList propPosts={posts} />}
          />
        </Box>
      </Box>
    </>
  );
}
