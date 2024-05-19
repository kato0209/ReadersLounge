import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { type NextRequest } from 'next/server';

export async function GET(request: NextRequest): Promise<NextResponse> {
  const searchParams = request.nextUrl.searchParams;
  const state = searchParams.get('state');
  if (!state) {
    return NextResponse.json({ error: 'State is required' }, { status: 500 });
  }
  cookies().set({
    name: 'state',
    value: state,
    path: '/',
    domain: process.env.COOKIE_DOMAIN,
    expires: new Date(Date.now() + 1000 * 60 * 60 * 24),
  });
  return NextResponse.json({ status: 200 });
}
