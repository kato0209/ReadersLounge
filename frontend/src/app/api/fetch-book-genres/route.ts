import { apiInstance } from '../../../lib/api/apiInstance';
import { BookGenreNode } from '../../../openapi';
import { getAllCookies } from '../../../utils/getCookies';
import { NextResponse } from 'next/server';

export async function GET(): Promise<NextResponse> {
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
      return NextResponse.json({ data: fetchedBookGenres }, { status: 200 });
    } else {
      return NextResponse.json({ data: [] }, { status: 200 });
    }
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
