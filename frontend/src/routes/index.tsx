import { publicRoutes } from './public';
import { useRoutes } from 'react-router-dom';
import { protectedRoutes } from './protected';
import PageNotFound from '../components/Error/PageNotFound';
import { Navigate } from 'react-router-dom';
import { useAuthUserContext } from '../lib/auth/auth';
import AppHeader from '../components/Header/AppHeader';

export const AppRoutes = () => {

    const { isAuthenticated } = useAuthUserContext();

    const element = useRoutes([
        ...publicRoutes,
        ...protectedRoutes.map(route => ({
            ...route,
            element: isAuthenticated ? (
                <>
                    <AppHeader />
                    {route.element}
                </>
            ) : <Navigate to="/login" />
        })),
        { path: '*', element: <PageNotFound /> }
    ]);

    return <>{element}</>;
};