import { useState, useRef, useEffect, useCallback } from 'react';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { Message } from '../../openapi';
import { SendMessageReqBody } from '../../openapi';
import SendIcon from '@mui/icons-material/Send';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import ReconnectingWebSocket from 'reconnecting-websocket';
import { useAuthUserContext } from '../../lib/auth/auth';
import { Typography, Paper } from '@mui/material';

type RoomProps = {
  roomID: number;
};

export default function Room(props: RoomProps) {
  const errorHandler = useErrorHandler();
  const [input, setInput] = useState<string>('');
  const [messages, setMessages] = useState<Message[]>([]);
  const socketRef = useRef<ReconnectingWebSocket | null>(null);
  const isConnectedRef = useRef<boolean>(false);
  const { user } = useAuthUserContext();

  useEffect(() => {
    const connect = (): Promise<ReconnectingWebSocket> => {
      isConnectedRef.current = false;
      return new Promise((resolve, reject) => {
        socketRef.current = new ReconnectingWebSocket(
          `${import.meta.env.VITE_WEBSOCKET_URL}/chats?room_id=${props.roomID}`,
        );

        if (socketRef.current) {
          socketRef.current.onopen = () => {
            isConnectedRef.current = true;
            if (socketRef.current) {
              resolve(socketRef.current);
            }
          };
          socketRef.current.onerror = (err: unknown) => {
            reject(err);
          };
        }
      });
    };

    connect()
      .then((socket) => {
        socket.onmessage = (msg: MessageEvent) => {
          const newMessageJson = JSON.parse(msg.data);
          const newMessage: Message = {
            message_id: newMessageJson.message_id,
            user_id: newMessageJson.user_id,
            content: newMessageJson.content,
            sent_at: newMessageJson.sent_at,
          };
          setMessages((prevMessages) => [...prevMessages, newMessage]);
        };
      })
      .catch((err) => {
        errorHandler(err);
      });

    return () => {
      if (isConnectedRef.current) {
        socketRef.current?.close();
      }
    };
  }, [props.roomID]);

  useEffect(() => {
    const fetchMessages = async () => {
      try {
        const api = await apiInstance;
        const res = await api.getMessages(props.roomID);

        if (res.data && Array.isArray(res.data)) {
          const fetchedMessages: Message[] = res.data.map((item) => ({
            message_id: item.message_id,
            user_id: item.user_id,
            content: item.content,
            sent_at: item.sent_at,
          }));
          setMessages(fetchedMessages);
        }
      } catch (error: unknown) {
        errorHandler(error);
      }
    };

    fetchMessages();
  }, [props.roomID]);

  const messagesEndRef = useRef<HTMLDivElement>(null);
  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'instant' });
  }, [messages]);

  const handleSendMessage = useCallback(() => {
    if (input.length === 0) return;
    const inputMessage: SendMessageReqBody = {
      room_id: props.roomID,
      content: input,
    };
    socketRef.current?.send(JSON.stringify(inputMessage));
    setInput('');
  }, [input]);

  return (
    <Container
      sx={{
        marginTop: '1rem',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        height: 'calc(100vh - 3rem - 8rem)',
        overflowY: 'auto',
        MsOverflowStyle: 'none',
        scrollbarWidth: 'none',
        '&::-webkit-scrollbar': {
          display: 'none',
        },
      }}
    >
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          width: '100%',
          flex: 1,
        }}
      >
        <>
          {messages.map((message) => (
            <Box key={message.message_id} sx={{ mb: 2 }}>
              <Box
                sx={{
                  display: 'flex',
                  justifyContent:
                    message.user_id === user.user_id
                      ? 'flex-end'
                      : 'flex-start',
                }}
              >
                <Box
                  sx={{
                    display: 'flex',
                    flexDirection:
                      message.user_id === user.user_id ? 'row-reverse' : 'row',
                    alignItems: 'center',
                  }}
                >
                  <Paper
                    variant="outlined"
                    sx={{
                      p: 2,
                      ml: message.user_id === user.user_id ? 0 : 1,
                      mr: message.user_id === user.user_id ? 1 : 0,
                      borderRadius:
                        message.user_id === user.user_id
                          ? '20px 20px 5px 20px'
                          : '20px 20px 20px 5px',
                    }}
                  >
                    <Typography variant="body1">{message.content}</Typography>
                  </Paper>
                </Box>
              </Box>
              <span
                style={{
                  display: 'flex',
                  justifyContent:
                    message.user_id === user.user_id
                      ? 'flex-end'
                      : 'flex-start',
                  fontSize: '0.7rem',
                  color: 'gray',
                  margin: '0.5rem',
                }}
              >
                {message.sent_at}
              </span>
            </Box>
          ))}
        </>
        <div ref={messagesEndRef} />
      </Box>
      <Box
        component="form"
        onSubmit={(e) => {
          e.preventDefault();
          handleSendMessage();
        }}
        noValidate
        sx={{
          display: 'flex',
          justifyContent: 'center',
          position: 'fixed',
          bottom: '1rem',
        }}
      >
        <TextField
          id="message-input"
          label="メッセージを入力"
          fullWidth
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
            },
          }}
          type="submit"
        >
          <SendIcon />
        </Button>
      </Box>
    </Container>
  );
}
