import * as React from 'react';
import { apiInstance } from '../../lib/api/apiInstance';
import { Post, Comment } from '../../openapi';
import { CommentCC } from './CommentCC';
import { getAllCookies } from '../../utils/getCookies';

export async function CommentSC({ postID }: { postID: number }) {
  const fetchPost = async (postID: number): Promise<Post> => {
    try {
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getPostByPostID(postID, {
        headers: { Cookie: cookie },
      });
      if (res.data) {
        const post: Post = {
          post_id: res.data.post_id,
          user: res.data.user,
          content: res.data.content,
          rating: res.data.rating,
          image: res.data.image,
          created_at: res.data.created_at,
          book: res.data.book,
          likes: res.data.likes,
        };
        return post;
      } else {
        return Promise.reject('Post not found');
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const fetchComments = async (postID: number): Promise<Comment[]> => {
    try {
      const cookie = getAllCookies();
      const api = await apiInstance;
      const res = await api.getCommentsByPostID(postID, {
        headers: { Cookie: cookie },
      });
      if (res.data) {
        const comments: Comment[] = res.data.map((item) => ({
          comment_id: item.comment_id,
          user: item.user,
          content: item.content,
          likes: item.likes,
          created_at: item.created_at,
        }));
        return comments;
      } else {
        return Promise.reject('Comments not found');
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const fetchLikedPostIDs = async (): Promise<number[]> => {
    try {
      const cookie = getAllCookies();
      const api = await apiInstance;
      const res = await api.getLikedPostList({ headers: { Cookie: cookie } });
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

  const fetchLikedCommentIDs = async (): Promise<number[]> => {
    try {
      const cookie = getAllCookies();
      const api = await apiInstance;
      const res = await api.getLikedCommentList({
        headers: { Cookie: cookie },
      });
      if (res.data && Array.isArray(res.data)) {
        const fetchedLikedCommentIDs: number[] = res.data.map(
          (item) => item.comment_id,
        );
        return fetchedLikedCommentIDs;
      } else {
        return Promise.reject(new Error('Failed to fetch liked comments'));
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const post = await fetchPost(postID);
  const comments = await fetchComments(postID);
  const likedPostIDs = await fetchLikedPostIDs();
  const likedCommentIDs = await fetchLikedCommentIDs();

  return (
    <CommentCC
      post={post}
      initialComments={comments}
      initialLikedPostIDs={likedPostIDs}
      initialLikedCommentIDs={likedCommentIDs}
    />
  );
}
