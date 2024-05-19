'use server';
import { z } from 'zod';
import { ReqSignupBody } from '../../openapi/models';
import { AxiosError } from 'axios';
import { apiInstance } from '../../lib/api/apiInstance';
import { redirect } from 'next/navigation';
import { setJwtTokenInCookie } from '../../lib/jwt/setJwtToken';
import { getAllCookies } from '../../utils/getCookies';

export type State = {
  error?: string;
  fieldErrors?: {
    email?: string;
    username?: string;
    password?: string;
    confirmationPassword?: string;
  };
};

export async function signup(state: State, formData: FormData): Promise<State> {
  const SignupSchema = z
    .object({
      email: z
        .string()
        .nonempty('メールアドレスは必須です')
        .email('有効なメールアドレスを入力してください'),
      username: z.string().nonempty('ユーザー名は必須です'),
      password: z
        .string()
        .nonempty('パスワードは必須です')
        .regex(
          /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,100}$/,
          '半角英小文字大文字数字をそれぞれ1種類以上含む8文字以上100文字以下のパスワードを設定して下さい',
        ),
      confirmationPassword: z.string().nonempty('パスワードの再入力は必須です'),
    })
    .refine((data) => data.password === data.confirmationPassword, {
      path: ['confirmationPassword'],
      message: 'パスワードが一致しません',
    });

  const validatedFields = SignupSchema.safeParse({
    email: formData.get('email'),
    username: formData.get('username'),
    password: formData.get('password'),
    confirmationPassword: formData.get('confirmationPassword'),
  });

  if (validatedFields.success === false) {
    const fieldErrors = validatedFields.error.flatten().fieldErrors;
    const emailError = fieldErrors?.email ? fieldErrors.email[0] : undefined;
    const usernameError = fieldErrors?.username
      ? fieldErrors.username[0]
      : undefined;
    const passwordError = fieldErrors?.password
      ? fieldErrors.password[0]
      : undefined;
    const confirmationPasswordError = fieldErrors?.confirmationPassword
      ? fieldErrors.confirmationPassword[0]
      : undefined;
    return {
      fieldErrors: {
        email: emailError,
        username: usernameError,
        password: passwordError,
        confirmationPassword: confirmationPasswordError,
      },
    };
  }

  const { email, username, password } = validatedFields.data;
  const reqSignupBody: ReqSignupBody = {
    identifier: email,
    username: username,
    credential: password,
  };

  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.signup(reqSignupBody, {
      headers: { Cookie: cookie },
    });
    setJwtTokenInCookie(res);
    redirect('/');
  } catch (error: unknown) {
    if (error instanceof AxiosError) {
      if (
        error.response &&
        error.response.data &&
        error.response.data === 'EMAIL_ALREADY_USED'
      ) {
        return {
          fieldErrors: {
            email: 'このメールアドレスは既に使用されています。',
          },
        };
      } else {
        throw error;
      }
    } else {
      throw error;
    }
  }
}
