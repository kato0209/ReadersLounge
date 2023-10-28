import { publicRoutes } from './public';
import { useRoutes } from 'react-router-dom';
import { protectedRoutes } from './protected';
import PageNotFound from '../components/Error/PageNotFound';
import AuthProvider from '../lib/auth/auth';

export const AppRoutes = () => {

    const element = useRoutes([
        ...publicRoutes,
        ...protectedRoutes.map(route => ({
            ...route,
            element: (
                <AuthProvider>
                    {route.element}
                </AuthProvider>
            )
        })),
        { path: '*', element: <PageNotFound /> }
    ]);

    return <>{element}</>;
};