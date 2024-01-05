import * as React from 'react';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import Avatar from '@mui/material/Avatar';
import { isValidUrl } from '../../utils/isValidUrl';
import { useParams } from 'react-router-dom';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { User } from '../../openapi';

export default function UserProfile() {

  const { id } = useParams<{ id: string }>();
  const idNumber = id ? parseInt(id, 10) : 0;
  const [user, setUser] = React.useState<User | null>(null);
  const errorHandler = useErrorHandler();

  React.useEffect(() => {
    const fetchUser = async () => {
    
        try {
            const api = await apiInstance;
            const res = await api.getUser(idNumber);
            
            if (res.data) {
              const targetUser: User = {
                user_id: res.data.user_id,
                name: res.data.name,
                profile_image: res.data.profile_image,
              };
                setUser(targetUser);
            }
        } catch (error: unknown) {
            errorHandler(error);
        }
            
    };
    fetchUser();
}, []);

  return (
    <Container component="main" maxWidth="xs">
        <Box
          sx={{
            marginTop: '8rem',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Box 
            component="img"
            src={isValidUrl(user?.profile_image) ? user?.profile_image : `data:image/png;base64,${user?.profile_image}` }
            sx={{ width: 100, height: 100, borderRadius: '50%' }}  
          >                  
          </Box>
        </Box> 
    </Container>
  );
}