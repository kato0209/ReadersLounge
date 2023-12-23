import * as React from 'react';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { ChatRoom } from '../../openapi';
import { useLocation } from "react-router-dom"
import SendIcon from '@mui/icons-material/Send';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import ReconnectingWebSocket from 'reconnecting-websocket'

export default function Room() {

    const errorHandler = useErrorHandler();
    const { state } = useLocation();
    const [input, setInput] = React.useState<string>("");
    const [messages, setMessages] = React.useState<string[]>([]);
    const socketRef = React.useRef<ReconnectingWebSocket | null>(null);
    const isConnectedRef = React.useRef<boolean>(false);


    React.useEffect(() => {
        const connect = (): Promise<ReconnectingWebSocket> => {
            isConnectedRef.current = false;
            return new Promise((resolve, reject) => {
                socketRef.current = new ReconnectingWebSocket(`${import.meta.env.VITE_WEBSOCKET_URL}/chats?room_id=${state.roomID}`);
                
                if (socketRef.current) {
                    socketRef.current.onopen = () => {
                        isConnectedRef.current = true;
                        resolve(socketRef.current);
                    };
                    socketRef.current.onerror = (err: unknown) => {
                        reject(err);
                    };
                }
            });
        };

        connect().then((socket) => {
            socket.onmessage = (msg: MessageEvent) => {
                const newMessage = JSON.parse(msg.data as string);
                setMessages(prevMessages => [...prevMessages, newMessage]);
            };
            
        }).catch((err) => {
            errorHandler(err);
        });

        return () => {
            if (isConnectedRef.current) {
                socketRef.current.close();
            }
        };
    }, []);

    const handleSendMessage = React.useCallback(() => {
        if (input.length === 0) return;
        socketRef.current.send(JSON.stringify(input));
        setInput("");
    }, [input]);

    return (
    <Container component="main">
        <Box
            sx={{
            marginTop: '8rem',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            }}
        >
            <>
                {messages.map((message, index) => (
                    <div key={index}>{message}</div>
                ))}
            </>
            <>
            <Box 
                component="form" 
                noValidate
                sx={{
                    display: 'flex',
                    justifyContent: 'center',
                    width: '95%',
                }}
            >
                <TextField
                    id="message-input"
                    label="メッセージを入力"
                    sx = {{width: '100%'}}
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                />
                <Button 
                    variant="contained"
                    sx={{
                    backgroundColor: '#FF7E73',
                    color: '#fff',
                    '&:hover': {
                    backgroundColor: '#E56A67',
                    },
                    '&.Mui-disabled': {
                        backgroundColor: '#FFA49D',
                        color: '#fff',
                    }
                    }}
                    onClick={handleSendMessage}
                >
                    <SendIcon />
                </Button>
            </Box>
        </>
        </Box> 
    </Container>
    );
}