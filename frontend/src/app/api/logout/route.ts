import { apiInstance } from '../../../lib/api/apiInstance';
import { getAllCookies } from '../../../utils/getCookies';
import { NextResponse } from 'next/server';

export async function GET(): Promise<NextResponse> {
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    await api.logout({ headers: { Cookie: cookie } });
    return NextResponse.json({ status: 200 });
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
