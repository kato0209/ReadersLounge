import * as React from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import ErrorFallback from '../components/Error/ErrorFallback';
import { AuthProvider } from '../lib/auth/auth';
import { CookiesProvider } from 'react-cookie';
import { BrowserRouter } from 'react-router-dom';
import { MobileProvider } from './mobile/isMobile';

type AppProviderProps = {
  children: React.ReactNode;
};

export const AppProvider: React.FC<AppProviderProps> = ({ children }) => {
  return (
    <ErrorBoundary FallbackComponent={ErrorFallback}>
      <CookiesProvider>
        <AuthProvider>
          <MobileProvider>
            <BrowserRouter>{children}</BrowserRouter>
          </MobileProvider>
        </AuthProvider>
      </CookiesProvider>
    </ErrorBoundary>
  );
};
