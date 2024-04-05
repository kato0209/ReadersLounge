import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const postID = searchParams.get('postID');
  if (!postID) {
    return NextResponse.json({ error: 'postID is required' }, { status: 500 });
  }
  try {
    const api = apiInstance;
    await api.deletePostLike(Number(postID));
  } catch (error: unknown) {
    return Promise.reject(error);
  }
  return NextResponse.json({ status: 200 });
}
