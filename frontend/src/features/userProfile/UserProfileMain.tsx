import * as React from 'react';
import Container from '@mui/material/Container';
import { isValidUrl } from '../../utils/isValidUrl';
import { useParams } from 'react-router-dom';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { User } from '../../openapi';
import { CreateConnectionRequest, Connection } from '../../openapi/';
import { Box, Card, CardContent, Typography, Button, Avatar, Stack, CardMedia } from '@mui/material';
import UserHeaderImage from '../../assets/images/UserProfileHeader.jpg';
import { useAuthUserContext } from '../../lib/auth/auth';
import { CreateChatRoomRequest } from '../../openapi';
import { useNavigate } from 'react-router-dom';
import { EditProfile } from './EditProfile';
import { ConnectionList } from './ConnectionList';

export default function UserProfileMain() {

  const { id } = useParams<{ id: string }>();
  const idNumber = id ? parseInt(id, 10) : 0;
  const { user: loginUser } = useAuthUserContext();
  const [user, setUser] = React.useState<User | null>(null);
  const [followerConnections, setfollowerConnections] = React.useState<Connection[]>([]);
  const [followingConnections, setfollowingConnections] = React.useState<Connection[]>([]);
  const [followingConnection, setFollowingConnection] = React.useState<Connection | null>(null);
  const [isFollowActionLoading, setIsFollowActionLoading] = React.useState<boolean>(false);
  const errorHandler = useErrorHandler();
  const navigate = useNavigate();
  const [activeConnectionList, setActiveConnectionList] = React.useState<string | null>(null);

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

  const fetchfollowerConnections = async () => {
    try {
        setIsFollowActionLoading(true);
        const api = await apiInstance;
        const res = await api.getFollowerConnections(idNumber);
        if (res.data && res.data.length > 0) {
          const followerConnections: Connection[] = res.data.map((connection) => {
            return {
              connection_id: connection.connection_id,
              target_user_id: connection.target_user_id,
              target_user_name: connection.target_user_name,
              target_user_profile_image: connection.target_user_profile_image,
            }
          });
          setfollowerConnections(followerConnections);
        } else {
          setfollowerConnections([]);
        }
    } catch (error: unknown) {
        errorHandler(error);
    } finally {
      setIsFollowActionLoading(false);
    }
  }

  const fetchfollowingConnections = async () => {
    try {
        setIsFollowActionLoading(true);
        const api = await apiInstance;
        const res = await api.getFollowingConnections(idNumber);
        if (res.data && res.data.length > 0) {
          const followingConnections: Connection[] = res.data.map((connection) => {
            return {
              connection_id: connection.connection_id,
              target_user_id: connection.target_user_id,
              target_user_name: connection.target_user_name,
              target_user_profile_image: connection.target_user_profile_image,
            }
          });
          setfollowingConnections(followingConnections);
        } else {
          setfollowingConnections([]);
        }
    } catch (error: unknown) {
        errorHandler(error);
    } finally {
      setIsFollowActionLoading(false);
    }
  }

  React.useEffect(() => {
    fetchUser();
    fetchfollowerConnections();
    fetchfollowingConnections();
  }, [idNumber]);

  React.useEffect(() => {
    const connection = followerConnections.find(connection => connection.target_user_id === loginUser.user_id);
    if (connection) {
      setFollowingConnection(connection);
    } else {
      setFollowingConnection(null);
    }
  }, [followerConnections]);

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
        fetchfollowerConnections();
      }
    } catch (error: unknown) {
        errorHandler(error);
    }
  }

  const handleUnFollowClick = async (connectionID: number) => {
    try {
      const api = await apiInstance;
      const res = await api.deleteConnection(connectionID);
      if (res.status === 204) {
        fetchfollowerConnections();
      }
    } catch (error: unknown) {
        errorHandler(error);
    }
  }

  const handleFollowerClick = () => {
    setActiveConnectionList(activeConnectionList === 'followers' ? null : 'followers');
  };
  
  const handleFollowingClick = () => {
    setActiveConnectionList(activeConnectionList === 'followings' ? null : 'followings');
  };

  return (
    <div style={{display: "flex"}}>
      <Box sx={{ display: 'flex', flexDirection: 'column', p: 2, flex: '0 0 50%' }}>
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
            <Stack direction="row" justifyContent="space-between" alignItems="center" spacing={2} mt={4}>
              {loginUser?.user_id !== idNumber ? (
                <Box sx={{display: "flex"}}>
                  {followingConnection 
                    ? (<Button disabled={isFollowActionLoading} variant="outlined" onClick={() => handleUnFollowClick(followingConnection.connection_id)} sx={{marginRight: "1rem", color: "black", borderColor: "black","&:hover": {borderColor: "black", color: 'black', backgroundColor: "rgba(0, 0, 0, 0.1)" }  }}>UnFollow</Button>)
                    : (<Button disabled={isFollowActionLoading} variant="outlined" onClick={handleFollowClick} sx={{marginRight: "1rem", color: "black", borderColor: "black","&:hover": {borderColor: "black", color: 'black', backgroundColor: "rgba(0, 0, 0, 0.1)" }  }}>Follow</Button>)
                  }
                  
                  <Button variant="outlined" onClick={handleMessageClick} sx={{color: "black", borderColor: "black","&:hover": {borderColor: "black", color: 'black', backgroundColor: "rgba(0, 0, 0, 0.1)" }  }}>Message</Button>
                </Box>
              ): <div></div>}
              <Box sx={{display: "flex"}}>
                <Box 
                  sx={{
                    marginRight: "1rem", 
                    cursor: "pointer",
                    "&:hover": {
                      textDecoration: "underline",
                    }
                  }} 
                  onClick={handleFollowerClick}
                >
                  <Typography variant="body2" color="text.secondary">
                    {followerConnections.length} Followers
                  </Typography>
                </Box>
                <Box 
                  sx={{ 
                    cursor: "pointer",
                    "&:hover": {
                      textDecoration: "underline",
                    }
                  }} 
                  onClick={handleFollowingClick}
                >
                  <Typography variant="body2" color="text.secondary">
                    {followingConnections.length} Following
                  </Typography>
                </Box>
              </Box>
            </Stack>
          </CardContent>
        </Card>
      </Box>
      <Box sx={{flex: 1}}>
        {activeConnectionList === 'followings' && <ConnectionList connections={followingConnections} />}
        {activeConnectionList === 'followers' && <ConnectionList connections={followerConnections} />}
      </Box>
    </div>
  );
}