import Login from '../features/login/Login';
import SignUp from '../features/signup/Signup';

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
