import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET(request: NextRequest): Promise<NextResponse> {
  const searchParams = request.nextUrl.searchParams;
  const commentID = searchParams.get('commentID');
  if (!commentID) {
    return NextResponse.json({ error: 'postID is required' }, { status: 500 });
  }
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    await api.deleteCommentLike(Number(commentID), {
      headers: { Cookie: cookie },
    });
  } catch (error: unknown) {
    return Promise.reject(error);
  }
  return NextResponse.json({ status: 204 });
}
