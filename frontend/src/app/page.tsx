import { authenticate } from '../lib/auth/authenticate';
import { redirect } from 'next/navigation';

export default function RootPage() {
  if (!authenticate()) {
    redirect('/login');
  }
  redirect('/home');
}
