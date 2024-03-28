'use client';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import SubmitButton from '../../components/Button/SubmitButton';
import PortalLogo from '../../components/Logo/PortalLogo';
import { signup } from './SubmitSignup';
import { State } from './SubmitSignup';
import { useFormState } from 'react-dom';

export const initialState: State = {
  error: '',
  fieldErrors: {
    email: '',
    username: '',
    password: '',
    confirmationPassword: '',
  },
};

export default function Signup({
  GoogleAuth,
}: {
  GoogleAuth: React.ReactNode;
}) {
  const [state, formAction] = useFormState(signup, initialState);
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
          Sign up
        </Typography>
        <form action={formAction}>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <TextField
                required
                fullWidth
                id="email"
                label="メールアドレス"
                name="email"
                autoComplete="email"
              />
              {state.fieldErrors?.email && (
                <span style={{ color: 'red' }}>{state.fieldErrors?.email}</span>
              )}
            </Grid>
            <Grid item xs={12}>
              <TextField
                required
                fullWidth
                name="username"
                label="ユーザー名"
                id="username"
              />
              {state.fieldErrors?.username && (
                <span style={{ color: 'red' }}>
                  {state.fieldErrors?.username}
                </span>
              )}
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                name="password"
                label="パスワード"
                type="password"
                id="password"
                autoComplete="new-password"
              />
              {state.fieldErrors?.password && (
                <span style={{ color: 'red' }}>
                  {state.fieldErrors?.password}
                </span>
              )}
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                name="confirmationPassword"
                label="パスワード(再入力)"
                type="password"
                id="confirmationPassword"
                autoComplete="new-password"
              />
              {state.fieldErrors?.confirmationPassword && (
                <span style={{ color: 'red' }}>
                  {state.fieldErrors?.confirmationPassword}
                </span>
              )}
            </Grid>
          </Grid>
          <SubmitButton content="SIGN UP" />
          <Grid container justifyContent="flex-end">
            <Grid item>
              <Link href="/login" variant="body2">
                アカウントをお持ちの方はこちら
              </Link>
            </Grid>
          </Grid>
          <Typography
            component="h1"
            variant="h5"
            sx={{ mt: 2, mb: 1, textAlign: 'center' }}
          >
            Sign up with another provider
          </Typography>
          {GoogleAuth}
        </form>
      </Box>
    </Container>
  );
}
