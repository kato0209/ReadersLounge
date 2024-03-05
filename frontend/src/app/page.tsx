import { authenticate } from '../lib/auth/authenticate';
import { redirect } from 'next/navigation';
import Home from '../features/home/Home';

export default function HomePage() {
  if (!authenticate()) {
    redirect('/login');
  }
  return <Home />;
}
