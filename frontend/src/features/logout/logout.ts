import * as React from 'react';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { useNavigate } from 'react-router-dom';
import { useAuthUserContext } from '../../lib/auth/auth';

export default function useLogout() {
  const navigate = useNavigate();
  const errorHandler = useErrorHandler();
  const { logout } = useAuthUserContext();

  return React.useCallback(() => {
    async function handleLogout() {
      try {
        logout();
        const api = await apiInstance;
        await api.logout();
        navigate('/login');
      } catch (error: unknown) {
        errorHandler(error);
      }
    }
    handleLogout();
  }, []);
}
