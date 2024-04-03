import * as React from 'react';
import { apiInstance } from '../../lib/api/apiInstance';
import { Post } from '../../openapi';
import { ListSection } from './ListSection';

type PostListProps = {
  propPosts: Post[];
};
export const PostList: React.FC<PostListProps> = async ({ propPosts }) => {
  const fetchLikedPostIDs = async (): Promise<number[]> => {
    try {
      const api = apiInstance;
      const res = await api.getLikedPostList();
      if (res.data && Array.isArray(res.data)) {
        const fetchedLikedPostIDs: number[] = res.data.map(
          (item) => item.post_id,
        );
        return fetchedLikedPostIDs;
      } else {
        return Promise.reject(new Error('Failed to fetch liked posts'));
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const likedPostIDs = await fetchLikedPostIDs();

  return <ListSection propPosts={propPosts} likedPostIDs={likedPostIDs} />;
};
