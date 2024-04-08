import { apiInstance } from '../../../lib/api/apiInstance';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET() {
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    await api.logout({ headers: { Cookie: cookie } });
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
