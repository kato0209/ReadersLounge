import * as React from 'react';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { ChatRoom } from '../../openapi';
import { Link } from "react-router-dom"

export default function RoomList() {

    const errorHandler = useErrorHandler();
    const [chatRooms, setChatRooms] = React.useState<ChatRoom[]>([]);

    React.useEffect(() => {
    const fetchPosts = async () => {
    
        try {
            const api = await apiInstance;
            const res = await api.getChatRooms();
            
            if (res.data && Array.isArray(res.data)) {
                const fetchedRooms: ChatRoom[] = res.data.map(item => ({
                    room_id: item.room_id,
                }));
                setChatRooms(fetchedRooms);
            }
        } catch (error: unknown) {
            errorHandler(error);
        }
            
    };

    fetchPosts();
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
            {chatRooms.length > 0 ? (
          <>
            {chatRooms.map(chatRoom => (
                <Box key={chatRoom.room_id}>
                    <Link to={"/chat-room"} state={{ roomID: chatRoom.room_id }}>
                        room
                    </Link>
                </Box>
            ))}
          </>
            ) : (
                <h2>チャットルームがありません</h2>
            )}
        </Box> 
    </Container>
    );
}