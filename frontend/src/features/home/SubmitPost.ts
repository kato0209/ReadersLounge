'use server';
import { apiInstance } from '../../lib/api/apiInstance';
import { PostSchema } from '../../types/PostSchema';
import { getAllCookies } from '../../utils/getCookies';

export type State = {
  error?: string;
  fieldErrors?: {
    content?: string;
    rating?: string;
    ISBNcode?: string;
    postImage?: string;
  };
};

export async function post(state: State, formData: FormData): Promise<State> {
  const validatedFields = PostSchema.safeParse({
    content: formData.get('content'),
    rating: formData.get('rating'),
    ISBNcode: formData.get('ISBNcode'),
    postImage: formData.get('postImage'),
  });

  if (!validatedFields.success) {
    throw validatedFields.error.flatten().fieldErrors;
  }

  const { content, rating, ISBNcode, postImage } = validatedFields.data;

  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    if (postImage) {
      await api.createPost(content, rating, ISBNcode, postImage, {
        headers: { Cookie: cookie },
      });
      return {};
    } else {
      await api.createPost(content, rating, ISBNcode, undefined, {
        headers: { Cookie: cookie },
      });
      return {};
    }
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
