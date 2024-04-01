'use client';
import Link from '@mui/material/Link';
import { getGoogleAuthUrl } from '../../lib/auth/oauth/getGoogleAuthUrl';
import { FcGoogle } from 'react-icons/fc';
import axios from 'axios';
import { generateRandomState } from '../../utils/generateRandomState';
import { useEffect, useState } from 'react';
import { useErrorHandler } from 'react-error-boundary';

export default function GoogleAuth() {
  const errorHandler = useErrorHandler();
  async function setStateToCookie(state: string) {
    try {
      await axios.get(`/api/set-state?state=${state}`);
    } catch (error: unknown) {
      errorHandler(error);
    }
  }
  const [state, setState] = useState<string>('');
  useEffect(() => {
    const state = generateRandomState(10);
    setStateToCookie(state);
    setState(state);
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
