import { apiInstance } from '../../../lib/api/apiInstance';
import { getAllCookies } from '../../../utils/getCookies';
import { NextResponse } from 'next/server';
import { setJwtTokenInCookie } from '../../../lib/jwt/setJwtToken';

export async function POST(): Promise<NextResponse> {
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.logout({ headers: { Cookie: cookie } });
    if (res.status === 200) {
      setJwtTokenInCookie(res);
      return NextResponse.json({ status: 200 });
    }
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
