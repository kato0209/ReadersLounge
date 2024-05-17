'use server';
import { apiInstance } from '../../lib/api/apiInstance';
import { z } from 'zod';
import { Book } from '../../openapi';
import { getAllCookies } from '../../utils/getCookies';

export type State = {
  error?: string;
  fieldErrors?: {
    keyword?: string;
    bookGenreID?: string;
  };
  fetchedBooks?: Book[];
  bookNotFound?: boolean;
};

export async function searchBook(
  state: State,
  formData: FormData,
): Promise<State> {
  const searchBookSchema = z.object({
    keyword: z.string().optional(),
    bookGenreID: z.string().optional(),
  });
  const validatedFields = searchBookSchema.safeParse({
    keyword: formData.get('keyword'),
    bookGenreID: formData.get('bookGenreID'),
  });

  if (validatedFields.success === false) {
    return {
      fieldErrors: {
        keyword: validatedFields.error.flatten().fieldErrors.keyword[0],
        bookGenreID: validatedFields.error.flatten().fieldErrors.bookGenreID[0],
      },
    };
  }

  const { keyword, bookGenreID } = validatedFields.data;

  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.fetchBookData(bookGenreID, keyword, {
      headers: { Cookie: cookie },
    });
    if (res.data && Array.isArray(res.data)) {
      if (res.data.length === 0) {
        return { bookNotFound: true };
      }
      const fetchedBooks: Book[] = res.data.map((item) => ({
        book_id: item.book_id,
        ISBNcode: item.ISBNcode,
        title: item.title,
        author: item.author,
        price: item.price,
        publisher: item.publisher,
        published_at: item.published_at,
        item_url: item.item_url,
        image: item.image,
      }));
      return { fetchedBooks: fetchedBooks };
    } else {
      return Promise.reject(new Error('Failed to fetch book data'));
    }
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
