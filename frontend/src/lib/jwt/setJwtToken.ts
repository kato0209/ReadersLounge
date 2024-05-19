import { AxiosResponse } from 'axios';
import { cookies } from 'next/headers';

export function setJwtTokenInCookie(res: AxiosResponse) {
  if (res.headers['set-cookie']) {
    const setCookieHeader = res.headers['set-cookie'];
    const resCookies = setCookieHeader.map((cookie) => cookie.split(';')[0]);
    const jwtTokenCookie = resCookies.find((cookie) =>
      cookie.startsWith('jwt_token='),
    );
    const jwtToken = jwtTokenCookie?.split('=')[1];
    cookies().set({
      name: 'jwt_token',
      value: jwtToken as string,
      httpOnly: true,
      path: '/',
    });
  }
}
