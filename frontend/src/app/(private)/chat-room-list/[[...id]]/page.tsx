import RoomList from '../../../../features/chat/RoomList';
import { authenticate } from '../../../../lib/auth/authenticate';
import { redirect } from 'next/navigation';

export default function SearchBookPage() {
  if (!authenticate()) {
    redirect('/login');
  }
  return <RoomList />;
}
