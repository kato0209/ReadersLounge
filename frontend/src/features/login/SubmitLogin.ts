'use server';
import { z } from 'zod';
import { ReqLoginBody } from '../../openapi/models';
import { AxiosError } from 'axios';
import { redirect } from 'next/navigation';
import { apiInstance } from '../../lib/api/apiInstance';

export type State = {
  error?: string;
  fieldErrors?: {
    email?: string;
    password?: string;
  };
};

export async function login(state: State, formData: FormData): Promise<State> {
  const LoginSchema = z.object({
    email: z.string().nonempty('メールアドレスは必須です'),
    password: z.string().nonempty('パスワードは必須です'),
  });

  const validatedFields = LoginSchema.safeParse({
    email: formData.get('email'),
    password: formData.get('password'),
  });

  if (!validatedFields.success) {
    throw validatedFields.error.flatten().fieldErrors;
  }

  const { email, password } = validatedFields.data;
  const reqLoginBody: ReqLoginBody = {
    identifier: email,
    credential: password,
  };

  try {
    const api = await apiInstance;
    await api.login(reqLoginBody);
    redirect('/');
  } catch (error: unknown) {
    if (error instanceof AxiosError) {
      if (error.response && error.response.status === 500) {
        return { error: 'メールアドレスまたはパスワードが間違っています' };
      } else {
        throw error;
      }
    } else {
      throw error;
    }
  }
}
