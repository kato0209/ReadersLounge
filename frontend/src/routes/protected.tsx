import Home from '../features/home/Home';
import { SearchBook } from '../features/searchBook/SearchBook';
import UserProfile  from '../features/userProfile/UserProfile';
import RoomList from '../features/chat/RoomList';
import Room from '../features/chat/Room';

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
    {
        path: '/chat-room-list',
        element: <RoomList />,
    },
    {
        path: '/chat-room',
        element: <Room />,
    },
];
