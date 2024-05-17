import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';
import { CreatePostLikeReqBody } from '../../../openapi';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET(request: NextRequest): Promise<NextResponse> {
  const searchParams = request.nextUrl.searchParams;
  const postID = searchParams.get('postID');
  if (!postID) {
    return NextResponse.json({ error: 'postID is required' }, { status: 500 });
  }
  try {
    const req: CreatePostLikeReqBody = {
      post_id: Number(postID),
    };
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.createPostLike(req, { headers: { Cookie: cookie } });
    return NextResponse.json({ postLike: res.data });
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
