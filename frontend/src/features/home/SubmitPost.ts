'use server';
import { apiInstance } from '../../lib/api/apiInstance';
import { PostSchema } from '../../types/PostSchema';

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

  const api = apiInstance;
  if (postImage) {
    await api.createPost(content, rating, ISBNcode, postImage);
    return {};
  } else {
    await api.createPost(content, rating, ISBNcode);
    return {};
  }
}
