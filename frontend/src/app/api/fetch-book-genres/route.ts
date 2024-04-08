import { apiInstance } from '../../../lib/api/apiInstance';
import { BookGenreNode } from '../../../openapi';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET(): Promise<BookGenreNode[]> {
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.getBooksGenres({ headers: { Cookie: cookie } });
    if (res.data && Array.isArray(res.data)) {
      const fetchedBookGenres: BookGenreNode[] = res.data.map((item) => ({
        id: item.id,
        books_genre_id: item.books_genre_id,
        books_genre_name: item.books_genre_name,
        genre_level: item.genre_level,
        parent_genre_id: item.parent_genre_id,
        children: item.children,
      }));
      return fetchedBookGenres;
    } else {
      return [];
    }
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
