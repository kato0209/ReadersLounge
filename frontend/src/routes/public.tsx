import Login from '../features/login/Login';
import SignUp from '../features/signup/SignUp';

export const publicRoutes = [
  {
    path: '/signup',
    element: <SignUp />,
  },
  {
    path: '/login',
    element: <Login />,
  },
];
