import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET(request: NextRequest): Promise<NextResponse> {
  const searchParams = request.nextUrl.searchParams;
  const selectedCommentID = searchParams.get('selectedCommentId');
  if (!selectedCommentID) {
    return NextResponse.json(
      { error: 'selectedPostId is required' },
      { status: 500 },
    );
  }
  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    await api.deleteComment(Number(selectedCommentID), {
      headers: { Cookie: cookie },
    });
    return NextResponse.json({ status: 204 });
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
