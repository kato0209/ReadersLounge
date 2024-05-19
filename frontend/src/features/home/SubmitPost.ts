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
    const fieldErrors = validatedFields.error.flatten().fieldErrors;
    const contentError = fieldErrors?.content
      ? fieldErrors.content[0]
      : undefined;
    const ratingError = fieldErrors?.rating ? fieldErrors.rating[0] : undefined;
    const ISBNcodeError = fieldErrors?.ISBNcode
      ? fieldErrors.ISBNcode[0]
      : undefined;
    const postImageError = fieldErrors?.postImage
      ? fieldErrors.postImage[0]
      : undefined;

    return {
      fieldErrors: {
        content: contentError,
        rating: ratingError,
        ISBNcode: ISBNcodeError,
        postImage: postImageError,
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
