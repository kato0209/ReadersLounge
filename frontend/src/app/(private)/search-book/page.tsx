import { SearchBook } from '../../../features/searchBook/SearchBook';
import { authenticate } from '../../../lib/auth/authenticate';
import { redirect } from 'next/navigation';

export default function SearchBookPage() {
  if (!authenticate()) {
    redirect('/login');
  }
  return <SearchBook />;
}
