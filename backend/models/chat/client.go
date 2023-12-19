package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ClientID          int
	RoomID            int
	WsConnect         *websocket.Conn
	SendCh            chan []byte
	ContextCancelFunc func()
}

func NewClient(ws *websocket.Conn, clientID, roomID int) *Client {
	return &Client{
		ClientID:  clientID,
		RoomID:    roomID,
		WsConnect: ws,
		SendCh:    make(chan []byte),
	}
}

func (c *Client) Disconnect(unregister chan<- *Client) {
	unregister <- c
	c.WsConnect.Close()
}
