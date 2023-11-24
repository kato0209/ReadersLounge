import React from 'react';
import { User } from '../../openapi';

type AuthUserContextType = {
    isAuthenticated: boolean;
    user: User;
    login: (user:User) => void;
    logout: () => void;
};

const defaultUser: User = {
    user_id: 0,
    name: '',
    profile_image: '',
};

const AuthUserContext = React.createContext<AuthUserContextType>({} as AuthUserContextType);

type AuthRouteProps = {
    children: React.ReactNode;
};

export const AuthProvider: React.FC<AuthRouteProps> = ({ children }) => {
    const [user, setUser] = React.useState<User>(defaultUser);
    const [isAuthenticated, setIsAuthenticated] = React.useState<boolean>(false);

    const login = (newUser: User) => {
        setIsAuthenticated(true);
        setUser(newUser);
    }

    const logout = () => {
        setIsAuthenticated(false);
        setUser(defaultUser);
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