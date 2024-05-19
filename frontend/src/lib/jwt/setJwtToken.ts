import { AxiosResponse } from 'axios';
import { cookies } from 'next/headers';

type ParsedCookie = {
  name: string;
  value: string;
  path?: string;
  domain?: string;
  expires?: Date;
  httpOnly?: boolean;
};

function parseCookieString(cookieString: string): ParsedCookie {
  const parts = cookieString.split(';').map((part) => part.trim());
  const [nameValue, ...attributes] = parts;
  const [name, value] = nameValue.split('=');

  const parsedCookie: ParsedCookie = {
    name,
    value,
  };

  attributes.forEach((attribute) => {
    const [key, val] = attribute.split('=');
    switch (key.toLowerCase()) {
      case 'path':
        parsedCookie.path = val;
        break;
      case 'domain':
        parsedCookie.domain = val;
        break;
      case 'expires':
        parsedCookie.expires = new Date(val);
        break;
      case 'httponly':
        parsedCookie.httpOnly = true;
        break;
    }
  });

  return parsedCookie;
}

export function setJwtTokenInCookie(res: AxiosResponse) {
  if (res.headers['set-cookie']) {
    const setCookieHeader = res.headers['set-cookie'];
    setCookieHeader.forEach((cookieString) => {
      const parsedCookie = parseCookieString(cookieString);
      if (parsedCookie.name === 'jwt_token') {
        cookies().set({
          name: parsedCookie.name,
          value: parsedCookie.value,
          httpOnly: parsedCookie.httpOnly,
          path: parsedCookie.path,
          domain: parsedCookie.domain,
          expires: parsedCookie.expires,
        });
      }
    });
  }
}
