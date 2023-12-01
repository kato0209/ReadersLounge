import Home from '../features/home/Home';
import SearchBook from '../features/searchBook/SearchBook';

export const protectedRoutes = [
    { 
        path: '/', 
        element: <Home /> 
    },
    {
        path: '/search-book',
        element: <SearchBook />,
    },  
];
