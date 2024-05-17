import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET(request: NextRequest): Promise<NextResponse> {
  const searchParams = request.nextUrl.searchParams;
  const postID = searchParams.get('postID');
  if (!postID) {
    return NextResponse.json({ error: 'postID is required' }, { status: 500 });
  }
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.deletePostLike(Number(postID), {
      headers: { Cookie: cookie },
    });
    return NextResponse.json({ postLike: res.data });
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
