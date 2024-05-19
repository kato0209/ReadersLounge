import UserSearch from '../../../features/userSearch/UserSearch';
import { authenticate } from '../../../lib/auth/authenticate';
import { redirect } from 'next/navigation';

export default function UserSearchPage() {
  if (!authenticate()) {
    redirect('/login');
  }
  return <UserSearch />;
}
