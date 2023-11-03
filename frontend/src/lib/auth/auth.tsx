import React from 'react';
import { User } from '../../openapi';

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
    const localIsAuthenticated = Boolean(localStorage.getItem('isAuthenticated'));
    const [user, setUser] = React.useState<User | null>(null);
    const [isAuthenticated, setIsAuthenticated] = React.useState<boolean>(localIsAuthenticated);

    const login = (newUser: User, callback: () => void) => {
        localStorage.setItem('isAuthenticated', 'true');
        setUser(newUser);
        setIsAuthenticated(true);
        console.log(isAuthenticated);
        callback();
    }

    const logout = () => {
        console.log('logout');
        localStorage.setItem('isAuthenticated', 'false');
        setUser(null);
        setIsAuthenticated(false);
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