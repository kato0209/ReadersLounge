import type { ReactNode } from 'react';
import AppHeader from '../../components/Header/AppHeader';

export default function SignupLayout({ children }: { children: ReactNode }) {
  return (
    <>
      <AppHeader />
      {children}
    </>
  );
}
