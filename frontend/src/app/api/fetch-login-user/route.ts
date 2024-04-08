import { apiInstance } from '../../../lib/api/apiInstance';
import { User } from '../../../openapi';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET(): Promise<User> {
  console.log('GET /api/fetch-login-user');
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.getLoginUser({ headers: { Cookie: cookie } });
    console.log(res);
    const user: User = {
      user_id: res.data.user_id,
      name: res.data.name,
      profile_image: res.data.profile_image,
    };
    return user;
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
