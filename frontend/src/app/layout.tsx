import type { ReactNode } from 'react';
import { StrictMode } from 'react';
import { AppProvider } from '../providers/app';
import '../styles/globals.css';

export const metadata = {
  title: 'ReadersLounge',
};

export default function HomeLayout({ children }: { children: ReactNode }) {
  return (
    <html>
      <body>
        <StrictMode>
          <AppProvider>{children}</AppProvider>
        </StrictMode>
      </body>
    </html>
  );
}
