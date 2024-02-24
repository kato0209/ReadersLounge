import * as React from 'react';
import Link from '@mui/material/Link';
import { getGoogleAuthUrl } from '../../utils/getGoogleAuthUrl';
import { FcGoogle } from 'react-icons/fc';
import { useCookies } from 'react-cookie';
import { generateRandomState } from '../../utils/generateRandomState';

export default function GoogleAuth() {
  const [_, setCookie] = useCookies(['state']);
  const [state, setState] = React.useState<string>('');

  React.useEffect(() => {
    const state = generateRandomState(10);
    setState(state);

    const now = new Date();
    const oneDay = 24 * 60 * 60 * 1000;
    const expires = new Date(now.getTime() + oneDay);
    setCookie('state', state, { path: '/', expires: expires });
  }, []);

  return (
    <Link
      href={getGoogleAuthUrl(state)}
      sx={{
        backgroundColor: '#EFEBE5',
        borderRadius: 1,
        py: '0.6rem',
        columnGap: '1rem',
        textDecoration: 'none',
        color: '#393e45',
        cursor: 'pointer',
        fontWeight: 500,
        border: '1px solid #a9a9a9',
        '&:hover': {
          color: 'inherit',
          backgroundColor: '#EAE6E0',
          border: '1px solid black',
        },
        '&:active': {
          backgroundColor: '#D5D0CA',
        },
      }}
      display="flex"
      justifyContent="center"
      alignItems="center"
    >
      <FcGoogle />
      Google
    </Link>
  );
}
