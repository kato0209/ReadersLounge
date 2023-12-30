import * as React from 'react';
import useMediaQuery from '@mui/material/useMediaQuery';

type MobileProviderProps = {
    children: React.ReactNode;
}

const MobileContext = React.createContext<boolean>(false);

export const MobileProvider: React.FC<MobileProviderProps> = ({ children }) => {
    const isMobile = useMediaQuery('(max-width:500px)');

    return (
        <MobileContext.Provider value={isMobile}>
            {children}
        </MobileContext.Provider>
    );
};

export const useIsMobileContext = ():boolean => {
    return React.useContext<boolean>(MobileContext);
};