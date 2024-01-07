import * as React from 'react';
import Container from '@mui/material/Container';
import { isValidUrl } from '../../utils/isValidUrl';
import { useParams } from 'react-router-dom';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { User } from '../../openapi';
import { CreateConnectionRequest } from '../../openapi/';
import { Box, Card, CardContent, Typography, Button, Avatar, Stack, CardMedia } from '@mui/material';
import UserHeaderImage from '../../assets/images/UserProfileHeader.jpg';
import { useAuthUserContext } from '../../lib/auth/auth';
import { CreateChatRoomRequest } from '../../openapi';
import { useNavigate } from 'react-router-dom';
import { EditProfile } from './EditProfile';

export default function UserProfileMain() {

  const { id } = useParams<{ id: string }>();
  const idNumber = id ? parseInt(id, 10) : 0;
  const { user: loginUser } = useAuthUserContext();
  const [user, setUser] = React.useState<User | null>(null);
  const errorHandler = useErrorHandler();
  const navigate = useNavigate();

  const fetchUser = async () => {
    
    try {
        const api = await apiInstance;
        const res = await api.getUser(idNumber);
        if (res.data) {
          const targetUser: User = {
            user_id: res.data.user_id,
            name: res.data.name,
            profile_image: res.data.profile_image,
            profile_text: res.data.profile_text,
          };
            setUser(targetUser);
        }
    } catch (error: unknown) {
        errorHandler(error);
    }
        
};
  React.useEffect(() => {
    fetchUser();
  }, []);

  const handleMessageClick = async () => {
    try {
      const req: CreateChatRoomRequest = {
        chat_partner_id: idNumber,
      }
      const api = await apiInstance;
      const res = await api.createChatRoom(req);
      if (res.status === 201) {
        navigate(`/chat-room-list/${res.data}`);
      }
    } catch (error: unknown) {
        errorHandler(error);
    }
  };

  const handleFollowClick = async () => {
    try {
      const req: CreateConnectionRequest = {
        target_user_id: idNumber,
      }
      const api = await apiInstance;
      const res = await api.createConnection(req);
      if (res.status === 201) {
        console.log(res.data);
      }
    } catch (error: unknown) {
        errorHandler(error);
    }
  }


  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', p: 2 }}>
      <Card sx={{ maxWidth: 500, width: '100%', mb: 2 }}>
        <CardMedia
          component="img"
          height="200"
          image={UserHeaderImage}
          alt="Cover image"
        />
        <Box sx={{ display: 'flex', justifyContent: 'center', mt: -8, position: "relative" }}>
          <Avatar
            sx={{ width: 128, height: 128, border: '4px solid', borderColor: 'background.paper'}}
            src={isValidUrl(user?.profile_image) ? user?.profile_image : `data:image/png;base64,${user?.profile_image}` }
          />
          {loginUser?.user_id === idNumber && (
            <Box sx={{display: "flex", flexDirection: "column", position: "absolute", top: "50%", right: "0.3rem"}}>
              {user && (
                <EditProfile user={user} fetchUser={fetchUser}/>
              )}
            </Box>
          )}
        </Box>
        <CardContent>
          <Typography variant="h5" component="div" textAlign="center" sx={{marginBottom: "1rem"}}>
            {user?.name}
          </Typography>
          <Typography variant="body2" color="text.secondary" textAlign="center">
            {user?.profile_text}
          </Typography>
          <Stack direction="row" justifyContent="space-between" alignItems="center" spacing={2} mt={2}>
            {loginUser?.user_id !== idNumber ? (
              <Box sx={{display: "flex"}}>
                <Button variant="outlined" onClick={handleFollowClick} sx={{marginRight: "1rem", color: "black", borderColor: "black","&:hover": {borderColor: "black", color: 'black', backgroundColor: "rgba(0, 0, 0, 0.1)" }  }}>Follow</Button>
                <Button variant="outlined" onClick={handleMessageClick} sx={{color: "black", borderColor: "black","&:hover": {borderColor: "black", color: 'black', backgroundColor: "rgba(0, 0, 0, 0.1)" }  }}>Message</Button>
              </Box>
            ): <div></div>}
            <Box sx={{display: "flex"}}>
              <Box sx={{marginRight: "1rem"}}>
                <Typography variant="h6" component="div">
                  120
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Followers
                </Typography>
              </Box>
              <Box>
                <Typography variant="h6" component="div">
                  80
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Following
                </Typography>
              </Box>
            </Box>
          </Stack>
        </CardContent>
      </Card>

    </Box>
  );
}