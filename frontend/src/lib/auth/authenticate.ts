import { cookies } from 'next/headers';

export function authenticate(): boolean {
  const jwtToken = cookies().get('jwt_token');
  if (jwtToken === undefined) {
    return false;
  } else if (jwtToken.value === '') {
    return false;
  }
  return true;
}
