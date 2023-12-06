import Home from '../features/home/Home';
import { SearchBook } from '../features/searchBook/SearchBook';
import UserProfile  from '../features/userProfile/UserProfile';

export const protectedRoutes = [
    { 
        path: '/', 
        element: <Home /> 
    },
    {
        path: '/search-book',
        element: <SearchBook />,
    },  
    {
        path: '/user-profile',
        element: <UserProfile />,
    },  
];
