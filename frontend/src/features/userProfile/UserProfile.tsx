import * as React from 'react';
import Sidebar from '../../components/Sidebar/Sidebar';
import UserProfileMain from './UserProfileMain';
import { useIsMobileContext } from '../../providers/mobile/isMobile';
import { Connection } from '../../openapi/';

export default function UserProfile() {
    const isMobile = useIsMobileContext();
    
  return (
    <div style={{ display: 'flex'}}>
        {!isMobile && (
            <div style={{ flex: '0 0 30%', display: 'flex' }}>
                <Sidebar />
            </div>
        )}
        <div style={{ flex: 1, overflowX: 'hidden' }}>
            <UserProfileMain />
        </div>
    </div>
  );
}