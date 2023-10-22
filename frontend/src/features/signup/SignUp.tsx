import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import ReadersLoungeLogo from '../../assets/images/ReadersLounge-logo-book.png';
import { ReqSignupBody } from '../../openapi/models';
import { apiInstance } from '../../lib/api/apiInstance';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';

const SignupSchema = z.object({
    email: z.string().nonempty('メールアドレスは必須です').email('有効なメールアドレスを入力してください'),
    username: z.string().nonempty('ユーザー名は必須です'),
    password: z.string()
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
    const { register, handleSubmit, formState: { errors } } = useForm<FormData>({
        resolver: zodResolver(SignupSchema),
    });

    const onSubmit = async (data: FormData) => {

        const reqSignupBody: ReqSignupBody = {
            identifier: data.email,
            username: data.username,
            credential: data.password,
        };
        try {
            const api = await apiInstance;
            const res = await api.signup(reqSignupBody);
            console.log(res.data.user_id);
        } catch (error) {
            console.error(error);
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
                <img src={ReadersLoungeLogo}/>
                <Typography component="h1" variant="h3" style={{ color: '#FF7E73' }}>
                    ReadersLounge
                </Typography>
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
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2, backgroundColor: '#FF7E73', '&:hover': { backgroundColor: '#E56A67' } }}
                    >
                        Sign Up
                    </Button>
                    <Grid container justifyContent="flex-end">
                        <Grid item>
                            <Link href="../login" variant="body2">
                                アカウントをお持ちの方はこちら
                            </Link>
                        </Grid>
                    </Grid>
                </Box>
            </Box>
        </Container>
    );
}