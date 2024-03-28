import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { type NextRequest } from 'next/server';

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const state = searchParams.get('state');
  if (!state) {
    return NextResponse.json({ error: 'State is required' }, { status: 500 });
  }
  cookies().set({
    name: 'state',
    value: state,
    path: '/',
  });
  return NextResponse.json({ status: 200 });
}
