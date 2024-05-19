'use server';
import { z } from 'zod';
import { ReqLoginBody } from '../../openapi/models';
import { AxiosError } from 'axios';
import { apiInstance } from '../../lib/api/apiInstance';
import { redirect } from 'next/navigation';
import { setJwtTokenInCookie } from '../../lib/jwt/setJwtToken';
import { getAllCookies } from '../../utils/getCookies';

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

  if (validatedFields.success === false) {
    const fieldErrors = validatedFields.error.flatten().fieldErrors;
    const emailError = fieldErrors?.email ? fieldErrors.email[0] : undefined;
    const passwordError = fieldErrors?.password
      ? fieldErrors.password[0]
      : undefined;
    return {
      fieldErrors: {
        email: emailError,
        password: passwordError,
      },
    };
  }

  const { email, password } = validatedFields.data;
  const reqLoginBody: ReqLoginBody = {
    identifier: email,
    credential: password,
  };

  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.login(reqLoginBody, { headers: { Cookie: cookie } });
    setJwtTokenInCookie(res);
    redirect('/');
  } catch (error: unknown) {
    if (error instanceof AxiosError) {
      if (error.response && error.response.status === 500) {
        return { error: 'メールアドレスまたはパスワードが間違っています' };
      } else {
        return Promise.reject(error);
      }
    } else {
      return Promise.reject(error);
    }
  }
}
