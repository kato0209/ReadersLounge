import * as React from 'react';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { ReqSignupBody } from '../../openapi/models';
import { apiInstance } from '../../lib/api/apiInstance';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { AxiosError } from 'axios' ;
import { useErrorHandler } from 'react-error-boundary';
import SubmitButton from '../../components/Button/SubmitButton';
import PortalLogo from '../../components/Logo/PortalLogo';
import GoogleAuth from '../../components/OAuth/GoogleAuth';
import { useNavigate } from 'react-router-dom';
import { User } from '../../openapi';
import { useAuthUserContext } from '../../lib/auth/auth';


const SignupSchema = z.object({
    email: z.string().nonempty('メールアドレスは必須です').email('有効なメールアドレスを入力してください'),
    username: z.string().nonempty('ユーザー名は必須です'),
    password: z.string().nonempty('パスワードは必須です')
      .regex(
        /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,100}$/,
        '半角英小文字大文字数字をそれぞれ1種類以上含む8文字以上100文字以下のパスワードを設定して下さい'
      ),
    confirmationPassword: z.string().nonempty('パスワードの再入力は必須です'),
})
.refine(data => data.password === data.confirmationPassword, {
    path: ['confirmationPassword'],
    message: 'パスワードが一致しません',
});

type FormData = z.infer<typeof SignupSchema>;

export default function SignUp() {
    const { register, handleSubmit, setError, formState: { errors } } = useForm<FormData>({
        resolver: zodResolver(SignupSchema),
    });
    const errorHandler = useErrorHandler();
    const navigate = useNavigate();
    const { login } = useAuthUserContext();

    const onSubmit = async (data: FormData) => {

        const reqSignupBody: ReqSignupBody = {
            identifier: data.email,
            username: data.username,
            credential: data.password,
        };
        try {
            const api = await apiInstance;
            const res = await api.signup(reqSignupBody);
            const user: User = {
                user_id: res.data.user_id,
                name: res.data.name,
                profile_image: res.data.profile_image,
              }
            login(user);
            navigate('/');
        } catch (error: unknown) {
            if (error instanceof AxiosError) {
                if (error.response && error.response.data && error.response.data === 'email already exists') {
                    setError('email', {
                        type: 'manual',
                        message: 'このメールアドレスは既に使用されています。',
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
                <PortalLogo />
                <Typography component="h1" variant="h5" sx={{ mt: 1 }}>
                    Sign up
                </Typography>
                <Box component="form" noValidate onSubmit={handleSubmit(onSubmit)} sx={{ mt: 1 }}>
                    <Grid container spacing={2}>
                        <Grid item xs={12}>
                            <TextField
                                {...register("email")}
                                required
                                fullWidth
                                id="email"
                                label="メールアドレス"
                                name="email"
                                autoComplete="email"
                            />
                            {errors.email && <span style={{ color: 'red' }}>{errors.email.message}</span>}
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                {...register("username")}
                                required
                                fullWidth
                                name="username"
                                label="ユーザー名"
                                id="username"
                            />
                            {errors.username && <span style={{ color: 'red' }}>{errors.username.message}</span>}
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField
                                {...register("password")}
                                required
                                fullWidth
                                name="password"
                                label="パスワード"
                                type="password"
                                id="password"
                                autoComplete="new-password"
                            />
                            {errors.password && <span style={{ color: 'red' }}>{errors.password.message}</span>}
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField
                                {...register("confirmationPassword")}
                                required
                                fullWidth
                                name="confirmationPassword"
                                label="パスワード(再入力)"
                                type="password"
                                id="confirmationPassword"
                                autoComplete="new-password"
                            />
                            {errors.confirmationPassword && <span style={{ color: 'red' }}>{errors.confirmationPassword.message}</span>}
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
                    <Typography component="h1" variant="h5" sx={{ mt: 2, mb: 1 }}>
                        Sign up with another provider
                    </Typography>
                    <GoogleAuth />
                </Box>
            </Box>
        </Container>
    );
}