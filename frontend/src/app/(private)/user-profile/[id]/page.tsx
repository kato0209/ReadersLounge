import UserProfile from '../../../../features/userProfile/UserProfile';
import { authenticate } from '../../../../lib/auth/authenticate';
import { redirect } from 'next/navigation';

export default function UserProfilePage({
  params,
}: {
  params: { id: string };
}) {
  const userID = Number(params.id);
  if (!authenticate()) {
    redirect('/login');
  }
  return <UserProfile userID={userID} />;
}
