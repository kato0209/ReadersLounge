'use server';
import { z } from 'zod';
import { apiInstance } from '../../lib/api/apiInstance';
import { ReqCreateCommentBody } from '../../openapi/models';
import { getAllCookies } from '../../utils/getCookies';
import { redirect } from 'next/navigation';

export type State = {
  error?: string;
  fieldErrors?: {
    content?: string;
    postID?: string;
  };
};

export async function createComment(
  state: State,
  formData: FormData,
): Promise<State> {
  const CommentSchema = z.object({
    content: z.string().nonempty('投稿内容は必須です').max(255, {
      message: '投稿内容は255文字以内で入力してください',
    }),
    postID: z.number().positive(),
  });

  const validatedFields = CommentSchema.safeParse({
    content: formData.get('content'),
    postID: Number(formData.get('postID')),
  });

  if (validatedFields.success === false) {
    return {
      fieldErrors: {
        content: validatedFields.error.flatten().fieldErrors.content
          ? validatedFields.error.flatten().fieldErrors.content[0]
          : undefined,
        postID: validatedFields.error.flatten().fieldErrors.postID
          ? validatedFields.error.flatten().fieldErrors.postID[0]
          : undefined,
      },
    };
  }

  const { content, postID } = validatedFields.data;
  const req: ReqCreateCommentBody = {
    content: content,
    post_id: postID,
  };

  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    await api.createComment(req, { headers: { Cookie: cookie } });
    redirect('/post/' + postID);
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
