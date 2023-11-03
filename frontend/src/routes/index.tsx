import { publicRoutes } from './public';
import { useRoutes } from 'react-router-dom';
import { protectedRoutes } from './protected';
import PageNotFound from '../components/Error/PageNotFound';
import { Navigate } from 'react-router-dom';
import { useAuthUserContext } from '../lib/auth/auth';

export const AppRoutes = () => {

    const { isAuthenticated } = useAuthUserContext();
    console.log('AppRoutes');
    console.log(isAuthenticated);

    const element = useRoutes([
        ...publicRoutes,
        ...protectedRoutes.map(route => ({
            ...route,
            element: isAuthenticated ? route.element : <Navigate to="/login" />
        })),
        { path: '*', element: <PageNotFound /> }
    ]);

    return <>{element}</>;
};