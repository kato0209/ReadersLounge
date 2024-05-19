import { apiInstance } from '../../../lib/api/apiInstance';
import { User } from '../../../openapi';
import { getAllCookies } from '../../../utils/getCookies';
import { NextResponse } from 'next/server';

export async function GET(): Promise<NextResponse> {
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.getLoginUser({ headers: { Cookie: cookie } });
    const user: User = {
      user_id: res.data.user_id,
      name: res.data.name,
      profile_image: res.data.profile_image,
      profile_text: res.data.profile_text,
    };
    return NextResponse.json({ data: user }, { status: 200 });
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
