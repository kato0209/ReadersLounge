import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import ReadersLoungeLogo from '../../assets/images/ReadersLounge-logo-book.png';
import { DefaultApi } from '../../openapi/api';
import { Configuration } from '../../openapi';
import axios from 'axios';

export default function SignUp() {
    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);

        if(!import.meta.env.VITE_API_URL){
            console.error('環境変数BACKEND_API_URLが設定されていません');
            process.exit();
        }

        let csrfToken = '';
        axios.get(`${import.meta.env.VITE_API_URL}/csrftoken`)
            .then((results) => {
                csrfToken = results.data.csrf_token;
            })
            .catch((error) => {
                console.log(error.status);
            });
        const config = new Configuration({
            basePath: import.meta.env.VITE_API_URL,
            apiKey: csrfToken,
        });

        const apiInstance = new DefaultApi(config);
        try {
            const csrfResponse = await apiInstance.csrftoken();
            console.log(csrfResponse.data.csrf_token);
        } catch (error) {
            console.error(error);
        }
        console.log({
            email: data.get('email'),
            password: data.get('password'),
        });
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
                <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 1 }}>
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
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                required
                                fullWidth
                                name="username"
                                label="ユーザー名"
                                id="username"
                            />
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
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField
                                required
                                fullWidth
                                name="password2"
                                label="パスワード(再入力)"
                                type="password"
                                id="password2"
                                autoComplete="new-password"
                            />
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
                            <Link href="#" variant="body2">
                                アカウントをお持ちの方はこちら
                            </Link>
                        </Grid>
                    </Grid>
                </Box>
            </Box>
        </Container>
    );
}