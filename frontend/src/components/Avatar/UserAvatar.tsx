import React from 'react';
import Avatar from '@mui/material/Avatar';
import { isValidUrl } from '../../utils/isValidUrl';

function UserAvatar(props: { image: string }) {
    return (
        <Avatar 
            src={isValidUrl(props.image) ? props.image : `data:image/png;base64,${props.image}` }>                  
        </Avatar>
    );
}

export default UserAvatar;
