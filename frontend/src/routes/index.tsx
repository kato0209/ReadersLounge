import * as React from 'react';
import { publicRoutes } from './public';
import { useRoutes } from 'react-router-dom';
import { protectedRoutes } from './protected';
import PageNotFound from '../components/Error/PageNotFound';
import { Navigate, useLocation  } from 'react-router-dom';
import { useAuthUserContext } from '../lib/auth/auth';
import AppHeader from '../components/Header/AppHeader';
import Box from '@mui/material/Box';
import { apiInstance } from '../lib/api/apiInstance';
import { User } from '../openapi';
import { AxiosError } from 'axios';
import { useErrorHandler } from 'react-error-boundary';

export const AppRoutes = () => {

    const { isAuthenticated, login } = useAuthUserContext();
    const [isAuthChecked, setIsAuthChecked] = React.useState(false);
    const location = useLocation();
    const errorHandler = useErrorHandler();

    const fetchUserData = async () => {
        try {
            const api = await apiInstance;
            const res = await api.getLoginUser();
            const user: User = {
                user_id: res.data.user_id,
                name: res.data.name,
                profile_image: res.data.profile_image,
            };
            login(user);
        } catch (error: unknown) {
            if (error instanceof AxiosError) {
                if (error.response && error.response.status && error.response.status ===  401 ) {
                    //console.log("UnAuthorized")
                } else {
                    errorHandler(error);
                }
            } else {
                errorHandler(error);
            }
        } finally {
            setIsAuthChecked(true);
        }
    };

    React.useEffect(() => {
        const isProtectedRoute = protectedRoutes.some(route => {
            const matchPattern = new RegExp(`^${route.path.replace(/:\w+/g, '\\w+')}$`);
            return matchPattern.test(location.pathname);
        });
        if (isProtectedRoute){
            if (!isAuthenticated && !isAuthChecked) {
                fetchUserData();
            }
        } else {
            if (!isAuthChecked) {
                setIsAuthChecked(true);
            }
        }
    }, []);

    const routes = useRoutes([
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

    if (!isAuthChecked) {
        return <div>Loading...</div>;
    }

    return <>{routes}</>;
};
