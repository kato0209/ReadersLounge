import { Home } from '../../../features/home/Home';
import { authenticate } from '../../../lib/auth/authenticate';
import { redirect } from 'next/navigation';

export default function HomePage() {
  if (!authenticate()) {
    redirect('/login');
  }
  return <Home />;
}
