'use client';
import * as React from 'react';
import { Connection } from '../../openapi/';
import { Box, Typography } from '@mui/material';
import UserAvatar from '../../components/Avatar/UserAvatar';
import { useRouter } from 'next/navigation';

type ConnectionListProps = {
  connections: Connection[];
};

export const ConnectionList: React.FC<ConnectionListProps> = ({
  connections,
}) => {
  const router = useRouter();
  return (
    <Box sx={{ marginTop: '2rem' }}>
      {connections.map((connection) => (
        <Box
          key={connection.connection_id}
          onClick={() =>
            router.push(`/user-profile/${connection.target_user_id}`)
          }
          sx={{
            display: 'flex',
            alignItems: 'center',
            padding: '0.5rem',
            cursor: 'pointer',
            '&:hover': {
              color: 'inherit',
              backgroundColor: '#EAE6E0',
            },
          }}
        >
          <Box sx={{ marginRight: '1rem' }}>
            <UserAvatar
              image={connection.target_user_profile_image}
              userID={connection.target_user_id}
            />
          </Box>
          <Box>
            <Typography variant="body1">
              {connection.target_user_name}
            </Typography>
          </Box>
        </Box>
      ))}
    </Box>
  );
};
