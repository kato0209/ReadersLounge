import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
  console.log('middleware');
  const response = NextResponse.next();
  if (request.cookies.has('_csrf')) {
    const csrf = request.cookies.get('_csrf')?.value as string;
    response.cookies.set({
      name: '_csrf',
      value: csrf,
      path: '/',
    });
  }

  if (request.cookies.has('jwt_token')) {
    const jwt = request.cookies.get('jwt_token')?.value as string;
    response.cookies.set({
      name: 'jwt_token',
      value: jwt,
      path: '/',
    });
  }

  return response;
}
