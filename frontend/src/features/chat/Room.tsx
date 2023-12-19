import * as React from 'react';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { ChatRoom } from '../../openapi';
import { useLocation } from "react-router-dom"
import { set } from 'zod';



export default function Room() {

    const errorHandler = useErrorHandler();
    const { state } = useLocation();
    const [roomID, setRoomID] = React.useState<number>();

    React.useEffect(() => {
        setRoomID(state.roomID);
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
            <div>{roomID}</div>
        </Box> 
    </Container>
    );
}