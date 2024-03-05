import { apiInstance } from '../../lib/api/apiInstance';
import { User } from '../../openapi';

export const fetchUserData = async (): Promise<User> => {
  try {
    const api = await apiInstance;
    const res = await api.getLoginUser();
    const user: User = {
      user_id: res.data.user_id,
      name: res.data.name,
      profile_image: res.data.profile_image,
    };
    return user;
  } catch (error: unknown) {
    return Promise.reject(error);
  }
};
