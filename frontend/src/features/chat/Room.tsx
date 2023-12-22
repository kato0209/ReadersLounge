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
import * as WebSocket from "websocket";


export default function Room() {

    const errorHandler = useErrorHandler();
    const location = useLocation();
    const { state } = useLocation();
    const [input, setInput] = React.useState<string>("");
    const [webSocket, setWebSocket] = React.useState<WebSocket.w3cwebsocket>();
    const [messages, setMessages] = React.useState<string[]>([]);

    React.useEffect(() => {
        const connect = (): Promise<WebSocket.w3cwebsocket> => {
            return new Promise((resolve, reject) => {
              const socket = new WebSocket.w3cwebsocket(`${import.meta.env.VITE_WEBSOCKET_URL}/chats?room_id=${state.roomID}`);
              socket.onopen = () => {
                console.log("connected");
                resolve(socket);
              };
              socket.onclose = () => {
                alert("接続が切れました。");
                console.log("disconnected");
              };
              socket.onerror = (err: unknown) => {
                reject(err);
              };
            });
        };
        
        connect().then((socket) => {
            socket.onmessage = (msg: MessageEvent) => {
                const newMessage = JSON.parse(msg.data as string);
                setMessages(prevMessages => [...prevMessages, newMessage]);
            };
            setWebSocket(socket);
        }).catch((err) => {
            errorHandler(err);
        });
    }, []);

    const handleSendMessage = () => {
        if (input.length === 0) return;
        webSocket.send(JSON.stringify(input));
        setInput("");
    };

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