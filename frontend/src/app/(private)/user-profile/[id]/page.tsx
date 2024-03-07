import UserProfile from '../../../../features/userProfile/UserProfile';
import { authenticate } from '../../../../lib/auth/authenticate';
import { redirect } from 'next/navigation';

export default function UserSearchPage() {
  if (!authenticate()) {
    redirect('/login');
  }
  return <UserProfile />;
}
