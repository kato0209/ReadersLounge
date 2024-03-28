'use client';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import SubmitButton from '../../components/Button/SubmitButton';
import PortalLogo from '../../components/Logo/PortalLogo';
import { useFormState } from 'react-dom';
import { login } from './SubmitLogin';
import { State } from './SubmitLogin';

export const initialState: State = {
  error: '',
  fieldErrors: {
    email: '',
    password: '',
  },
};

export default function Login({ GoogleAuth }: { GoogleAuth: React.ReactNode }) {
  const [state, formAction] = useFormState(login, initialState);
  return (
    <Container component="main" maxWidth="xs">
      <Box
        sx={{
          marginTop: '8rem',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <PortalLogo />
        <Typography component="h1" variant="h5" sx={{ mt: 1 }}>
          Login
        </Typography>
        <form action={formAction}>
          <TextField
            required
            fullWidth
            id="email"
            label="メールアドレス"
            name="email"
            autoComplete="email"
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="パスワード"
            type="password"
            id="password"
            autoComplete="current-password"
          />
          {state.error && <span style={{ color: 'red' }}>{state.error}</span>}
          <SubmitButton content="LOGIN" />
          <Grid container justifyContent="flex-end">
            <Grid item>
              <Link href="/signup" variant="body2">
                アカウント作成はこちら
              </Link>
            </Grid>
          </Grid>
          <Typography
            component="h1"
            variant="h5"
            sx={{ mt: 2, mb: 1, textAlign: 'center' }}
          >
            Login with another provider
          </Typography>
          {GoogleAuth}
        </form>
      </Box>
    </Container>
  );
}
