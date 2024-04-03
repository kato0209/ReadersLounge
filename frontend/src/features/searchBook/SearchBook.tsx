import * as React from 'react';
import { SearchBookComponent } from './SearchBookComponent';
import { z } from 'zod';
import { PostSchema } from '../../types/PostSchema';
import { BookGenreNode } from '../../openapi/models';
import { apiInstance } from '../../lib/api/apiInstance';

type PostFormData = z.infer<typeof PostSchema>;
type SearchBookProps = {
  formData?: PostFormData;
};

export const SearchBook: React.FC<SearchBookProps> = async ({ formData }) => {
  const fetchBookGenres = async (): Promise<BookGenreNode[]> => {
    try {
      const api = apiInstance;
      const res = await api.getBooksGenres();

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
        return Promise.reject(new Error('Failed to fetch book genres'));
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const bookGenreNodes = await fetchBookGenres();

  return (
    <SearchBookComponent formData={formData} bookGenreNodes={bookGenreNodes} />
  );
};
