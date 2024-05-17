'use client';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import Sidebar from '../../components/Sidebar/Sidebar';
import UserAvatar from '../../components/Avatar/UserAvatar';
import Typography from '@mui/material/Typography';
import Room from './Room';
import { useRouter } from 'next/navigation';
import { ChatRoom, Message } from '../../openapi';

export default async function RoomList({
  roomID,
  chatRooms,
  messages,
}: {
  roomID: number;
  chatRooms: ChatRoom[];
  messages: Message[];
}) {
  const router = useRouter();
  return (
    <div style={{ display: 'flex' }}>
      <div className="isMobile" style={{ flex: '0 0 30%', display: 'flex' }}>
        <Sidebar />
      </div>
      <div
        style={{
          flex: '0 0 35%',
          borderRight: '1px solid #BDBDBD',
          borderLeft: '1px solid #BDBDBD',
          height: 'calc(100vh - 3rem)',
        }}
      >
        <h3 style={{ marginLeft: '1rem' }}>Rooms</h3>
        <Box
          sx={{
            marginTop: '1rem',
            display: 'flex',
            flexDirection: 'column',
          }}
        >
          {chatRooms.length > 0 ? (
            <>
              {chatRooms.map((chatRoom) => (
                <Box
                  key={chatRoom.room_id}
                  onClick={() =>
                    router.push(`/chat-room-list/${chatRoom.room_id}`)
                  }
                  sx={{
                    display: 'flex',
                    alignItems: 'center',
                    paddingLeft: '1rem',
                    cursor: 'pointer',
                    '&:hover': {
                      color: 'inherit',
                      backgroundColor: '#EAE6E0',
                    },
                  }}
                >
                  <UserAvatar
                    image={chatRoom.target_user_profile_image}
                    userID={chatRoom.target_user_id}
                  />
                  <Box sx={{ margin: '0.5rem' }}>
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <Typography variant="h6" color="black">
                        {chatRoom.target_user_name}
                      </Typography>
                      <Typography color="gray" sx={{ marginLeft: '0.5rem' }}>
                        {chatRoom.last_message_sent_at}
                      </Typography>
                    </Box>
                    <Typography color="gray" style={{ wordWrap: 'break-word' }}>
                      {chatRoom.last_message}
                    </Typography>
                  </Box>
                </Box>
              ))}
            </>
          ) : (
            <Box sx={{ marginLeft: '1rem' }}>
              <h2>チャットルームがありません</h2>
            </Box>
          )}
        </Box>
      </div>
      <div style={{ flex: '1', display: 'flex' }}>
        {roomID ? (
          <Room roomID={roomID} messages={messages} />
        ) : (
          <Container component="main">
            <h3>Select a ChatRoom</h3>
          </Container>
        )}
      </div>
    </div>
  );
}
