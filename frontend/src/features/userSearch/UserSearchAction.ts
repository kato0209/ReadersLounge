'use server';
import { z } from 'zod';
import { apiInstance } from '../../lib/api/apiInstance';
import { User } from '../../openapi';
import { getAllCookies } from '../../utils/getCookies';

export type State = {
  error?: string;
  fieldErrors?: {
    keyword?: string;
  };
  users?: User[];
  userNotFound?: boolean;
};

export async function searchUser(
  state: State,
  formData: FormData,
): Promise<State> {
  const searchUserSchema = z.object({
    keyword: z.string().nonempty(),
  });

  const validatedFields = searchUserSchema.safeParse({
    keyword: formData.get('keyword'),
  });

  if (validatedFields.success === false) {
    console.log(validatedFields.error.flatten().fieldErrors);
    return {
      fieldErrors: {
        keyword: validatedFields.error.flatten().fieldErrors.keyword[0],
      },
    };
  }

  const { keyword } = validatedFields.data;

  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.searchUser(keyword, { headers: { Cookie: cookie } });
    if (res.data && Array.isArray(res.data)) {
      if (res.data.length === 0) {
        return { userNotFound: true, users: [] };
      }
      const SearchedUsers: User[] = res.data.map((item) => ({
        user_id: item.user_id,
        name: item.name,
        profile_image: item.profile_image,
      }));
      return { users: SearchedUsers, userNotFound: false };
    } else {
      return Promise.reject('failed to search user');
    }
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
