'use client';
import React, { useState, useEffect } from 'react';
import { isValidUrl } from '../../utils/isValidUrl';
import { useErrorHandler } from 'react-error-boundary';
import { User } from '../../openapi';
import { Connection } from '../../openapi';
import {
  Box,
  Card,
  CardContent,
  Typography,
  Button,
  Avatar,
  Stack,
  CardMedia,
} from '@mui/material';
import { EditProfile } from './EditProfile';
import { ConnectionList } from './ConnectionList';
import { Post } from '../../openapi';
import { redirect } from 'next/navigation';
import axios from 'axios';

export default function UserProfileComponent({
  user,
  posts,
  followerConnections,
  followingConnections,
  postListComponent,
}: {
  user: User;
  posts: Post[];
  followerConnections: Connection[];
  followingConnections: Connection[];
  postListComponent: React.ReactNode;
}) {
  const [followingConnection, setFollowingConnection] =
    useState<Connection | null>(null);
  const errorHandler = useErrorHandler();
  const [activeConnectionList, setActiveConnectionList] = useState<
    string | null
  >(null);

  const [loginUser, setLoginUser] = useState<User | null>(null);

  async function fetchLoginUser() {
    try {
      const res = await axios.get(`/api/fetch-login-user`);
      return res.data;
    } catch (error: unknown) {
      errorHandler(error);
    }
  }
  useEffect(() => {
    fetchLoginUser().then((res) => {
      setLoginUser(res.data);
    });
  }, []);

  useEffect(() => {
    const connection = followerConnections.find(
      (connection) => connection.target_user_id === loginUser?.user_id,
    );
    if (connection) {
      setFollowingConnection(connection);
    } else {
      setFollowingConnection(null);
    }
  }, [followerConnections]);

  const handleMessageClick = async (chatPartnerID: number) => {
    try {
      const roomID = await axios.get(
        `/api/create-chat-room?chatPartnerID=${chatPartnerID}`,
      );
      redirect(`/chat-room-list/${roomID}`);
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleFollowClick = async (connectionID: number) => {
    try {
      await axios.get(`/api/create-connection?connectionID=${connectionID}`);
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleUnFollowClick = async (connectionID: number) => {
    try {
      await axios.get(`/api/delete-connection?connectionID=${connectionID}`);
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  const handleFollowerClick = () => {
    setActiveConnectionList(
      activeConnectionList === 'followers' ? null : 'followers',
    );
  };

  const handleFollowingClick = () => {
    setActiveConnectionList(
      activeConnectionList === 'followings' ? null : 'followings',
    );
  };

  return (
    <Box sx={{ display: 'flex' }}>
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          p: 2,
          flex: '0 0 50%',
          '@media (max-width: 500px)': { flex: '0 0 60%' },
        }}
      >
        <Card sx={{ maxWidth: 500, width: '100%', mb: 2 }}>
          <CardMedia
            component="img"
            height="200"
            image="/images/UserProfileHeader.jpg"
            alt="Cover image"
          />
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'center',
              mt: -8,
              position: 'relative',
            }}
          >
            <Avatar
              sx={{
                width: 128,
                height: 128,
                border: '4px solid',
                borderColor: 'background.paper',
                '@media (max-width: 1000px)': {
                  width: 64,
                  height: 64,
                },
                '@media (max-width: 770px)': {
                  width: 50,
                  height: 50,
                },
                '@media (max-width: 600px)': {
                  width: 40,
                  height: 40,
                },
              }}
              src={
                isValidUrl(user?.profile_image)
                  ? user?.profile_image
                  : `data:image/png;base64,${user?.profile_image}`
              }
            />
            {loginUser?.user_id === user.user_id && (
              <Box
                sx={{
                  display: 'flex',
                  position: 'absolute',
                  top: '50%',
                  right: '0.3rem',
                }}
              >
                {user && <EditProfile user={user} />}
              </Box>
            )}
          </Box>
          <CardContent>
            <Typography
              variant="h5"
              component="div"
              textAlign="center"
              sx={{ marginBottom: '1rem' }}
            >
              {user?.name}
            </Typography>
            <Typography
              variant="body2"
              color="text.secondary"
              textAlign="center"
            >
              {user?.profile_text}
            </Typography>
            <Stack
              direction="row"
              justifyContent="space-between"
              alignItems="center"
              spacing={2}
              mt={4}
            >
              {loginUser?.user_id !== user.user_id ? (
                <Box sx={{ display: 'flex' }}>
                  {followingConnection ? (
                    <Button
                      variant="outlined"
                      onClick={() =>
                        handleUnFollowClick(followingConnection.connection_id)
                      }
                      sx={{
                        marginRight: '1rem',
                        color: 'black',
                        borderColor: 'black',
                        '&:hover': {
                          borderColor: 'black',
                          color: 'black',
                          backgroundColor: 'rgba(0, 0, 0, 0.1)',
                        },
                      }}
                    >
                      UnFollow
                    </Button>
                  ) : (
                    <Button
                      variant="outlined"
                      onClick={() => handleFollowClick(user.user_id)}
                      sx={{
                        marginRight: '1rem',
                        color: 'black',
                        borderColor: 'black',
                        '&:hover': {
                          borderColor: 'black',
                          color: 'black',
                          backgroundColor: 'rgba(0, 0, 0, 0.1)',
                        },
                      }}
                    >
                      Follow
                    </Button>
                  )}

                  <Button
                    variant="outlined"
                    onClick={() => handleMessageClick(user.user_id)}
                    sx={{
                      color: 'black',
                      borderColor: 'black',
                      '&:hover': {
                        borderColor: 'black',
                        color: 'black',
                        backgroundColor: 'rgba(0, 0, 0, 0.1)',
                      },
                    }}
                  >
                    Message
                  </Button>
                </Box>
              ) : (
                <div></div>
              )}
              <Box sx={{ display: 'flex' }}>
                <Box
                  sx={{
                    marginRight: '1rem',
                    cursor: 'pointer',
                    '&:hover': {
                      textDecoration: 'underline',
                    },
                  }}
                  onClick={handleFollowerClick}
                >
                  <Typography variant="body2" color="text.secondary">
                    {followerConnections.length} Followers
                  </Typography>
                </Box>
                <Box
                  sx={{
                    cursor: 'pointer',
                    '&:hover': {
                      textDecoration: 'underline',
                    },
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
        {posts && posts.length > 0 && (
          <Box sx={{ flex: 1, maxWidth: 500, width: '100%' }}>
            {postListComponent}
          </Box>
        )}
      </Box>
      <Box sx={{ flex: 1 }}>
        {activeConnectionList === 'followings' && (
          <ConnectionList connections={followingConnections} />
        )}
        {activeConnectionList === 'followers' && (
          <ConnectionList connections={followerConnections} />
        )}
      </Box>
    </Box>
  );
}
