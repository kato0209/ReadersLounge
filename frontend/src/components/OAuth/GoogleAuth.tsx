import Link from '@mui/material/Link';
import { getGoogleAuthUrl } from '../../utils/getGoogleAuthUrl';
import { FcGoogle } from 'react-icons/fc';
import { generateRandomState } from '../../utils/generateRandomState';
import { apiInstance } from '../../lib/api/apiInstance';

export default function GoogleAuth() {
  async function setState(state: string) {
    const api = await apiInstance;
    await api.setState(state);
  }
  const state = generateRandomState(10);
  setState(state);
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
