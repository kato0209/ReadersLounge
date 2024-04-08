import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';
import { CreateCommentLikeReqBody } from '../../../openapi';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET(request: NextRequest): Promise<NextResponse> {
  const searchParams = request.nextUrl.searchParams;
  const commentID = searchParams.get('commentID');
  if (!commentID) {
    return NextResponse.json(
      { error: 'commentID is required' },
      { status: 500 },
    );
  }
  try {
    const req: CreateCommentLikeReqBody = {
      comment_id: Number(commentID),
    };
    const cookie = getAllCookies();
    const api = apiInstance;
    await api.createCommentLike(req, { headers: { Cookie: cookie } });
  } catch (error: unknown) {
    return Promise.reject(error);
  }
  return NextResponse.json({ status: 201 });
}
