import type { ReactNode } from 'react';

export const metadata = {
  title: 'ReadersLounge',
};

export default function HomeLayout({ children }: { children: ReactNode }) {
  return (
    <html>
      <body>{children}</body>
    </html>
  );
}
