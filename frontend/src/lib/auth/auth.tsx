import React from 'react';
import { User } from '../../openapi';

type AuthUserContextType = {
    isAuthenticated: boolean;
    user: User | null;
    login: (user:User, callback:() => void) => void;
    logout: (callback:() => void) => void;
};

const AuthUserContext = React.createContext<AuthUserContextType>({} as AuthUserContextType);

type AuthRouteProps = {
    children: React.ReactNode;
};

export const AuthProvider: React.FC<AuthRouteProps> = ({ children }) => {
    const [user, setUser] = React.useState<User | null>(null);
    const [isAuthenticated, setIsAuthenticated] = React.useState<boolean>(false);

    const login = (newUser: User, callback: () => void) => {
        setUser(newUser);
        setIsAuthenticated(true);
        callback();
    }

    const logout = (callback: () => void) => {
        setUser(null);
        setIsAuthenticated(false);
        callback();
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