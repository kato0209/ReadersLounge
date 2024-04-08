import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import { apiInstance } from '../../lib/api/apiInstance';
import { ChatRoom } from '../../openapi';
import Sidebar from '../../components/Sidebar/Sidebar';
import UserAvatar from '../../components/Avatar/UserAvatar';
import Typography from '@mui/material/Typography';
import Room from './Room';
import { useSearchParams, redirect } from 'next/navigation';
import useMediaQuery from '@mui/material/useMediaQuery';
import { getAllCookies } from '../../utils/getCookies';

export default async function RoomList() {
  const isMobile = useMediaQuery('(max-width:650px)');
  const searchParams = useSearchParams();
  const id = searchParams.get('id');
  const roomID = id ? parseInt(id, 10) : 0;

  const fetchPosts = async (): Promise<ChatRoom[]> => {
    try {
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getChatRooms({ headers: { Cookie: cookie } });
      if (res.data && Array.isArray(res.data)) {
        const fetchedRooms: ChatRoom[] = res.data.map((item) => ({
          room_id: item.room_id,
          target_user_id: item.target_user_id,
          target_user_name: item.target_user_name,
          target_user_profile_image: item.target_user_profile_image,
          last_message: item.last_message,
          last_message_sent_at: item.last_message_sent_at,
        }));
        return fetchedRooms;
      } else {
        return [];
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const chatRooms = await fetchPosts();

  return (
    <div style={{ display: 'flex' }}>
      {!isMobile && (
        <div style={{ flex: '0 0 30%', display: 'flex' }}>
          <Sidebar />
        </div>
      )}
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
                    redirect(`/chat-room-list/${chatRoom.room_id}`)
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
          <Room roomID={roomID} />
        ) : (
          <Container component="main">
            <h3>Select a ChatRoom</h3>
          </Container>
        )}
      </div>
    </div>
  );
}
