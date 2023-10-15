import { useRoutes } from 'react-router-dom';

import { publicRoutes } from './public';
import Home from '../features/home/Home';

export const AppRoutes = () => {

    const commonRoutes = [{ path: '/', element: <Home /> }];

    //const routes = auth.user ? protectedRoutes : publicRoutes;
    const routes = publicRoutes;

    const element = useRoutes([...routes, ...commonRoutes]);

    return <>{element}</>;
};