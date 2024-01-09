import * as React from 'react';
import Sidebar from '../../components/Sidebar/Sidebar';
import { useIsMobileContext } from '../../providers/mobile/isMobile';
import UserSearchComponent from './UserSearchComponent';

export default function UserSearch() {
    const isMobile = useIsMobileContext();

  return (
    <div style={{ display: 'flex'}}>
        {!isMobile && (
            <div style={{ flex: '0 0 30%', display: 'flex' }}>
                <Sidebar />
            </div>
        )}
        <div style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
            <UserSearchComponent />
        </div>
    </div>
  );
}