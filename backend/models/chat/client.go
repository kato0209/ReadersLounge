package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	WsConnect *websocket.Conn
	SendCh    chan []byte
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		WsConnect: ws,
		SendCh:    make(chan []byte),
	}
}

func (c *Client) Disconnect(unregister chan<- *Client) {
	unregister <- c
	c.WsConnect.Close()
}
