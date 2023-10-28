import React from 'react';
import { Navigate } from 'react-router-dom';
import { useCookies } from 'react-cookie'; 

type AuthRouteProps = {
    children: React.ReactNode;
};

const AuthProvider: React.FC<AuthRouteProps> = ({ children }) => {
    const [cookies] = useCookies(['jwt_token']);
    console.log(cookies);
    const isAuthenticated = Boolean(cookies.jwt_token);
    console.log(isAuthenticated);

    return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
};

export default AuthProvider;