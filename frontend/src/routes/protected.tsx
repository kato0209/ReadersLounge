import Home from '../features/home/Home';
import { SearchBook } from '../features/searchBook/SearchBook';
import UserProfile  from '../features/userProfile/UserProfile';
import RoomList from '../features/chat/RoomList';

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
        path: '/user-profile/:id',
        element: <UserProfile />,
    }, 
    {
        path: '/chat-room-list',
        element: <RoomList />,
    }, 
    {
        path: '/chat-room-list/:id',
        element: <RoomList />,
    },
];
