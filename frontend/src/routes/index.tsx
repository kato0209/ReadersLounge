import * as React from 'react';
import { publicRoutes } from './public';
import { useRoutes } from 'react-router-dom';
import { protectedRoutes } from './protected';
import PageNotFound from '../components/Error/PageNotFound';
import { Navigate } from 'react-router-dom';
import { useAuthUserContext } from '../lib/auth/auth';
import AppHeader from '../components/Header/AppHeader';
import Box from '@mui/material/Box';
import { apiInstance } from '../lib/api/apiInstance';
import { User } from '../openapi';

export const AppRoutes = () => {

    const fetchUserData = async () => {
        try {
            const api = await apiInstance;
            const res = await api.user();
            const user: User = {
            user_id: res.data.user_id,
            name: res.data.name,
            profile_image: res.data.profile_image,
            }
            console.log(user)
        } catch (error) {
            console.log("not authenticated")
        }
    };

    React.useEffect(() => {
        fetchUserData();
    }, []);
    const { isAuthenticated } = useAuthUserContext();

    const element = useRoutes([
        ...publicRoutes,
        ...protectedRoutes.map(route => ({
            ...route,
            element: isAuthenticated ? (
                <Box sx={{ pt: '3rem' }}>
                    <AppHeader />
                    {route.element}
                </Box>
            ) : <Navigate to="/login" />
        })),
        { path: '*', element: <PageNotFound /> }
    ]);

    return <>{element}</>;
};