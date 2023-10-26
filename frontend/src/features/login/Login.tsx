import * as React from 'react';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import LogoTitle from '../../components/Logo/LogoTitle';
import SubmitButton from '../../components/Button/SubmitButton';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { useErrorHandler } from 'react-error-boundary';
import { ReqLoginBody } from '../../openapi/models';
import { apiInstance } from '../../lib/api/apiInstance';
import { AxiosError } from 'axios';


const LoginSchema = z.object({
  email: z.string().nonempty('メールアドレスは必須です'),
  password: z.string().nonempty('パスワードは必須です')
});

type FormData = z.infer<typeof LoginSchema>;

export default function Login() {
  const { register, handleSubmit, setError, formState: { errors } } = useForm<FormData>({
    resolver: zodResolver(LoginSchema),
});

const errorHandler = useErrorHandler();

const onSubmit = async (data: FormData) => {

  const reqLoginBody: ReqLoginBody = {
    identifier: data.email,
    credential: data.password,
  };

  try {
    const api = await apiInstance;
    const res = await api.login(reqLoginBody);
    console.log(res);
  } catch (error: unknown) {
    if (error instanceof AxiosError) {
      if (error.response && error.response.status === 500) {
          setError('password', {
              type: 'manual',
              message: 'メールアドレスまたはパスワードが間違っています',
          });
      } else {
          errorHandler(error);
      }
    } else {
        errorHandler(error);
    }
  }
  

};

  return (
      <Container component="main" maxWidth="xs">
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <LogoTitle />
          <Typography component="h1" variant="h5" sx={{ mt: 1 }}>
            Login
          </Typography>
          <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 1 }}>
            <TextField
              {...register("email")}
              required
              fullWidth
              id="email"
              label="メールアドレス"
              name="email"
              autoComplete="email"
              autoFocus
            />
            {errors.email && <span style={{ color: 'red' }}>{errors.email.message}</span>}
            <TextField
              {...register("password")}
              margin="normal"
              required
              fullWidth
              name="password"
              label="パスワード"
              type="password"
              id="password"
              autoComplete="current-password"
            />
            {errors.password && <span style={{ color: 'red' }}>{errors.password.message}</span>}
            <SubmitButton content="LOGIN" />
            <Grid container justifyContent="flex-end">
                <Grid item>
                    <Link href="/signup" variant="body2">
                        アカウント作成はこちら
                    </Link>
                </Grid>
            </Grid>
          </Box>
        </Box>
      </Container>
  );
}