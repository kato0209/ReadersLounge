import Comments from '../../../../features/comment/Comments';
import { authenticate } from '../../../../lib/auth/authenticate';
import { redirect } from 'next/navigation';

export default function PostPage(params: { [key: string]: string }) {
  const postID = Number(params.id);
  if (!authenticate()) {
    redirect('/login');
  }
  return <Comments postID={postID} />;
}
