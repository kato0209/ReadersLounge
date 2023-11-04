import React from 'react';
import { User } from '../../openapi';
import { useCookies } from 'react-cookie';

type AuthUserContextType = {
    isAuthenticated: boolean;
    user: User | null;
    login: (user:User, callback:() => void) => void;
    logout: () => void;
};

const AuthUserContext = React.createContext<AuthUserContextType>({} as AuthUserContextType);

type AuthRouteProps = {
    children: React.ReactNode;
};

export const AuthProvider: React.FC<AuthRouteProps> = ({ children }) => {
    const [cookies, setCookie, removeCookie] = useCookies(['isAuthenticated']);
    const [user, setUser] = React.useState<User | null>(null);
    const isAuthenticated = Boolean(cookies.isAuthenticated);

    const login = (newUser: User, callback: () => void) => {
        setCookie('isAuthenticated', 'true', { path: '/', secure: true, sameSite: 'strict' });
        setUser(newUser);
        callback();
    }

    const logout = () => {
        removeCookie('isAuthenticated', { path: '/' });
        setUser(null);
    }

    const value:AuthUserContextType = { isAuthenticated, user, login, logout };
    return (
    <AuthUserContext.Provider value={value}>
        {children}
    </AuthUserContext.Provider>
    );
};

export const useAuthUserContext = ():AuthUserContextType => {
    return React.useContext<AuthUserContextType>(AuthUserContext);
};