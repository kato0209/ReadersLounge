import { apiInstance } from '../../../lib/api/apiInstance';

export async function GET() {
  try {
    const api = apiInstance;
    await api.logout();
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
