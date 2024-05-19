'use client';
import * as React from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import ErrorFallback from '../components/Error/ErrorFallback';

type AppProviderProps = {
  children: React.ReactNode;
};

export const AppProvider: React.FC<AppProviderProps> = ({ children }) => {
  return (
    <ErrorBoundary FallbackComponent={ErrorFallback}>{children}</ErrorBoundary>
  );
};
