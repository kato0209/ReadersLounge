'use server';
import { apiInstance } from '../../lib/api/apiInstance';
import { PostSchema } from '../../types/PostSchema';
import { getAllCookies } from '../../utils/getCookies';
import { redirect } from 'next/navigation';

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
    rating: Number(formData.get('rating')),
    ISBNcode: formData.get('ISBNcode'),
    postImage: formData.get('postImage') || undefined,
  });

  if (validatedFields.success === false) {
    return {
      fieldErrors: {
        content: validatedFields.error.flatten().fieldErrors.content
          ? validatedFields.error.flatten().fieldErrors.content[0]
          : undefined,
        rating: validatedFields.error.flatten().fieldErrors.rating
          ? validatedFields.error.flatten().fieldErrors.rating[0]
          : undefined,
        ISBNcode: validatedFields.error.flatten().fieldErrors.ISBNcode
          ? validatedFields.error.flatten().fieldErrors.ISBNcode[0]
          : undefined,
        postImage: validatedFields.error.flatten().fieldErrors.postImage
          ? validatedFields.error.flatten().fieldErrors.postImage[0]
          : undefined,
      },
    };
  }

  const { content, rating, ISBNcode, postImage } = validatedFields.data;

  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    if (postImage) {
      await api.createPost(content, rating, ISBNcode, postImage, {
        headers: { Cookie: cookie },
      });
    } else {
      await api.createPost(content, rating, ISBNcode, undefined, {
        headers: { Cookie: cookie },
      });
    }
    redirect('/home');
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
